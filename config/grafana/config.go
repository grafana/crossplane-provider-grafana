/*
Copyright 2021 Upbound Inc.
*/

package grafana

import (
	ujconfig "github.com/upbound/upjet/pkg/config"
)

// Configure configures the grafana group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("grafana_dashboard", func(r *ujconfig.Resource) {
	})
}
