package clients

import (
	"context"
	"encoding/json"
	"testing"

	v1 "github.com/crossplane/crossplane-runtime/v2/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	resourcefake "github.com/crossplane/crossplane-runtime/v2/pkg/resource/fake"
	"github.com/google/go-cmp/cmp"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/utils/ptr"
	ctrlclient "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	cv1beta1 "github.com/grafana/crossplane-provider-grafana/v2/apis/cluster/v1beta1"
	nv1beta1 "github.com/grafana/crossplane-provider-grafana/v2/apis/namespaced/v1beta1"
)

func createScheme(t *testing.T) *runtime.Scheme {
	t.Helper()
	scheme := runtime.NewScheme()
	_ = nv1beta1.SchemeBuilder.AddToScheme(scheme)
	_ = cv1beta1.SchemeBuilder.AddToScheme(scheme)
	_ = corev1.AddToScheme(scheme)
	return scheme
}

func createClusterProviderConfig(t *testing.T, orgID, stackID *int) *nv1beta1.ClusterProviderConfig {
	t.Helper()
	return &nv1beta1.ClusterProviderConfig{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-config",
			Namespace: "default",
		},
		Spec: nv1beta1.ProviderConfigSpec{
			URL:     "https://example.grafana.com",
			OrgID:   orgID,
			StackID: stackID,
			Credentials: nv1beta1.ProviderCredentials{
				Source: v1.CredentialsSourceSecret,
				CommonCredentialSelectors: v1.CommonCredentialSelectors{
					SecretRef: &v1.SecretKeySelector{
						SecretReference: v1.SecretReference{
							Name:      "test-secret",
							Namespace: "default",
						},
						Key: "credentials",
					},
				},
			},
		},
		Status: nv1beta1.ProviderConfigStatus{
			ProviderConfigStatus: v1.ProviderConfigStatus{
				ConditionedStatus: v1.ConditionedStatus{
					Conditions: []v1.Condition{},
				},
			},
		},
	}
}

func createNamespacedProviderConfig(t *testing.T, orgID, stackID *int) *nv1beta1.ProviderConfig {
	t.Helper()
	return &nv1beta1.ProviderConfig{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-config",
			Namespace: "default",
		},
		Spec: nv1beta1.ProviderConfigSpec{
			URL:     "https://example.grafana.com",
			OrgID:   orgID,
			StackID: stackID,
			Credentials: nv1beta1.ProviderCredentials{
				Source: v1.CredentialsSourceSecret,
				CommonCredentialSelectors: v1.CommonCredentialSelectors{
					SecretRef: &v1.SecretKeySelector{
						SecretReference: v1.SecretReference{
							Name:      "test-secret",
							Namespace: "default",
						},
						Key: "credentials",
					},
				},
			},
		},
		Status: nv1beta1.ProviderConfigStatus{
			ProviderConfigStatus: v1.ProviderConfigStatus{
				ConditionedStatus: v1.ConditionedStatus{
					Conditions: []v1.Condition{},
				},
			},
		},
	}
}

func createLegacyProviderConfig(t *testing.T, orgID, stackID *int) *cv1beta1.ProviderConfig {
	t.Helper()
	return &cv1beta1.ProviderConfig{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-config",
		},
		Spec: cv1beta1.ProviderConfigSpec{
			OrgID:   orgID,
			StackID: stackID,
			Credentials: cv1beta1.ProviderCredentials{
				Source: v1.CredentialsSourceSecret,
				CommonCredentialSelectors: v1.CommonCredentialSelectors{
					SecretRef: &v1.SecretKeySelector{
						SecretReference: v1.SecretReference{
							Name:      "test-secret",
							Namespace: "default",
						},
						Key: "credentials",
					},
				},
			},
		},
		Status: cv1beta1.ProviderConfigStatus{
			ProviderConfigStatus: v1.ProviderConfigStatus{
				ConditionedStatus: v1.ConditionedStatus{
					Conditions: []v1.Condition{},
				},
			},
		},
	}
}

func createModernManaged(t *testing.T, kind string) resource.ModernManaged {
	t.Helper()
	return &resourcefake.ModernManaged{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-resource",
			Namespace: "default",
			UID:       "test-uid-12345",
		},
		TypedProviderConfigReferencer: resourcefake.TypedProviderConfigReferencer{
			Ref: &v1.ProviderConfigReference{
				Kind: kind,
				Name: "test-config",
			},
		},
	}
}

func createLegacyManaged(t *testing.T) resource.LegacyManaged {
	t.Helper()
	return &resourcefake.LegacyManaged{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-resource",
			Namespace: "default",
			UID:       "test-uid-12345",
		},
		LegacyProviderConfigReferencer: resourcefake.LegacyProviderConfigReferencer{
			Ref: &v1.Reference{Name: "test-config"},
		},
	}
}

// setupTestNamespaced creates a namespaced ProviderConfig for testing the modern namespaced code path
func setupTestNamespaced(t *testing.T, credentials map[string]string, orgID, stackID *int) (ctrlclient.Client, resource.Managed) {
	t.Helper()

	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-secret",
			Namespace: "default",
		},
		Data: map[string][]byte{},
	}

	pc := createNamespacedProviderConfig(t, orgID, stackID)

	credData, err := json.Marshal(credentials)
	if err != nil {
		t.Fatal(err)
	}
	secret.Data["credentials"] = credData

	mg := createModernManaged(t, "ProviderConfig")
	scheme := createScheme(t)

	client := fake.NewClientBuilder().
		WithScheme(scheme).
		WithObjects(secret, pc).
		WithStatusSubresource(&nv1beta1.ProviderConfig{}).
		Build()

	return client, mg
}

// setupTestNamespaced creates a ClusterProviderConfig for testing the modern cluster-scoped code path
func setupTestClusterScoped(t *testing.T, credentials map[string]string, orgID, stackID *int) (ctrlclient.Client, resource.Managed) {
	t.Helper()

	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-secret",
			Namespace: "default",
		},
		Data: map[string][]byte{},
	}

	pc := createClusterProviderConfig(t, orgID, stackID)

	credData, err := json.Marshal(credentials)
	if err != nil {
		t.Fatal(err)
	}
	secret.Data["credentials"] = credData

	mg := createModernManaged(t, "ClusterProviderConfig")
	scheme := createScheme(t)

	client := fake.NewClientBuilder().
		WithScheme(scheme).
		WithObjects(secret, pc).
		WithStatusSubresource(&nv1beta1.ClusterProviderConfig{}).
		Build()

	return client, mg
}

// setupTestLegacy creates a cluster-scoped ProviderConfig for testing the legacy cluster-scoped code path
// This mirrors the main branch behavior where only cluster-scoped ProviderConfigs existed
func setupTestLegacy(t *testing.T, credentials map[string]string, orgID, stackID *int) (ctrlclient.Client, resource.Managed) {
	t.Helper()

	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-secret",
			Namespace: "default",
		},
		Data: map[string][]byte{},
	}

	pc := createLegacyProviderConfig(t, orgID, stackID)

	credData, err := json.Marshal(credentials)
	if err != nil {
		t.Fatal(err)
	}
	secret.Data["credentials"] = credData

	mg := createLegacyManaged(t)
	scheme := createScheme(t)

	client := fake.NewClientBuilder().
		WithScheme(scheme).
		WithObjects(secret, pc).
		WithStatusSubresource(&cv1beta1.ProviderConfig{}).
		Build()

	return client, mg
}

func TestTerraformSetupBuilderNamespaced(t *testing.T) {
	cases := []struct {
		name        string
		credentials map[string]string
		orgID       *int
		stackID     *int
		want        map[string]any
		wantMissing []string
	}{
		{
			name: "OrgID override takes precedence over credentials",
			credentials: map[string]string{
				"url":    "https://example.grafana.com",
				"auth":   "token",
				"org_id": "999",
			},
			orgID: ptr.To(123),
			want: map[string]any{
				"auth":   "token",
				"url":    "https://example.grafana.com",
				"org_id": 123,
			},
			wantMissing: []string{"stack_id"},
		},
		{
			name: "StackID override takes precedence over credentials",
			credentials: map[string]string{
				"url":      "https://example.grafana.com",
				"auth":     "token",
				"stack_id": "999",
			},
			stackID: ptr.To(456),
			want: map[string]any{
				"auth":     "token",
				"url":      "https://example.grafana.com",
				"stack_id": 456,
			},
			wantMissing: []string{"org_id"},
		},
		{
			name: "OrgID and StackID zero overrides are used",
			credentials: map[string]string{
				"url":      "https://example.grafana.com",
				"auth":     "token",
				"org_id":   "999",
				"stack_id": "888",
			},
			orgID:   ptr.To(0),
			stackID: ptr.To(0),
			want: map[string]any{
				"auth":     "token",
				"url":      "https://example.grafana.com",
				"org_id":   0,
				"stack_id": 0,
			},
		},
		{
			name: "Credentials are used when overrides are absent",
			credentials: map[string]string{
				"url":      "https://example.grafana.com",
				"auth":     "token",
				"org_id":   "789",
				"stack_id": "321",
			},
			want: map[string]any{
				"auth":     "token",
				"url":      "https://example.grafana.com",
				"org_id":   "789",
				"stack_id": "321",
			},
		},
		{
			name: "OrgID and StackID omitted when not provided",
			credentials: map[string]string{
				"url":  "https://example.grafana.com",
				"auth": "token",
			},
			want: map[string]any{
				"auth": "token",
				"url":  "https://example.grafana.com",
			},
			wantMissing: []string{"org_id", "stack_id"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			client, mg := setupTestNamespaced(t, tc.credentials, tc.orgID, tc.stackID)

			setupFn := TerraformSetupBuilder()
			setup, err := setupFn(context.Background(), client, mg)
			if err != nil {
				t.Fatalf("TerraformSetupBuilder() error = %v", err)
			}

			pc := &nv1beta1.ProviderConfig{}
			key := types.NamespacedName{Name: "test-config", Namespace: "default"}
			client.Get(context.Background(), key, pc)

			if len(pc.Status.Conditions) == 0 {
				t.Fatalf("expected ProviderConfig status to have Ready condition set, got none")
			}

			for key, val := range tc.want {
				got, ok := setup.Configuration[key]
				if !ok {
					t.Errorf("expected %s to be set in configuration", key)
					continue
				}

				if diff := cmp.Diff(val, got); diff != "" {
					t.Errorf("%s mismatch (-want +got):\n%s", key, diff)
				}
			}

			for _, key := range tc.wantMissing {
				if _, ok := setup.Configuration[key]; ok {
					t.Errorf("expected %s to be absent from configuration", key)
				}
			}
		})
	}
}

func TestTerraformSetupBuilderClusterScoped(t *testing.T) {
	cases := []struct {
		name        string
		credentials map[string]string
		orgID       *int
		stackID     *int
		want        map[string]any
		wantMissing []string
	}{
		{
			name: "OrgID override takes precedence over credentials",
			credentials: map[string]string{
				"url":    "https://example.grafana.com",
				"auth":   "token",
				"org_id": "999",
			},
			orgID: ptr.To(123),
			want: map[string]any{
				"auth":   "token",
				"url":    "https://example.grafana.com",
				"org_id": 123,
			},
			wantMissing: []string{"stack_id"},
		},
		{
			name: "StackID override takes precedence over credentials",
			credentials: map[string]string{
				"url":      "https://example.grafana.com",
				"auth":     "token",
				"stack_id": "999",
			},
			stackID: ptr.To(456),
			want: map[string]any{
				"auth":     "token",
				"url":      "https://example.grafana.com",
				"stack_id": 456,
			},
			wantMissing: []string{"org_id"},
		},
		{
			name: "OrgID and StackID zero overrides are used",
			credentials: map[string]string{
				"url":      "https://example.grafana.com",
				"auth":     "token",
				"org_id":   "999",
				"stack_id": "888",
			},
			orgID:   ptr.To(0),
			stackID: ptr.To(0),
			want: map[string]any{
				"auth":     "token",
				"url":      "https://example.grafana.com",
				"org_id":   0,
				"stack_id": 0,
			},
		},
		{
			name: "Credentials are used when overrides are absent",
			credentials: map[string]string{
				"url":      "https://example.grafana.com",
				"auth":     "token",
				"org_id":   "789",
				"stack_id": "321",
			},
			want: map[string]any{
				"auth":     "token",
				"url":      "https://example.grafana.com",
				"org_id":   "789",
				"stack_id": "321",
			},
		},
		{
			name: "OrgID and StackID omitted when not provided",
			credentials: map[string]string{
				"url":  "https://example.grafana.com",
				"auth": "token",
			},
			want: map[string]any{
				"auth": "token",
				"url":  "https://example.grafana.com",
			},
			wantMissing: []string{"org_id", "stack_id"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			client, mg := setupTestClusterScoped(t, tc.credentials, tc.orgID, tc.stackID)

			setupFn := TerraformSetupBuilder()
			setup, err := setupFn(context.Background(), client, mg)
			if err != nil {
				t.Fatalf("TerraformSetupBuilder() error = %v", err)
			}

			pc := &nv1beta1.ClusterProviderConfig{}
			key := types.NamespacedName{Name: "test-config", Namespace: "default"}
			client.Get(context.Background(), key, pc)

			if len(pc.Status.Conditions) == 0 {
				t.Fatalf("expected ClusterProviderConfig to have Ready condition set, got none")
			}

			for key, val := range tc.want {
				got, ok := setup.Configuration[key]
				if !ok {
					t.Errorf("expected %s to be set in configuration", key)
					continue
				}

				if diff := cmp.Diff(val, got); diff != "" {
					t.Errorf("%s mismatch (-want +got):\n%s", key, diff)
				}
			}

			for _, key := range tc.wantMissing {
				if _, ok := setup.Configuration[key]; ok {
					t.Errorf("expected %s to be absent from configuration", key)
				}
			}
		})
	}
}

func TestTerraformSetupBuilderLegacy(t *testing.T) {
	cases := []struct {
		name        string
		credentials map[string]string
		orgID       *int
		stackID     *int
		want        map[string]any
		wantMissing []string
	}{
		{
			name: "OrgID override takes precedence over credentials",
			credentials: map[string]string{
				"url":    "https://example.grafana.com",
				"auth":   "token",
				"org_id": "999",
			},
			orgID: ptr.To(123),
			want: map[string]any{
				"auth":   "token",
				"url":    "https://example.grafana.com",
				"org_id": 123,
			},
			wantMissing: []string{"stack_id"},
		},
		{
			name: "StackID override takes precedence over credentials",
			credentials: map[string]string{
				"url":      "https://example.grafana.com",
				"auth":     "token",
				"stack_id": "999",
			},
			stackID: ptr.To(456),
			want: map[string]any{
				"auth":     "token",
				"url":      "https://example.grafana.com",
				"stack_id": 456,
			},
			wantMissing: []string{"org_id"},
		},
		{
			name: "OrgID and StackID zero overrides are used",
			credentials: map[string]string{
				"url":      "https://example.grafana.com",
				"auth":     "token",
				"org_id":   "999",
				"stack_id": "888",
			},
			orgID:   ptr.To(0),
			stackID: ptr.To(0),
			want: map[string]any{
				"auth":     "token",
				"url":      "https://example.grafana.com",
				"org_id":   0,
				"stack_id": 0,
			},
		},
		{
			name: "Credentials are used when overrides are absent",
			credentials: map[string]string{
				"url":      "https://example.grafana.com",
				"auth":     "token",
				"org_id":   "789",
				"stack_id": "321",
			},
			want: map[string]any{
				"auth":     "token",
				"url":      "https://example.grafana.com",
				"org_id":   "789",
				"stack_id": "321",
			},
		},
		{
			name: "OrgID and StackID omitted when not provided",
			credentials: map[string]string{
				"url":  "https://example.grafana.com",
				"auth": "token",
			},
			want: map[string]any{
				"auth": "token",
				"url":  "https://example.grafana.com",
			},
			wantMissing: []string{"org_id", "stack_id"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			client, mg := setupTestLegacy(t, tc.credentials, tc.orgID, tc.stackID)

			setupFn := TerraformSetupBuilder()
			setup, err := setupFn(context.Background(), client, mg)
			if err != nil {
				t.Fatalf("TerraformSetupBuilder() error = %v", err)
			}

			pc := &cv1beta1.ProviderConfig{}
			key := types.NamespacedName{Name: "test-config"}
			client.Get(context.Background(), key, pc)

			if len(pc.Status.Conditions) == 0 {
				t.Fatalf("expected ProviderConfig status to have Ready condition set, got none")
			}

			for key, val := range tc.want {
				got, ok := setup.Configuration[key]
				if !ok {
					t.Errorf("expected %s to be set in configuration", key)
					continue
				}

				if diff := cmp.Diff(val, got); diff != "" {
					t.Errorf("%s mismatch (-want +got):\n%s", key, diff)
				}
			}

			for _, key := range tc.wantMissing {
				if _, ok := setup.Configuration[key]; ok {
					t.Errorf("expected %s to be absent from configuration", key)
				}
			}
		})
	}
}

func TestLegacyNoProviderConfig(t *testing.T) {
	client := fake.NewClientBuilder().
		WithScheme(createScheme(t)).
		Build()

	mg := createLegacyManaged(t)

	setupFn := TerraformSetupBuilder()
	_, err := setupFn(context.Background(), client, mg)

	if err == nil {
		t.Fatalf("expected error when ProviderConfig is missing, got none")
	}
}

func TestClusterScopedNoProviderConfig(t *testing.T) {
	client := fake.NewClientBuilder().
		WithScheme(createScheme(t)).
		Build()

	mg := createModernManaged(t, "ClusterProviderConfig")

	setupFn := TerraformSetupBuilder()
	_, err := setupFn(context.Background(), client, mg)

	if err == nil {
		t.Fatalf("expected error when ProviderConfig is missing, got none")
	}
}

func TestNamespacedNoProviderConfig(t *testing.T) {
	client := fake.NewClientBuilder().
		WithScheme(createScheme(t)).
		Build()

	mg := createModernManaged(t, "ProviderConfig")

	setupFn := TerraformSetupBuilder()
	_, err := setupFn(context.Background(), client, mg)

	if err == nil {
		t.Fatalf("expected error when ProviderConfig is missing, got none")
	}
}
