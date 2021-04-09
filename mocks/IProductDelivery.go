// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// IProductDelivery is an autogenerated mock type for the IProductDelivery type
type IProductDelivery struct {
	mock.Mock
}

// BuyProduct provides a mock function with given fields: _a0, _a1
func (_m *IProductDelivery) BuyProduct(_a0 http.ResponseWriter, _a1 *http.Request) {
	_m.Called(_a0, _a1)
}

// DeleteProduct provides a mock function with given fields: _a0, _a1
func (_m *IProductDelivery) DeleteProduct(_a0 http.ResponseWriter, _a1 *http.Request) {
	_m.Called(_a0, _a1)
}

// GetAvailableProducts provides a mock function with given fields: _a0, _a1
func (_m *IProductDelivery) GetAvailableProducts(_a0 http.ResponseWriter, _a1 *http.Request) {
	_m.Called(_a0, _a1)
}

// GetProduct provides a mock function with given fields: _a0, _a1
func (_m *IProductDelivery) GetProduct(_a0 http.ResponseWriter, _a1 *http.Request) {
	_m.Called(_a0, _a1)
}

// GetSellerProducts provides a mock function with given fields: _a0, _a1
func (_m *IProductDelivery) GetSellerProducts(_a0 http.ResponseWriter, _a1 *http.Request) {
	_m.Called(_a0, _a1)
}

// InsertProduct provides a mock function with given fields: _a0, _a1
func (_m *IProductDelivery) InsertProduct(_a0 http.ResponseWriter, _a1 *http.Request) {
	_m.Called(_a0, _a1)
}

// UpdateProduct provides a mock function with given fields: _a0, _a1
func (_m *IProductDelivery) UpdateProduct(_a0 http.ResponseWriter, _a1 *http.Request) {
	_m.Called(_a0, _a1)
}
