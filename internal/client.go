// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package internal is a generated GoMock package.
package internal

import (
	gomock "github.com/golang/mock/gomock"
	fosite "github.com/ory/fosite"
	go_jose_v2 "gopkg.in/square/go-jose.v2"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// GetID mocks base method
func (m *MockClient) GetID() string {
	ret := m.ctrl.Call(m, "GetID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetID indicates an expected call of GetID
func (mr *MockClientMockRecorder) GetID() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetID", reflect.TypeOf((*MockClient)(nil).GetID))
}

// GetHashedSecret mocks base method
func (m *MockClient) GetHashedSecret() []byte {
	ret := m.ctrl.Call(m, "GetHashedSecret")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// GetHashedSecret indicates an expected call of GetHashedSecret
func (mr *MockClientMockRecorder) GetHashedSecret() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHashedSecret", reflect.TypeOf((*MockClient)(nil).GetHashedSecret))
}

// GetRedirectURIs mocks base method
func (m *MockClient) GetRedirectURIs() []string {
	ret := m.ctrl.Call(m, "GetRedirectURIs")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetRedirectURIs indicates an expected call of GetRedirectURIs
func (mr *MockClientMockRecorder) GetRedirectURIs() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRedirectURIs", reflect.TypeOf((*MockClient)(nil).GetRedirectURIs))
}

// GetGrantTypes mocks base method
func (m *MockClient) GetGrantTypes() fosite.Arguments {
	ret := m.ctrl.Call(m, "GetGrantTypes")
	ret0, _ := ret[0].(fosite.Arguments)
	return ret0
}

// GetGrantTypes indicates an expected call of GetGrantTypes
func (mr *MockClientMockRecorder) GetGrantTypes() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGrantTypes", reflect.TypeOf((*MockClient)(nil).GetGrantTypes))
}

// GetResponseTypes mocks base method
func (m *MockClient) GetResponseTypes() fosite.Arguments {
	ret := m.ctrl.Call(m, "GetResponseTypes")
	ret0, _ := ret[0].(fosite.Arguments)
	return ret0
}

// GetResponseTypes indicates an expected call of GetResponseTypes
func (mr *MockClientMockRecorder) GetResponseTypes() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResponseTypes", reflect.TypeOf((*MockClient)(nil).GetResponseTypes))
}

// GetScopes mocks base method
func (m *MockClient) GetScopes() fosite.Arguments {
	ret := m.ctrl.Call(m, "GetScopes")
	ret0, _ := ret[0].(fosite.Arguments)
	return ret0
}

// GetScopes indicates an expected call of GetScopes
func (mr *MockClientMockRecorder) GetScopes() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScopes", reflect.TypeOf((*MockClient)(nil).GetScopes))
}

// IsPublic mocks base method
func (m *MockClient) IsPublic() bool {
	ret := m.ctrl.Call(m, "IsPublic")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsPublic indicates an expected call of IsPublic
func (mr *MockClientMockRecorder) IsPublic() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsPublic", reflect.TypeOf((*MockClient)(nil).IsPublic))
}

// GetAudience mocks base method
func (m *MockClient) GetAudience() fosite.Arguments {
	ret := m.ctrl.Call(m, "GetAudience")
	ret0, _ := ret[0].(fosite.Arguments)
	return ret0
}

// GetAudience indicates an expected call of GetAudience
func (mr *MockClientMockRecorder) GetAudience() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAudience", reflect.TypeOf((*MockClient)(nil).GetAudience))
}

// GetAccessTokenExpiresAt mocks base method
func (m *MockClient) GetAccessTokenExpiresAt() int64 {
	ret := m.ctrl.Call(m, "GetAccessTokenExpiresAt")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetAccessTokenExpiresAt indicates an expected call of GetAccessTokenExpiresAt
func (mr *MockClientMockRecorder) GetAccessTokenExpiresAt() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessTokenExpiresAt", reflect.TypeOf((*MockClient)(nil).GetAccessTokenExpiresAt))
}

// MockOpenIDConnectClient is a mock of OpenIDConnectClient interface
type MockOpenIDConnectClient struct {
	ctrl     *gomock.Controller
	recorder *MockOpenIDConnectClientMockRecorder
}

// MockOpenIDConnectClientMockRecorder is the mock recorder for MockOpenIDConnectClient
type MockOpenIDConnectClientMockRecorder struct {
	mock *MockOpenIDConnectClient
}

// NewMockOpenIDConnectClient creates a new mock instance
func NewMockOpenIDConnectClient(ctrl *gomock.Controller) *MockOpenIDConnectClient {
	mock := &MockOpenIDConnectClient{ctrl: ctrl}
	mock.recorder = &MockOpenIDConnectClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOpenIDConnectClient) EXPECT() *MockOpenIDConnectClientMockRecorder {
	return m.recorder
}

// GetRequestURIs mocks base method
func (m *MockOpenIDConnectClient) GetRequestURIs() []string {
	ret := m.ctrl.Call(m, "GetRequestURIs")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetRequestURIs indicates an expected call of GetRequestURIs
func (mr *MockOpenIDConnectClientMockRecorder) GetRequestURIs() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequestURIs", reflect.TypeOf((*MockOpenIDConnectClient)(nil).GetRequestURIs))
}

// GetJSONWebKeys mocks base method
func (m *MockOpenIDConnectClient) GetJSONWebKeys() *go_jose_v2.JSONWebKeySet {
	ret := m.ctrl.Call(m, "GetJSONWebKeys")
	ret0, _ := ret[0].(*go_jose_v2.JSONWebKeySet)
	return ret0
}

// GetJSONWebKeys indicates an expected call of GetJSONWebKeys
func (mr *MockOpenIDConnectClientMockRecorder) GetJSONWebKeys() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJSONWebKeys", reflect.TypeOf((*MockOpenIDConnectClient)(nil).GetJSONWebKeys))
}

// GetJSONWebKeysURI mocks base method
func (m *MockOpenIDConnectClient) GetJSONWebKeysURI() string {
	ret := m.ctrl.Call(m, "GetJSONWebKeysURI")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetJSONWebKeysURI indicates an expected call of GetJSONWebKeysURI
func (mr *MockOpenIDConnectClientMockRecorder) GetJSONWebKeysURI() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJSONWebKeysURI", reflect.TypeOf((*MockOpenIDConnectClient)(nil).GetJSONWebKeysURI))
}

// GetRequestObjectSigningAlgorithm mocks base method
func (m *MockOpenIDConnectClient) GetRequestObjectSigningAlgorithm() string {
	ret := m.ctrl.Call(m, "GetRequestObjectSigningAlgorithm")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetRequestObjectSigningAlgorithm indicates an expected call of GetRequestObjectSigningAlgorithm
func (mr *MockOpenIDConnectClientMockRecorder) GetRequestObjectSigningAlgorithm() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequestObjectSigningAlgorithm", reflect.TypeOf((*MockOpenIDConnectClient)(nil).GetRequestObjectSigningAlgorithm))
}

// GetTokenEndpointAuthMethod mocks base method
func (m *MockOpenIDConnectClient) GetTokenEndpointAuthMethod() string {
	ret := m.ctrl.Call(m, "GetTokenEndpointAuthMethod")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetTokenEndpointAuthMethod indicates an expected call of GetTokenEndpointAuthMethod
func (mr *MockOpenIDConnectClientMockRecorder) GetTokenEndpointAuthMethod() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTokenEndpointAuthMethod", reflect.TypeOf((*MockOpenIDConnectClient)(nil).GetTokenEndpointAuthMethod))
}

// GetTokenEndpointAuthSigningAlgorithm mocks base method
func (m *MockOpenIDConnectClient) GetTokenEndpointAuthSigningAlgorithm() string {
	ret := m.ctrl.Call(m, "GetTokenEndpointAuthSigningAlgorithm")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetTokenEndpointAuthSigningAlgorithm indicates an expected call of GetTokenEndpointAuthSigningAlgorithm
func (mr *MockOpenIDConnectClientMockRecorder) GetTokenEndpointAuthSigningAlgorithm() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTokenEndpointAuthSigningAlgorithm", reflect.TypeOf((*MockOpenIDConnectClient)(nil).GetTokenEndpointAuthSigningAlgorithm))
}
