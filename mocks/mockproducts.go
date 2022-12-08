// Code generated by MockGen. DO NOT EDIT.
// Source: database/database.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	models "sample/git/models"

	gomock "github.com/golang/mock/gomock"
)

// MockDbProducts is a mock of DbProducts interface.
type MockDbProducts struct {
	ctrl     *gomock.Controller
	recorder *MockDbProductsMockRecorder
}

// MockDbProductsMockRecorder is the mock recorder for MockDbProducts.
type MockDbProductsMockRecorder struct {
	mock *MockDbProducts
}

// NewMockDbProducts creates a new mock instance.
func NewMockDbProducts(ctrl *gomock.Controller) *MockDbProducts {
	mock := &MockDbProducts{ctrl: ctrl}
	mock.recorder = &MockDbProductsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDbProducts) EXPECT() *MockDbProductsMockRecorder {
	return m.recorder
}

// GetProducts mocks base method.
func (m *MockDbProducts) GetProducts() ([]models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProducts")
	ret0, _ := ret[0].([]models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProducts indicates an expected call of GetProducts.
func (mr *MockDbProductsMockRecorder) GetProducts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProducts", reflect.TypeOf((*MockDbProducts)(nil).GetProducts))
}