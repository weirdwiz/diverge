package main

import "github.com/stretchr/testify/mock"

// MockStore : the mock store created for unit testing
type MockStore struct {
	mock.Mock
}

// CreateBird Function for the MockStore
func (m *MockStore) CreateBird(bird *Bird) error {
	rets := m.Called(bird)
	return rets.Error(0)
}

// GetBirds function for the mockstore
func (m *MockStore) GetBirds() ([]*Bird, error) {
	rets := m.Called()
	return rets.Get(0).([]*Bird), rets.Error(1)
}

// InitMockStore initialises the MockStore
func InitMockStore() *MockStore {
	s := new(MockStore)
	store = s
	return s
}
