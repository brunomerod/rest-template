// Code generated by MockGen. DO NOT EDIT.
// Source: rest-template/pkg/service/endpoint (interfaces: Endpoint,Request,Response)

// Package endpoint is a generated GoMock package.
package endpoint

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockEndpoint is a mock of Endpoint interface.
type MockEndpoint struct {
	ctrl     *gomock.Controller
	recorder *MockEndpointMockRecorder
}

// MockEndpointMockRecorder is the mock recorder for MockEndpoint.
type MockEndpointMockRecorder struct {
	mock *MockEndpoint
}

// NewMockEndpoint creates a new mock instance.
func NewMockEndpoint(ctrl *gomock.Controller) *MockEndpoint {
	mock := &MockEndpoint{ctrl: ctrl}
	mock.recorder = &MockEndpointMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEndpoint) EXPECT() *MockEndpointMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockEndpoint) Execute(arg0 context.Context, arg1 interface{}) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", arg0, arg1)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockEndpointMockRecorder) Execute(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockEndpoint)(nil).Execute), arg0, arg1)
}

// MockRequest is a mock of Request interface.
type MockRequest struct {
	ctrl     *gomock.Controller
	recorder *MockRequestMockRecorder
}

// MockRequestMockRecorder is the mock recorder for MockRequest.
type MockRequestMockRecorder struct {
	mock *MockRequest
}

// NewMockRequest creates a new mock instance.
func NewMockRequest(ctrl *gomock.Controller) *MockRequest {
	mock := &MockRequest{ctrl: ctrl}
	mock.recorder = &MockRequestMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRequest) EXPECT() *MockRequestMockRecorder {
	return m.recorder
}

// Method mocks base method.
func (m *MockRequest) Method() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Method")
	ret0, _ := ret[0].(string)
	return ret0
}

// Method indicates an expected call of Method.
func (mr *MockRequestMockRecorder) Method() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Method", reflect.TypeOf((*MockRequest)(nil).Method))
}

// MockResponse is a mock of Response interface.
type MockResponse struct {
	ctrl     *gomock.Controller
	recorder *MockResponseMockRecorder
}

// MockResponseMockRecorder is the mock recorder for MockResponse.
type MockResponseMockRecorder struct {
	mock *MockResponse
}

// NewMockResponse creates a new mock instance.
func NewMockResponse(ctrl *gomock.Controller) *MockResponse {
	mock := &MockResponse{ctrl: ctrl}
	mock.recorder = &MockResponseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResponse) EXPECT() *MockResponseMockRecorder {
	return m.recorder
}

// Failed mocks base method.
func (m *MockResponse) Failed() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Failed")
	ret0, _ := ret[0].(error)
	return ret0
}

// Failed indicates an expected call of Failed.
func (mr *MockResponseMockRecorder) Failed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Failed", reflect.TypeOf((*MockResponse)(nil).Failed))
}
