// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/mayankkapoor/go-rest-mux-app/thirdparty (interfaces: Doer)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockDoer is a mock of Doer interface
type MockDoer struct {
	ctrl     *gomock.Controller
	recorder *MockDoerMockRecorder
}

// MockDoerMockRecorder is the mock recorder for MockDoer
type MockDoerMockRecorder struct {
	mock *MockDoer
}

// NewMockDoer creates a new mock instance
func NewMockDoer(ctrl *gomock.Controller) *MockDoer {
	mock := &MockDoer{ctrl: ctrl}
	mock.recorder = &MockDoerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDoer) EXPECT() *MockDoerMockRecorder {
	return m.recorder
}

// GetEmployeeSalary mocks base method
func (m *MockDoer) GetEmployeeSalary(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEmployeeSalary", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetEmployeeSalary indicates an expected call of GetEmployeeSalary
func (mr *MockDoerMockRecorder) GetEmployeeSalary(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmployeeSalary", reflect.TypeOf((*MockDoer)(nil).GetEmployeeSalary), arg0)
}
