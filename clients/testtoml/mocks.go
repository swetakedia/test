package testtoml

import "github.com/stretchr/testify/mock"

// MockClient is a mockable testtoml client.
type MockClient struct {
	mock.Mock
}

// GetTestToml is a mocking a method
func (m *MockClient) GetTestToml(domain string) (*Response, error) {
	a := m.Called(domain)
	return a.Get(0).(*Response), a.Error(1)
}

// GetTestTomlByAddress is a mocking a method
func (m *MockClient) GetTestTomlByAddress(address string) (*Response, error) {
	a := m.Called(address)
	return a.Get(0).(*Response), a.Error(1)
}
