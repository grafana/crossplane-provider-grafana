/*
 Copyright 2022 Upbound Inc
*/

package features

import "github.com/crossplane/crossplane-runtime/v2/pkg/feature"

// Feature flags.
const (

	// EnableBetaManagementPolicies enables beta support for
	// Management Policies. See the below design for more details.
	// https://github.com/crossplane/crossplane/pull/3531
	EnableBetaManagementPolicies feature.Flag = "EnableBetaManagementPolicies"

	// EnableSafeStart gates controller startup behind a phased rollout
	// mechanism. When enabled, controllers use their SetupGated variants
	// allowing for feature-flag based activation. This supports safer
	// upgrades and gradual enablement of new API groups.
	EnableSafeStart feature.Flag = "EnableSafeStart"
)
