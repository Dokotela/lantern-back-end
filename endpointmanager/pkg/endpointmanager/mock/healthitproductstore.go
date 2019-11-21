package mock

import (
	"github.com/onc-healthit/lantern-back-end/endpointmanager/pkg/endpointmanager"
)

// GetHealthITProduct mocks endpointmanager.HealthITProductStore.GetHealthITProduct and sets s.GetHealthITProductInvoked to true and calls s.GetHealthITProductFn with the given arguments.
func (s *Store) GetHealthITProduct(id int) (*endpointmanager.HealthITProduct, error) {
	s.GetHealthITProductInvoked = true
	return s.GetHealthITProductFn(id)
}

// GetHealthITProductUsingNameAndVersion mocks endpointmanager.HealthITProductStore.GetHealthITProductUsingNameAndVersion and sets s.GetHealthITProductUsingNameAndVersionInvoked to true and calls s.GetHealthITProductUsingNameAndVersionFn with the given arguments.
func (s *Store) GetHealthITProductUsingNameAndVersion(name string, version string) (*endpointmanager.HealthITProduct, error) {
	s.GetHealthITProductUsingNameAndVersionInvoked = true
	return s.GetHealthITProductUsingNameAndVersionFn(name, version)
}

// AddHealthITProduct mocks endpointmanager.HealthITProductStore.AddHealthITProduct and sets s.AddHealthITProductInvoked to true and calls s.AddHealthITProductFn with the given arguments.
func (s *Store) AddHealthITProduct(hitp *endpointmanager.HealthITProduct) error {
	s.AddHealthITProductInvoked = true
	return s.AddHealthITProductFn(hitp)
}

// UpdateHealthITProduct mocks endpointmanager.HealthITProductStore.UpdateHealthITProduct and sets s.UpdateHealthITProductInvoked to true and calls s.UpdateHealthITProductFn with the given arguments.
func (s *Store) UpdateHealthITProduct(hitp *endpointmanager.HealthITProduct) error {
	s.UpdateHealthITProductInvoked = true
	return s.UpdateHealthITProductFn(hitp)
}

// DeleteHealthITProduct mocks endpointmanager.HealthITProductStore.DeleteHealthITProduct and sets s.DeleteHealthITProductInvoked to true and calls s.DeleteHealthITProductFn with the given arguments.
func (s *Store) DeleteHealthITProduct(hitp *endpointmanager.HealthITProduct) error {
	s.DeleteHealthITProductInvoked = true
	return s.DeleteHealthITProductFn(hitp)
}