// Code generated by MockGen. DO NOT EDIT.
// Source: datastore.go
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/datastore.go -source datastore.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	datastore "github.com/stackrox/rox/central/complianceoperator/v2/rules/datastore"
	v1 "github.com/stackrox/rox/generated/api/v1"
	storage "github.com/stackrox/rox/generated/storage"
	gomock "go.uber.org/mock/gomock"
)

// MockDataStore is a mock of DataStore interface.
type MockDataStore struct {
	ctrl     *gomock.Controller
	recorder *MockDataStoreMockRecorder
	isgomock struct{}
}

// MockDataStoreMockRecorder is the mock recorder for MockDataStore.
type MockDataStoreMockRecorder struct {
	mock *MockDataStore
}

// NewMockDataStore creates a new mock instance.
func NewMockDataStore(ctrl *gomock.Controller) *MockDataStore {
	mock := &MockDataStore{ctrl: ctrl}
	mock.recorder = &MockDataStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataStore) EXPECT() *MockDataStoreMockRecorder {
	return m.recorder
}

// DeleteRule mocks base method.
func (m *MockDataStore) DeleteRule(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRule", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRule indicates an expected call of DeleteRule.
func (mr *MockDataStoreMockRecorder) DeleteRule(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRule", reflect.TypeOf((*MockDataStore)(nil).DeleteRule), ctx, id)
}

// DeleteRulesByCluster mocks base method.
func (m *MockDataStore) DeleteRulesByCluster(ctx context.Context, clusterID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRulesByCluster", ctx, clusterID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRulesByCluster indicates an expected call of DeleteRulesByCluster.
func (mr *MockDataStoreMockRecorder) DeleteRulesByCluster(ctx, clusterID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRulesByCluster", reflect.TypeOf((*MockDataStore)(nil).DeleteRulesByCluster), ctx, clusterID)
}

// GetControlsByRulesAndBenchmarks mocks base method.
func (m *MockDataStore) GetControlsByRulesAndBenchmarks(ctx context.Context, ruleNames, benchmarkNames []string) ([]*datastore.ControlResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetControlsByRulesAndBenchmarks", ctx, ruleNames, benchmarkNames)
	ret0, _ := ret[0].([]*datastore.ControlResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetControlsByRulesAndBenchmarks indicates an expected call of GetControlsByRulesAndBenchmarks.
func (mr *MockDataStoreMockRecorder) GetControlsByRulesAndBenchmarks(ctx, ruleNames, benchmarkNames any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetControlsByRulesAndBenchmarks", reflect.TypeOf((*MockDataStore)(nil).GetControlsByRulesAndBenchmarks), ctx, ruleNames, benchmarkNames)
}

// GetRulesByCluster mocks base method.
func (m *MockDataStore) GetRulesByCluster(ctx context.Context, clusterID string) ([]*storage.ComplianceOperatorRuleV2, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRulesByCluster", ctx, clusterID)
	ret0, _ := ret[0].([]*storage.ComplianceOperatorRuleV2)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRulesByCluster indicates an expected call of GetRulesByCluster.
func (mr *MockDataStoreMockRecorder) GetRulesByCluster(ctx, clusterID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRulesByCluster", reflect.TypeOf((*MockDataStore)(nil).GetRulesByCluster), ctx, clusterID)
}

// SearchRules mocks base method.
func (m *MockDataStore) SearchRules(ctx context.Context, query *v1.Query) ([]*storage.ComplianceOperatorRuleV2, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchRules", ctx, query)
	ret0, _ := ret[0].([]*storage.ComplianceOperatorRuleV2)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchRules indicates an expected call of SearchRules.
func (mr *MockDataStoreMockRecorder) SearchRules(ctx, query any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchRules", reflect.TypeOf((*MockDataStore)(nil).SearchRules), ctx, query)
}

// UpsertRule mocks base method.
func (m *MockDataStore) UpsertRule(ctx context.Context, rule *storage.ComplianceOperatorRuleV2) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertRule", ctx, rule)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertRule indicates an expected call of UpsertRule.
func (mr *MockDataStoreMockRecorder) UpsertRule(ctx, rule any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertRule", reflect.TypeOf((*MockDataStore)(nil).UpsertRule), ctx, rule)
}
