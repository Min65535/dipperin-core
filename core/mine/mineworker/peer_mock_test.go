// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dipperin/dipperin-core/core/chaincommunication (interfaces: PmAbstractPeer)

// Package mineworker is a generated GoMock package.
package mineworker

import (
	common "github.com/dipperin/dipperin-core/common"
	p2p "github.com/dipperin/dipperin-core/third_party/p2p"
	gomock "github.com/golang/mock/gomock"
	net "net"
	reflect "reflect"
)

// MockPmAbstractPeer is a mock of PmAbstractPeer interface
type MockPmAbstractPeer struct {
	ctrl     *gomock.Controller
	recorder *MockPmAbstractPeerMockRecorder
}

// MockPmAbstractPeerMockRecorder is the mock recorder for MockPmAbstractPeer
type MockPmAbstractPeerMockRecorder struct {
	mock *MockPmAbstractPeer
}

// NewMockPmAbstractPeer creates a new mock instance
func NewMockPmAbstractPeer(ctrl *gomock.Controller) *MockPmAbstractPeer {
	mock := &MockPmAbstractPeer{ctrl: ctrl}
	mock.recorder = &MockPmAbstractPeerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPmAbstractPeer) EXPECT() *MockPmAbstractPeerMockRecorder {
	return m.recorder
}

// DisconnectPeer mocks base method
func (m *MockPmAbstractPeer) DisconnectPeer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DisconnectPeer")
}

// DisconnectPeer indicates an expected call of DisconnectPeer
func (mr *MockPmAbstractPeerMockRecorder) DisconnectPeer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisconnectPeer", reflect.TypeOf((*MockPmAbstractPeer)(nil).DisconnectPeer))
}

// GetCsPeerInfo mocks base method
func (m *MockPmAbstractPeer) GetCsPeerInfo() *p2p.CsPeerInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCsPeerInfo")
	ret0, _ := ret[0].(*p2p.CsPeerInfo)
	return ret0
}

// GetCsPeerInfo indicates an expected call of GetCsPeerInfo
func (mr *MockPmAbstractPeerMockRecorder) GetCsPeerInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCsPeerInfo", reflect.TypeOf((*MockPmAbstractPeer)(nil).GetCsPeerInfo))
}

// GetHead mocks base method
func (m *MockPmAbstractPeer) GetHead() (common.Hash, uint64) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHead")
	ret0, _ := ret[0].(common.Hash)
	ret1, _ := ret[1].(uint64)
	return ret0, ret1
}

// GetHead indicates an expected call of GetHead
func (mr *MockPmAbstractPeerMockRecorder) GetHead() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHead", reflect.TypeOf((*MockPmAbstractPeer)(nil).GetHead))
}

// GetPeerRawUrl mocks base method
func (m *MockPmAbstractPeer) GetPeerRawUrl() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPeerRawUrl")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetPeerRawUrl indicates an expected call of GetPeerRawUrl
func (mr *MockPmAbstractPeerMockRecorder) GetPeerRawUrl() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPeerRawUrl", reflect.TypeOf((*MockPmAbstractPeer)(nil).GetPeerRawUrl))
}

// ID mocks base method
func (m *MockPmAbstractPeer) ID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ID indicates an expected call of ID
func (mr *MockPmAbstractPeerMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockPmAbstractPeer)(nil).ID))
}

// IsRunning mocks base method
func (m *MockPmAbstractPeer) IsRunning() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsRunning")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsRunning indicates an expected call of IsRunning
func (mr *MockPmAbstractPeerMockRecorder) IsRunning() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsRunning", reflect.TypeOf((*MockPmAbstractPeer)(nil).IsRunning))
}

// NodeName mocks base method
func (m *MockPmAbstractPeer) NodeName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodeName")
	ret0, _ := ret[0].(string)
	return ret0
}

// NodeName indicates an expected call of NodeName
func (mr *MockPmAbstractPeerMockRecorder) NodeName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeName", reflect.TypeOf((*MockPmAbstractPeer)(nil).NodeName))
}

// NodeType mocks base method
func (m *MockPmAbstractPeer) NodeType() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodeType")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// NodeType indicates an expected call of NodeType
func (mr *MockPmAbstractPeerMockRecorder) NodeType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeType", reflect.TypeOf((*MockPmAbstractPeer)(nil).NodeType))
}

// ReadMsg mocks base method
func (m *MockPmAbstractPeer) ReadMsg() (p2p.Msg, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadMsg")
	ret0, _ := ret[0].(p2p.Msg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadMsg indicates an expected call of ReadMsg
func (mr *MockPmAbstractPeerMockRecorder) ReadMsg() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadMsg", reflect.TypeOf((*MockPmAbstractPeer)(nil).ReadMsg))
}

// RemoteAddress mocks base method
func (m *MockPmAbstractPeer) RemoteAddress() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoteAddress")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// RemoteAddress indicates an expected call of RemoteAddress
func (mr *MockPmAbstractPeerMockRecorder) RemoteAddress() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteAddress", reflect.TypeOf((*MockPmAbstractPeer)(nil).RemoteAddress))
}

// RemoteVerifierAddress mocks base method
func (m *MockPmAbstractPeer) RemoteVerifierAddress() common.Address {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoteVerifierAddress")
	ret0, _ := ret[0].(common.Address)
	return ret0
}

// RemoteVerifierAddress indicates an expected call of RemoteVerifierAddress
func (mr *MockPmAbstractPeerMockRecorder) RemoteVerifierAddress() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteVerifierAddress", reflect.TypeOf((*MockPmAbstractPeer)(nil).RemoteVerifierAddress))
}

// SendMsg mocks base method
func (m *MockPmAbstractPeer) SendMsg(arg0 uint64, arg1 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockPmAbstractPeerMockRecorder) SendMsg(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockPmAbstractPeer)(nil).SendMsg), arg0, arg1)
}

// SetHead mocks base method
func (m *MockPmAbstractPeer) SetHead(arg0 common.Hash, arg1 uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetHead", arg0, arg1)
}

// SetHead indicates an expected call of SetHead
func (mr *MockPmAbstractPeerMockRecorder) SetHead(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHead", reflect.TypeOf((*MockPmAbstractPeer)(nil).SetHead), arg0, arg1)
}

// SetNodeName mocks base method
func (m *MockPmAbstractPeer) SetNodeName(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetNodeName", arg0)
}

// SetNodeName indicates an expected call of SetNodeName
func (mr *MockPmAbstractPeerMockRecorder) SetNodeName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNodeName", reflect.TypeOf((*MockPmAbstractPeer)(nil).SetNodeName), arg0)
}

// SetNodeType mocks base method
func (m *MockPmAbstractPeer) SetNodeType(arg0 uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetNodeType", arg0)
}

// SetNodeType indicates an expected call of SetNodeType
func (mr *MockPmAbstractPeerMockRecorder) SetNodeType(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNodeType", reflect.TypeOf((*MockPmAbstractPeer)(nil).SetNodeType), arg0)
}

// SetNotRunning mocks base method
func (m *MockPmAbstractPeer) SetNotRunning() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetNotRunning")
}

// SetNotRunning indicates an expected call of SetNotRunning
func (mr *MockPmAbstractPeerMockRecorder) SetNotRunning() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNotRunning", reflect.TypeOf((*MockPmAbstractPeer)(nil).SetNotRunning))
}

// SetPeerRawUrl mocks base method
func (m *MockPmAbstractPeer) SetPeerRawUrl(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetPeerRawUrl", arg0)
}

// SetPeerRawUrl indicates an expected call of SetPeerRawUrl
func (mr *MockPmAbstractPeerMockRecorder) SetPeerRawUrl(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPeerRawUrl", reflect.TypeOf((*MockPmAbstractPeer)(nil).SetPeerRawUrl), arg0)
}

// SetRemoteVerifierAddress mocks base method
func (m *MockPmAbstractPeer) SetRemoteVerifierAddress(arg0 common.Address) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetRemoteVerifierAddress", arg0)
}

// SetRemoteVerifierAddress indicates an expected call of SetRemoteVerifierAddress
func (mr *MockPmAbstractPeerMockRecorder) SetRemoteVerifierAddress(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRemoteVerifierAddress", reflect.TypeOf((*MockPmAbstractPeer)(nil).SetRemoteVerifierAddress), arg0)
}
