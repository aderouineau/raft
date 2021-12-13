// Code generated by MockGen. DO NOT EDIT.
// Source: internal/raftengine/types.go

// Package raftengine is a generated GoMock package.
package raftengine

import (
	context "context"
	io "io"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	membership "github.com/shaj13/raft/internal/membership"
	storage "github.com/shaj13/raft/internal/storage"
	transport "github.com/shaj13/raft/internal/transport"
	raftlog "github.com/shaj13/raft/raftlog"
	v3 "go.etcd.io/etcd/raft/v3"
)

// MockOperator is a mock of Operator interface.
type MockOperator struct {
	ctrl     *gomock.Controller
	recorder *MockOperatorMockRecorder
}

// MockOperatorMockRecorder is the mock recorder for MockOperator.
type MockOperatorMockRecorder struct {
	mock *MockOperator
}

// NewMockOperator creates a new mock instance.
func NewMockOperator(ctrl *gomock.Controller) *MockOperator {
	mock := &MockOperator{ctrl: ctrl}
	mock.recorder = &MockOperatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOperator) EXPECT() *MockOperatorMockRecorder {
	return m.recorder
}

// String mocks base method.
func (m *MockOperator) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String.
func (mr *MockOperatorMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockOperator)(nil).String))
}

// after mocks base method.
func (m *MockOperator) after(ost *operatorsState) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "after", ost)
	ret0, _ := ret[0].(error)
	return ret0
}

// after indicates an expected call of after.
func (mr *MockOperatorMockRecorder) after(ost interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "after", reflect.TypeOf((*MockOperator)(nil).after), ost)
}

// before mocks base method.
func (m *MockOperator) before(ost *operatorsState) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "before", ost)
	ret0, _ := ret[0].(error)
	return ret0
}

// before indicates an expected call of before.
func (mr *MockOperatorMockRecorder) before(ost interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "before", reflect.TypeOf((*MockOperator)(nil).before), ost)
}

// MockConfig is a mock of Config interface.
type MockConfig struct {
	ctrl     *gomock.Controller
	recorder *MockConfigMockRecorder
}

// MockConfigMockRecorder is the mock recorder for MockConfig.
type MockConfigMockRecorder struct {
	mock *MockConfig
}

// NewMockConfig creates a new mock instance.
func NewMockConfig(ctrl *gomock.Controller) *MockConfig {
	mock := &MockConfig{ctrl: ctrl}
	mock.recorder = &MockConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfig) EXPECT() *MockConfigMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockConfig) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockConfigMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockConfig)(nil).Context))
}

// Dial mocks base method.
func (m *MockConfig) Dial() transport.Dial {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Dial")
	ret0, _ := ret[0].(transport.Dial)
	return ret0
}

// Dial indicates an expected call of Dial.
func (mr *MockConfigMockRecorder) Dial() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dial", reflect.TypeOf((*MockConfig)(nil).Dial))
}

// DrainTimeout mocks base method.
func (m *MockConfig) DrainTimeout() time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DrainTimeout")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// DrainTimeout indicates an expected call of DrainTimeout.
func (mr *MockConfigMockRecorder) DrainTimeout() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DrainTimeout", reflect.TypeOf((*MockConfig)(nil).DrainTimeout))
}

// GroupID mocks base method.
func (m *MockConfig) GroupID() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GroupID")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GroupID indicates an expected call of GroupID.
func (mr *MockConfigMockRecorder) GroupID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GroupID", reflect.TypeOf((*MockConfig)(nil).GroupID))
}

// Logger mocks base method.
func (m *MockConfig) Logger() raftlog.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logger")
	ret0, _ := ret[0].(raftlog.Logger)
	return ret0
}

// Logger indicates an expected call of Logger.
func (mr *MockConfigMockRecorder) Logger() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logger", reflect.TypeOf((*MockConfig)(nil).Logger))
}

// Mux mocks base method.
func (m *MockConfig) Mux() Mux {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Mux")
	ret0, _ := ret[0].(Mux)
	return ret0
}

// Mux indicates an expected call of Mux.
func (mr *MockConfigMockRecorder) Mux() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mux", reflect.TypeOf((*MockConfig)(nil).Mux))
}

// Pool mocks base method.
func (m *MockConfig) Pool() membership.Pool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pool")
	ret0, _ := ret[0].(membership.Pool)
	return ret0
}

// Pool indicates an expected call of Pool.
func (mr *MockConfigMockRecorder) Pool() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pool", reflect.TypeOf((*MockConfig)(nil).Pool))
}

// RaftConfig mocks base method.
func (m *MockConfig) RaftConfig() *v3.Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RaftConfig")
	ret0, _ := ret[0].(*v3.Config)
	return ret0
}

// RaftConfig indicates an expected call of RaftConfig.
func (mr *MockConfigMockRecorder) RaftConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RaftConfig", reflect.TypeOf((*MockConfig)(nil).RaftConfig))
}

// SnapInterval mocks base method.
func (m *MockConfig) SnapInterval() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SnapInterval")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// SnapInterval indicates an expected call of SnapInterval.
func (mr *MockConfigMockRecorder) SnapInterval() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SnapInterval", reflect.TypeOf((*MockConfig)(nil).SnapInterval))
}

// StateMachine mocks base method.
func (m *MockConfig) StateMachine() StateMachine {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateMachine")
	ret0, _ := ret[0].(StateMachine)
	return ret0
}

// StateMachine indicates an expected call of StateMachine.
func (mr *MockConfigMockRecorder) StateMachine() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateMachine", reflect.TypeOf((*MockConfig)(nil).StateMachine))
}

// Storage mocks base method.
func (m *MockConfig) Storage() storage.Storage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Storage")
	ret0, _ := ret[0].(storage.Storage)
	return ret0
}

// Storage indicates an expected call of Storage.
func (mr *MockConfigMockRecorder) Storage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Storage", reflect.TypeOf((*MockConfig)(nil).Storage))
}

// TickInterval mocks base method.
func (m *MockConfig) TickInterval() time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TickInterval")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// TickInterval indicates an expected call of TickInterval.
func (mr *MockConfigMockRecorder) TickInterval() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TickInterval", reflect.TypeOf((*MockConfig)(nil).TickInterval))
}

// MockStateMachine is a mock of StateMachine interface.
type MockStateMachine struct {
	ctrl     *gomock.Controller
	recorder *MockStateMachineMockRecorder
}

// MockStateMachineMockRecorder is the mock recorder for MockStateMachine.
type MockStateMachineMockRecorder struct {
	mock *MockStateMachine
}

// NewMockStateMachine creates a new mock instance.
func NewMockStateMachine(ctrl *gomock.Controller) *MockStateMachine {
	mock := &MockStateMachine{ctrl: ctrl}
	mock.recorder = &MockStateMachineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStateMachine) EXPECT() *MockStateMachineMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockStateMachine) Apply(arg0 []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Apply", arg0)
}

// Apply indicates an expected call of Apply.
func (mr *MockStateMachineMockRecorder) Apply(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockStateMachine)(nil).Apply), arg0)
}

// Restore mocks base method.
func (m *MockStateMachine) Restore(arg0 io.ReadCloser) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Restore", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Restore indicates an expected call of Restore.
func (mr *MockStateMachineMockRecorder) Restore(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Restore", reflect.TypeOf((*MockStateMachine)(nil).Restore), arg0)
}

// Snapshot mocks base method.
func (m *MockStateMachine) Snapshot() (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Snapshot")
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Snapshot indicates an expected call of Snapshot.
func (mr *MockStateMachineMockRecorder) Snapshot() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Snapshot", reflect.TypeOf((*MockStateMachine)(nil).Snapshot))
}

// MockMux is a mock of Mux interface.
type MockMux struct {
	ctrl     *gomock.Controller
	recorder *MockMuxMockRecorder
}

// MockMuxMockRecorder is the mock recorder for MockMux.
type MockMuxMockRecorder struct {
	mock *MockMux
}

// NewMockMux creates a new mock instance.
func NewMockMux(ctrl *gomock.Controller) *MockMux {
	mock := &MockMux{ctrl: ctrl}
	mock.recorder = &MockMuxMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMux) EXPECT() *MockMuxMockRecorder {
	return m.recorder
}

// Start mocks base method.
func (m *MockMux) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start.
func (mr *MockMuxMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockMux)(nil).Start))
}

// Stop mocks base method.
func (m *MockMux) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockMuxMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockMux)(nil).Stop))
}

// add mocks base method.
func (m *MockMux) add(gid uint64, rn *v3.RawNode, cfg *v3.Config) v3.Node {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "add", gid, rn, cfg)
	ret0, _ := ret[0].(v3.Node)
	return ret0
}

// add indicates an expected call of add.
func (mr *MockMuxMockRecorder) add(gid, rn, cfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "add", reflect.TypeOf((*MockMux)(nil).add), gid, rn, cfg)
}