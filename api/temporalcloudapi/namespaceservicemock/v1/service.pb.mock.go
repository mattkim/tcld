// Code generated by MockGen. DO NOT EDIT.
// Source: namespaceservice/v1/service.pb.go

// Package namespaceservicemock is a generated GoMock package.
package namespaceservicemock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	namespaceservice "github.com/temporalio/tcld/api/temporalcloudapi/namespaceservice/v1"
	grpc "google.golang.org/grpc"
)

// MockNamespaceServiceClient is a mock of NamespaceServiceClient interface.
type MockNamespaceServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockNamespaceServiceClientMockRecorder
}

// MockNamespaceServiceClientMockRecorder is the mock recorder for MockNamespaceServiceClient.
type MockNamespaceServiceClientMockRecorder struct {
	mock *MockNamespaceServiceClient
}

// NewMockNamespaceServiceClient creates a new mock instance.
func NewMockNamespaceServiceClient(ctrl *gomock.Controller) *MockNamespaceServiceClient {
	mock := &MockNamespaceServiceClient{ctrl: ctrl}
	mock.recorder = &MockNamespaceServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNamespaceServiceClient) EXPECT() *MockNamespaceServiceClientMockRecorder {
	return m.recorder
}

// GetNamespace mocks base method.
func (m *MockNamespaceServiceClient) GetNamespace(ctx context.Context, in *namespaceservice.GetNamespaceRequest, opts ...grpc.CallOption) (*namespaceservice.GetNamespaceResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetNamespace", varargs...)
	ret0, _ := ret[0].(*namespaceservice.GetNamespaceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNamespace indicates an expected call of GetNamespace.
func (mr *MockNamespaceServiceClientMockRecorder) GetNamespace(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamespace", reflect.TypeOf((*MockNamespaceServiceClient)(nil).GetNamespace), varargs...)
}

// ListNamespaces mocks base method.
func (m *MockNamespaceServiceClient) ListNamespaces(ctx context.Context, in *namespaceservice.ListNamespacesRequest, opts ...grpc.CallOption) (*namespaceservice.ListNamespacesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListNamespaces", varargs...)
	ret0, _ := ret[0].(*namespaceservice.ListNamespacesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNamespaces indicates an expected call of ListNamespaces.
func (mr *MockNamespaceServiceClientMockRecorder) ListNamespaces(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNamespaces", reflect.TypeOf((*MockNamespaceServiceClient)(nil).ListNamespaces), varargs...)
}

// RegisterNamespace mocks base method.
func (m *MockNamespaceServiceClient) RegisterNamespace(ctx context.Context, in *namespaceservice.RegisterNamespaceRequest, opts ...grpc.CallOption) (*namespaceservice.RegisterNamespaceResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RegisterNamespace", varargs...)
	ret0, _ := ret[0].(*namespaceservice.RegisterNamespaceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterNamespace indicates an expected call of RegisterNamespace.
func (mr *MockNamespaceServiceClientMockRecorder) RegisterNamespace(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterNamespace", reflect.TypeOf((*MockNamespaceServiceClient)(nil).RegisterNamespace), varargs...)
}

// RenameCustomSearchAttribute mocks base method.
func (m *MockNamespaceServiceClient) RenameCustomSearchAttribute(ctx context.Context, in *namespaceservice.RenameCustomSearchAttributeRequest, opts ...grpc.CallOption) (*namespaceservice.RenameCustomSearchAttributeResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RenameCustomSearchAttribute", varargs...)
	ret0, _ := ret[0].(*namespaceservice.RenameCustomSearchAttributeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RenameCustomSearchAttribute indicates an expected call of RenameCustomSearchAttribute.
func (mr *MockNamespaceServiceClientMockRecorder) RenameCustomSearchAttribute(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenameCustomSearchAttribute", reflect.TypeOf((*MockNamespaceServiceClient)(nil).RenameCustomSearchAttribute), varargs...)
}

// UpdateNamespace mocks base method.
func (m *MockNamespaceServiceClient) UpdateNamespace(ctx context.Context, in *namespaceservice.UpdateNamespaceRequest, opts ...grpc.CallOption) (*namespaceservice.UpdateNamespaceResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateNamespace", varargs...)
	ret0, _ := ret[0].(*namespaceservice.UpdateNamespaceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateNamespace indicates an expected call of UpdateNamespace.
func (mr *MockNamespaceServiceClientMockRecorder) UpdateNamespace(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNamespace", reflect.TypeOf((*MockNamespaceServiceClient)(nil).UpdateNamespace), varargs...)
}

// MockNamespaceServiceServer is a mock of NamespaceServiceServer interface.
type MockNamespaceServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockNamespaceServiceServerMockRecorder
}

// MockNamespaceServiceServerMockRecorder is the mock recorder for MockNamespaceServiceServer.
type MockNamespaceServiceServerMockRecorder struct {
	mock *MockNamespaceServiceServer
}

// NewMockNamespaceServiceServer creates a new mock instance.
func NewMockNamespaceServiceServer(ctrl *gomock.Controller) *MockNamespaceServiceServer {
	mock := &MockNamespaceServiceServer{ctrl: ctrl}
	mock.recorder = &MockNamespaceServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNamespaceServiceServer) EXPECT() *MockNamespaceServiceServerMockRecorder {
	return m.recorder
}

// GetNamespace mocks base method.
func (m *MockNamespaceServiceServer) GetNamespace(arg0 context.Context, arg1 *namespaceservice.GetNamespaceRequest) (*namespaceservice.GetNamespaceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNamespace", arg0, arg1)
	ret0, _ := ret[0].(*namespaceservice.GetNamespaceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNamespace indicates an expected call of GetNamespace.
func (mr *MockNamespaceServiceServerMockRecorder) GetNamespace(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamespace", reflect.TypeOf((*MockNamespaceServiceServer)(nil).GetNamespace), arg0, arg1)
}

// ListNamespaces mocks base method.
func (m *MockNamespaceServiceServer) ListNamespaces(arg0 context.Context, arg1 *namespaceservice.ListNamespacesRequest) (*namespaceservice.ListNamespacesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNamespaces", arg0, arg1)
	ret0, _ := ret[0].(*namespaceservice.ListNamespacesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNamespaces indicates an expected call of ListNamespaces.
func (mr *MockNamespaceServiceServerMockRecorder) ListNamespaces(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNamespaces", reflect.TypeOf((*MockNamespaceServiceServer)(nil).ListNamespaces), arg0, arg1)
}

// RegisterNamespace mocks base method.
func (m *MockNamespaceServiceServer) RegisterNamespace(arg0 context.Context, arg1 *namespaceservice.RegisterNamespaceRequest) (*namespaceservice.RegisterNamespaceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterNamespace", arg0, arg1)
	ret0, _ := ret[0].(*namespaceservice.RegisterNamespaceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterNamespace indicates an expected call of RegisterNamespace.
func (mr *MockNamespaceServiceServerMockRecorder) RegisterNamespace(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterNamespace", reflect.TypeOf((*MockNamespaceServiceServer)(nil).RegisterNamespace), arg0, arg1)
}

// RenameCustomSearchAttribute mocks base method.
func (m *MockNamespaceServiceServer) RenameCustomSearchAttribute(arg0 context.Context, arg1 *namespaceservice.RenameCustomSearchAttributeRequest) (*namespaceservice.RenameCustomSearchAttributeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenameCustomSearchAttribute", arg0, arg1)
	ret0, _ := ret[0].(*namespaceservice.RenameCustomSearchAttributeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RenameCustomSearchAttribute indicates an expected call of RenameCustomSearchAttribute.
func (mr *MockNamespaceServiceServerMockRecorder) RenameCustomSearchAttribute(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenameCustomSearchAttribute", reflect.TypeOf((*MockNamespaceServiceServer)(nil).RenameCustomSearchAttribute), arg0, arg1)
}

// UpdateNamespace mocks base method.
func (m *MockNamespaceServiceServer) UpdateNamespace(arg0 context.Context, arg1 *namespaceservice.UpdateNamespaceRequest) (*namespaceservice.UpdateNamespaceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateNamespace", arg0, arg1)
	ret0, _ := ret[0].(*namespaceservice.UpdateNamespaceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateNamespace indicates an expected call of UpdateNamespace.
func (mr *MockNamespaceServiceServerMockRecorder) UpdateNamespace(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNamespace", reflect.TypeOf((*MockNamespaceServiceServer)(nil).UpdateNamespace), arg0, arg1)
}
