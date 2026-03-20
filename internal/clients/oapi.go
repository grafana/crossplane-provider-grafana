/*
Copyright 2025 Grafana
*/

package clients

import (
	"context"
	"encoding/json"
	"net/url"
	"strings"

	goapi "github.com/grafana/grafana-openapi-client-go/client"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
)

// ExtractModernConfig extracts the ProviderConfig and raw credential map for a
// namespaced managed resource. It reuses the unexported useModernProviderConfig
// logic in this package.
func ExtractModernConfig(ctx context.Context, c client.Client, mg resource.ModernManaged) (*Config, map[string]string, error) {
	cfg, err := useModernProviderConfig(ctx, c, mg)
	if err != nil {
		return nil, nil, err
	}

	data, err := resource.CommonCredentialExtractor(ctx, cfg.Credentials.Source, c, cfg.Credentials.CommonCredentialSelectors)
	if err != nil {
		return nil, nil, errors.Wrap(err, errExtractCredentials)
	}

	creds := map[string]string{}
	if err := json.Unmarshal(data, &creds); err != nil {
		return nil, nil, errors.Wrap(err, errUnmarshalCredentials)
	}

	return cfg, creds, nil
}

// NewOAPIClient creates a Grafana OpenAPI HTTP client from a Grafana URL and
// auth string. If auth is in "user:password" format, Basic auth is used;
// otherwise it is treated as a Bearer token.
func NewOAPIClient(grafanaURL, auth string) (*goapi.GrafanaHTTPAPI, error) {
	u, err := url.Parse(grafanaURL)
	if err != nil {
		return nil, errors.Wrap(err, "cannot parse Grafana URL")
	}

	cfg := goapi.DefaultTransportConfig()
	cfg.Host = u.Host
	cfg.BasePath = strings.TrimSuffix(u.Path, "/") + goapi.DefaultBasePath
	if u.Scheme != "" {
		cfg.Schemes = []string{u.Scheme}
	}

	if strings.Contains(auth, ":") {
		parts := strings.SplitN(auth, ":", 2)
		cfg.BasicAuth = url.UserPassword(parts[0], parts[1])
	} else {
		cfg.APIKey = auth
	}

	return goapi.NewHTTPClientWithConfig(nil, cfg), nil
}
