/*
Copyright 2026 The Crossplane Authors.
*/

package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	metricsserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"
)

func TestSecretCacheClientOptions(t *testing.T) {
	for _, enabled := range []bool{true, false} {
		name := "direct"
		if enabled {
			name = "cached"
		}
		t.Run(name, func(t *testing.T) {
			var requests []string
			var mu sync.Mutex
			observedRequests := func() []string {
				mu.Lock()
				defer mu.Unlock()
				return append([]string(nil), requests...)
			}
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				mu.Lock()
				requests = append(requests, r.Method+" "+r.URL.Path)
				mu.Unlock()
				w.Header().Set("Content-Type", "application/json")
				if r.Method == http.MethodGet && r.URL.Path == "/api/v1/namespaces/provider-system/secrets/credentials" {
					_ = json.NewEncoder(w).Encode(&corev1.Secret{
						TypeMeta:   metav1.TypeMeta{APIVersion: "v1", Kind: "Secret"},
						ObjectMeta: metav1.ObjectMeta{Name: "credentials", Namespace: "provider-system"},
						Data:       map[string][]byte{"credentials": []byte("test-value")},
					})
					return
				}
				w.WriteHeader(http.StatusForbidden)
				_ = json.NewEncoder(w).Encode(&metav1.Status{
					TypeMeta: metav1.TypeMeta{APIVersion: "v1", Kind: "Status"},
					Status:   metav1.StatusFailure, Reason: metav1.StatusReasonForbidden,
					Code: http.StatusForbidden, Message: "request is not authorized",
				})
			}))
			defer server.Close()

			scheme := runtime.NewScheme()
			if err := corev1.AddToScheme(scheme); err != nil {
				t.Fatal(err)
			}
			mapper := meta.NewDefaultRESTMapper([]schema.GroupVersion{corev1.SchemeGroupVersion})
			mapper.Add(corev1.SchemeGroupVersion.WithKind("Secret"), meta.RESTScopeNamespace)
			mapper.Add(corev1.SchemeGroupVersion.WithKind("Pod"), meta.RESTScopeNamespace)
			mgr, err := ctrl.NewManager(&rest.Config{Host: server.URL}, ctrl.Options{
				Scheme:         scheme,
				MapperProvider: func(*rest.Config, *http.Client) (meta.RESTMapper, error) { return mapper, nil },
				Client:         secretCacheClientOptions(enabled),
				Metrics:        metricsserver.Options{BindAddress: "0"},
			})
			if err != nil {
				t.Fatal(err)
			}
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			secret := &corev1.Secret{}
			err = mgr.GetClient().Get(ctx, types.NamespacedName{Namespace: "provider-system", Name: "credentials"}, secret)
			if enabled {
				recorded := observedRequests()
				// An unstarted manager cannot serve a cached read.
				if err == nil || len(recorded) != 0 {
					t.Fatalf("expected a cached read without direct API requests: err=%v requests=%v", err, recorded)
				}
				return
			}
			if err != nil || string(secret.Data["credentials"]) != "test-value" {
				t.Fatalf("named Secret read failed: %v", err)
			}
			err = mgr.GetClient().Get(ctx, types.NamespacedName{Namespace: "other-system", Name: "credentials"}, &corev1.Secret{})
			if !apierrors.IsForbidden(err) {
				t.Fatalf("expected the API authorization denial to be preserved, got %v", err)
			}
			// Other resource kinds must continue to use the manager cache.
			_ = mgr.GetClient().Get(ctx, types.NamespacedName{Namespace: "provider-system", Name: "workload"}, &corev1.Pod{})
			want := []string{
				"GET /api/v1/namespaces/provider-system/secrets/credentials",
				"GET /api/v1/namespaces/other-system/secrets/credentials",
			}
			got := observedRequests()
			if len(got) != len(want) {
				t.Fatalf("unexpected requests (no Secret list/watch should be needed): %v", got)
			}
			for i := range want {
				if got[i] != want[i] {
					t.Fatalf("request %d: got %q, want %q", i, got[i], want[i])
				}
			}
		})
	}
}
