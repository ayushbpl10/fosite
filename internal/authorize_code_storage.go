// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ory/fosite/handler/oauth2 (interfaces: AuthorizeCodeStorage)

// Package internal is a generated GoMock package.
package internal

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	fosite "github.com/ory/fosite"
)

// MockAuthorizeCodeStorage is a mock of AuthorizeCodeStorage interface
type MockAuthorizeCodeStorage struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizeCodeStorageMockRecorder
}

// MockAuthorizeCodeStorageMockRecorder is the mock recorder for MockAuthorizeCodeStorage
type MockAuthorizeCodeStorageMockRecorder struct {
	mock *MockAuthorizeCodeStorage
}

// NewMockAuthorizeCodeStorage creates a new mock instance
func NewMockAuthorizeCodeStorage(ctrl *gomock.Controller) *MockAuthorizeCodeStorage {
	mock := &MockAuthorizeCodeStorage{ctrl: ctrl}
	mock.recorder = &MockAuthorizeCodeStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthorizeCodeStorage) EXPECT() *MockAuthorizeCodeStorageMockRecorder {
	return m.recorder
}

// CreateAuthorizeCodeSession mocks base method
func (m *MockAuthorizeCodeStorage) CreateAuthorizeCodeSession(arg0 context.Context, arg1 string, arg2 fosite.Requester) error {
	ret := m.ctrl.Call(m, "CreateAuthorizeCodeSession", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAuthorizeCodeSession indicates an expected call of CreateAuthorizeCodeSession
func (mr *MockAuthorizeCodeStorageMockRecorder) CreateAuthorizeCodeSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAuthorizeCodeSession", reflect.TypeOf((*MockAuthorizeCodeStorage)(nil).CreateAuthorizeCodeSession), arg0, arg1, arg2)
}

// GetAuthorizeCodeSession mocks base method
func (m *MockAuthorizeCodeStorage) GetAuthorizeCodeSession(arg0 context.Context, arg1 string, arg2 fosite.Session) (fosite.Requester, error) {
	ret := m.ctrl.Call(m, "GetAuthorizeCodeSession", arg0, arg1, arg2)
	ret0, _ := ret[0].(fosite.Requester)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAuthorizeCodeSession indicates an expected call of GetAuthorizeCodeSession
func (mr *MockAuthorizeCodeStorageMockRecorder) GetAuthorizeCodeSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthorizeCodeSession", reflect.TypeOf((*MockAuthorizeCodeStorage)(nil).GetAuthorizeCodeSession), arg0, arg1, arg2)
}

// InvalidateAuthorizeCodeSession mocks base method
func (m *MockAuthorizeCodeStorage) InvalidateAuthorizeCodeSession(arg0 context.Context, arg1 string) error {
	ret := m.ctrl.Call(m, "InvalidateAuthorizeCodeSession", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InvalidateAuthorizeCodeSession indicates an expected call of InvalidateAuthorizeCodeSession
func (mr *MockAuthorizeCodeStorageMockRecorder) InvalidateAuthorizeCodeSession(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvalidateAuthorizeCodeSession", reflect.TypeOf((*MockAuthorizeCodeStorage)(nil).InvalidateAuthorizeCodeSession), arg0, arg1)
}
