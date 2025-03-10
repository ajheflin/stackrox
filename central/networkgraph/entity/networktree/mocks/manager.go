// Code generated by MockGen. DO NOT EDIT.
// Source: manager.go
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/manager.go -source manager.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	storage "github.com/stackrox/rox/generated/storage"
	tree "github.com/stackrox/rox/pkg/networkgraph/tree"
	gomock "go.uber.org/mock/gomock"
)

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
	isgomock struct{}
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// CreateDefaultNetworkTree mocks base method.
func (m *MockManager) CreateDefaultNetworkTree(ctx context.Context) tree.NetworkTree {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDefaultNetworkTree", ctx)
	ret0, _ := ret[0].(tree.NetworkTree)
	return ret0
}

// CreateDefaultNetworkTree indicates an expected call of CreateDefaultNetworkTree.
func (mr *MockManagerMockRecorder) CreateDefaultNetworkTree(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDefaultNetworkTree", reflect.TypeOf((*MockManager)(nil).CreateDefaultNetworkTree), ctx)
}

// CreateNetworkTree mocks base method.
func (m *MockManager) CreateNetworkTree(ctx context.Context, clusterID string) tree.NetworkTree {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNetworkTree", ctx, clusterID)
	ret0, _ := ret[0].(tree.NetworkTree)
	return ret0
}

// CreateNetworkTree indicates an expected call of CreateNetworkTree.
func (mr *MockManagerMockRecorder) CreateNetworkTree(ctx, clusterID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNetworkTree", reflect.TypeOf((*MockManager)(nil).CreateNetworkTree), ctx, clusterID)
}

// DeleteNetworkTree mocks base method.
func (m *MockManager) DeleteNetworkTree(ctx context.Context, clusterID string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteNetworkTree", ctx, clusterID)
}

// DeleteNetworkTree indicates an expected call of DeleteNetworkTree.
func (mr *MockManagerMockRecorder) DeleteNetworkTree(ctx, clusterID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNetworkTree", reflect.TypeOf((*MockManager)(nil).DeleteNetworkTree), ctx, clusterID)
}

// GetDefaultNetworkTree mocks base method.
func (m *MockManager) GetDefaultNetworkTree(ctx context.Context) tree.ReadOnlyNetworkTree {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDefaultNetworkTree", ctx)
	ret0, _ := ret[0].(tree.ReadOnlyNetworkTree)
	return ret0
}

// GetDefaultNetworkTree indicates an expected call of GetDefaultNetworkTree.
func (mr *MockManagerMockRecorder) GetDefaultNetworkTree(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDefaultNetworkTree", reflect.TypeOf((*MockManager)(nil).GetDefaultNetworkTree), ctx)
}

// GetNetworkTree mocks base method.
func (m *MockManager) GetNetworkTree(ctx context.Context, clusterID string) tree.NetworkTree {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNetworkTree", ctx, clusterID)
	ret0, _ := ret[0].(tree.NetworkTree)
	return ret0
}

// GetNetworkTree indicates an expected call of GetNetworkTree.
func (mr *MockManagerMockRecorder) GetNetworkTree(ctx, clusterID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetworkTree", reflect.TypeOf((*MockManager)(nil).GetNetworkTree), ctx, clusterID)
}

// GetReadOnlyNetworkTree mocks base method.
func (m *MockManager) GetReadOnlyNetworkTree(ctx context.Context, clusterID string) tree.ReadOnlyNetworkTree {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReadOnlyNetworkTree", ctx, clusterID)
	ret0, _ := ret[0].(tree.ReadOnlyNetworkTree)
	return ret0
}

// GetReadOnlyNetworkTree indicates an expected call of GetReadOnlyNetworkTree.
func (mr *MockManagerMockRecorder) GetReadOnlyNetworkTree(ctx, clusterID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReadOnlyNetworkTree", reflect.TypeOf((*MockManager)(nil).GetReadOnlyNetworkTree), ctx, clusterID)
}

// Initialize mocks base method.
func (m *MockManager) Initialize(entitiesByCluster map[string][]*storage.NetworkEntityInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Initialize", entitiesByCluster)
	ret0, _ := ret[0].(error)
	return ret0
}

// Initialize indicates an expected call of Initialize.
func (mr *MockManagerMockRecorder) Initialize(entitiesByCluster any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockManager)(nil).Initialize), entitiesByCluster)
}
