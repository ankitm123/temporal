// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go
//
// Generated by this command:
//
//	mockgen -package membership -source interfaces.go -destination interfaces_mock.go
//

// Package membership is a generated GoMock package.
package membership

import (
	context "context"
	reflect "reflect"
	time "time"

	primitives "go.temporal.io/server/common/primitives"
	gomock "go.uber.org/mock/gomock"
)

// MockMonitor is a mock of Monitor interface.
type MockMonitor struct {
	ctrl     *gomock.Controller
	recorder *MockMonitorMockRecorder
	isgomock struct{}
}

// MockMonitorMockRecorder is the mock recorder for MockMonitor.
type MockMonitorMockRecorder struct {
	mock *MockMonitor
}

// NewMockMonitor creates a new mock instance.
func NewMockMonitor(ctrl *gomock.Controller) *MockMonitor {
	mock := &MockMonitor{ctrl: ctrl}
	mock.recorder = &MockMonitorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMonitor) EXPECT() *MockMonitorMockRecorder {
	return m.recorder
}

// ApproximateMaxPropagationTime mocks base method.
func (m *MockMonitor) ApproximateMaxPropagationTime() time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApproximateMaxPropagationTime")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// ApproximateMaxPropagationTime indicates an expected call of ApproximateMaxPropagationTime.
func (mr *MockMonitorMockRecorder) ApproximateMaxPropagationTime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApproximateMaxPropagationTime", reflect.TypeOf((*MockMonitor)(nil).ApproximateMaxPropagationTime))
}

// EvictSelf mocks base method.
func (m *MockMonitor) EvictSelf() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EvictSelf")
	ret0, _ := ret[0].(error)
	return ret0
}

// EvictSelf indicates an expected call of EvictSelf.
func (mr *MockMonitorMockRecorder) EvictSelf() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EvictSelf", reflect.TypeOf((*MockMonitor)(nil).EvictSelf))
}

// EvictSelfAt mocks base method.
func (m *MockMonitor) EvictSelfAt(asOf time.Time) (time.Duration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EvictSelfAt", asOf)
	ret0, _ := ret[0].(time.Duration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EvictSelfAt indicates an expected call of EvictSelfAt.
func (mr *MockMonitorMockRecorder) EvictSelfAt(asOf any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EvictSelfAt", reflect.TypeOf((*MockMonitor)(nil).EvictSelfAt), asOf)
}

// GetReachableMembers mocks base method.
func (m *MockMonitor) GetReachableMembers() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReachableMembers")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReachableMembers indicates an expected call of GetReachableMembers.
func (mr *MockMonitorMockRecorder) GetReachableMembers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReachableMembers", reflect.TypeOf((*MockMonitor)(nil).GetReachableMembers))
}

// GetResolver mocks base method.
func (m *MockMonitor) GetResolver(service primitives.ServiceName) (ServiceResolver, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResolver", service)
	ret0, _ := ret[0].(ServiceResolver)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResolver indicates an expected call of GetResolver.
func (mr *MockMonitorMockRecorder) GetResolver(service any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResolver", reflect.TypeOf((*MockMonitor)(nil).GetResolver), service)
}

// SetDraining mocks base method.
func (m *MockMonitor) SetDraining(draining bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDraining", draining)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDraining indicates an expected call of SetDraining.
func (mr *MockMonitorMockRecorder) SetDraining(draining any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDraining", reflect.TypeOf((*MockMonitor)(nil).SetDraining), draining)
}

// Start mocks base method.
func (m *MockMonitor) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start.
func (mr *MockMonitorMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockMonitor)(nil).Start))
}

// WaitUntilInitialized mocks base method.
func (m *MockMonitor) WaitUntilInitialized(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitUntilInitialized", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitUntilInitialized indicates an expected call of WaitUntilInitialized.
func (mr *MockMonitorMockRecorder) WaitUntilInitialized(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitUntilInitialized", reflect.TypeOf((*MockMonitor)(nil).WaitUntilInitialized), arg0)
}

// MockServiceResolver is a mock of ServiceResolver interface.
type MockServiceResolver struct {
	ctrl     *gomock.Controller
	recorder *MockServiceResolverMockRecorder
	isgomock struct{}
}

// MockServiceResolverMockRecorder is the mock recorder for MockServiceResolver.
type MockServiceResolverMockRecorder struct {
	mock *MockServiceResolver
}

// NewMockServiceResolver creates a new mock instance.
func NewMockServiceResolver(ctrl *gomock.Controller) *MockServiceResolver {
	mock := &MockServiceResolver{ctrl: ctrl}
	mock.recorder = &MockServiceResolverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceResolver) EXPECT() *MockServiceResolverMockRecorder {
	return m.recorder
}

// AddListener mocks base method.
func (m *MockServiceResolver) AddListener(name string, notifyChannel chan<- *ChangedEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddListener", name, notifyChannel)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddListener indicates an expected call of AddListener.
func (mr *MockServiceResolverMockRecorder) AddListener(name, notifyChannel any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddListener", reflect.TypeOf((*MockServiceResolver)(nil).AddListener), name, notifyChannel)
}

// AvailableMemberCount mocks base method.
func (m *MockServiceResolver) AvailableMemberCount() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AvailableMemberCount")
	ret0, _ := ret[0].(int)
	return ret0
}

// AvailableMemberCount indicates an expected call of AvailableMemberCount.
func (mr *MockServiceResolverMockRecorder) AvailableMemberCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AvailableMemberCount", reflect.TypeOf((*MockServiceResolver)(nil).AvailableMemberCount))
}

// AvailableMembers mocks base method.
func (m *MockServiceResolver) AvailableMembers() []HostInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AvailableMembers")
	ret0, _ := ret[0].([]HostInfo)
	return ret0
}

// AvailableMembers indicates an expected call of AvailableMembers.
func (mr *MockServiceResolverMockRecorder) AvailableMembers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AvailableMembers", reflect.TypeOf((*MockServiceResolver)(nil).AvailableMembers))
}

// Lookup mocks base method.
func (m *MockServiceResolver) Lookup(key string) (HostInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Lookup", key)
	ret0, _ := ret[0].(HostInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Lookup indicates an expected call of Lookup.
func (mr *MockServiceResolverMockRecorder) Lookup(key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lookup", reflect.TypeOf((*MockServiceResolver)(nil).Lookup), key)
}

// LookupN mocks base method.
func (m *MockServiceResolver) LookupN(key string, n int) []HostInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LookupN", key, n)
	ret0, _ := ret[0].([]HostInfo)
	return ret0
}

// LookupN indicates an expected call of LookupN.
func (mr *MockServiceResolverMockRecorder) LookupN(key, n any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LookupN", reflect.TypeOf((*MockServiceResolver)(nil).LookupN), key, n)
}

// MemberCount mocks base method.
func (m *MockServiceResolver) MemberCount() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MemberCount")
	ret0, _ := ret[0].(int)
	return ret0
}

// MemberCount indicates an expected call of MemberCount.
func (mr *MockServiceResolverMockRecorder) MemberCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MemberCount", reflect.TypeOf((*MockServiceResolver)(nil).MemberCount))
}

// Members mocks base method.
func (m *MockServiceResolver) Members() []HostInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Members")
	ret0, _ := ret[0].([]HostInfo)
	return ret0
}

// Members indicates an expected call of Members.
func (mr *MockServiceResolverMockRecorder) Members() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Members", reflect.TypeOf((*MockServiceResolver)(nil).Members))
}

// RemoveListener mocks base method.
func (m *MockServiceResolver) RemoveListener(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveListener", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveListener indicates an expected call of RemoveListener.
func (mr *MockServiceResolverMockRecorder) RemoveListener(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveListener", reflect.TypeOf((*MockServiceResolver)(nil).RemoveListener), name)
}

// RequestRefresh mocks base method.
func (m *MockServiceResolver) RequestRefresh() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RequestRefresh")
}

// RequestRefresh indicates an expected call of RequestRefresh.
func (mr *MockServiceResolverMockRecorder) RequestRefresh() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestRefresh", reflect.TypeOf((*MockServiceResolver)(nil).RequestRefresh))
}

// MockHostInfoProvider is a mock of HostInfoProvider interface.
type MockHostInfoProvider struct {
	ctrl     *gomock.Controller
	recorder *MockHostInfoProviderMockRecorder
	isgomock struct{}
}

// MockHostInfoProviderMockRecorder is the mock recorder for MockHostInfoProvider.
type MockHostInfoProviderMockRecorder struct {
	mock *MockHostInfoProvider
}

// NewMockHostInfoProvider creates a new mock instance.
func NewMockHostInfoProvider(ctrl *gomock.Controller) *MockHostInfoProvider {
	mock := &MockHostInfoProvider{ctrl: ctrl}
	mock.recorder = &MockHostInfoProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHostInfoProvider) EXPECT() *MockHostInfoProviderMockRecorder {
	return m.recorder
}

// HostInfo mocks base method.
func (m *MockHostInfoProvider) HostInfo() HostInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HostInfo")
	ret0, _ := ret[0].(HostInfo)
	return ret0
}

// HostInfo indicates an expected call of HostInfo.
func (mr *MockHostInfoProviderMockRecorder) HostInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HostInfo", reflect.TypeOf((*MockHostInfoProvider)(nil).HostInfo))
}
