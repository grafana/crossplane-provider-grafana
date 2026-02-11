// Package schemeregistry provides thread-safe registration of API group schemes.
// This package ensures each API group is registered with the manager's scheme
// exactly once, even when multiple controllers for the same API group are
// activated concurrently.
package schemeregistry

import (
	"sync"

	ctrl "sigs.k8s.io/controller-runtime"
)

var (
	// registeredGroups tracks which API groups have been registered with the scheme.
	// Key: API group name (e.g., "alerting", "asserts")
	// Value: true if registered
	registeredGroups sync.Map
)

// RegisterClusterAPIGroup registers the scheme for a cluster-scoped API group if not already registered.
// This function is safe to call concurrently from multiple goroutines.
// It ensures each API group is registered exactly once with the manager's scheme.
func RegisterClusterAPIGroup(mgr ctrl.Manager, groupName string) error {
	// Check if already registered
	if _, loaded := registeredGroups.LoadOrStore(groupName, true); loaded {
		// Already registered, skip
		return nil
	}

	// Get the scheme registration function for this group
	schemeFunc := GetClusterGroupSchemeFunc(groupName)
	if schemeFunc == nil {
		// No scheme function found for this group - this shouldn't happen
		// but we'll log and continue rather than fail
		mgr.GetLogger().Info("No scheme registration function found for API group", "group", groupName)
		return nil
	}

	// Register the API group scheme
	scheme := mgr.GetScheme()
	if err := schemeFunc(scheme); err != nil {
		// If registration fails, remove from map so it can be retried
		registeredGroups.Delete(groupName)
		return err
	}

	// Log successful registration
	mgr.GetLogger().Info("Registered API group scheme", "group", groupName)
	return nil
}

// RegisterNamespacedAPIGroup registers the scheme for a namespaced API group if not already registered.
// This function is safe to call concurrently from multiple goroutines.
// It ensures each API group is registered exactly once with the manager's scheme.
func RegisterNamespacedAPIGroup(mgr ctrl.Manager, groupName string) error {
	// Check if already registered
	if _, loaded := registeredGroups.LoadOrStore(groupName, true); loaded {
		// Already registered, skip
		return nil
	}

	// Get the scheme registration function for this group
	schemeFunc := GetNamespacedGroupSchemeFunc(groupName)
	if schemeFunc == nil {
		// No scheme function found for this group - this shouldn't happen
		// but we'll log and continue rather than fail
		mgr.GetLogger().Info("No scheme registration function found for API group", "group", groupName)
		return nil
	}

	// Register the API group scheme
	scheme := mgr.GetScheme()
	if err := schemeFunc(scheme); err != nil {
		// If registration fails, remove from map so it can be retried
		registeredGroups.Delete(groupName)
		return err
	}

	// Log successful registration
	mgr.GetLogger().Info("Registered API group scheme", "group", groupName)
	return nil
}

// IsRegistered returns true if the API group has been registered.
// This is primarily useful for testing.
func IsRegistered(groupName string) bool {
	_, ok := registeredGroups.Load(groupName)
	return ok
}

// Reset clears all registered groups.
// This is primarily useful for testing.
func Reset() {
	registeredGroups = sync.Map{}
}

// GetRegisteredGroupCount returns the number of registered API groups.
// This is primarily useful for testing and diagnostics.
func GetRegisteredGroupCount() int {
	count := 0
	registeredGroups.Range(func(key, value interface{}) bool {
		count++
		return true
	})
	return count
}
