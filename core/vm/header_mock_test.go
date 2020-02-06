// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dipperin/dipperin-core/core/model (interfaces: AbstractHeader)

// Package vm is a generated GoMock package.
package vm

import (
	ecdsa "crypto/ecdsa"
	common "github.com/dipperin/dipperin-core/common"
	model "github.com/dipperin/dipperin-core/core/model"
	gomock "github.com/golang/mock/gomock"
	big "math/big"
	reflect "reflect"
)

// MockAbstractHeader is a mock of AbstractHeader interface
type MockAbstractHeader struct {
	ctrl     *gomock.Controller
	recorder *MockAbstractHeaderMockRecorder
}

// MockAbstractHeaderMockRecorder is the mock recorder for MockAbstractHeader
type MockAbstractHeaderMockRecorder struct {
	mock *MockAbstractHeader
}

// NewMockAbstractHeader creates a new mock instance
func NewMockAbstractHeader(ctrl *gomock.Controller) *MockAbstractHeader {
	mock := &MockAbstractHeader{ctrl: ctrl}
	mock.recorder = &MockAbstractHeaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAbstractHeader) EXPECT() *MockAbstractHeaderMockRecorder {
	return m.recorder
}

// CoinBaseAddress mocks base method
func (m *MockAbstractHeader) CoinBaseAddress() common.Address {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CoinBaseAddress")
	ret0, _ := ret[0].(common.Address)
	return ret0
}

// CoinBaseAddress indicates an expected call of CoinBaseAddress
func (mr *MockAbstractHeaderMockRecorder) CoinBaseAddress() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CoinBaseAddress", reflect.TypeOf((*MockAbstractHeader)(nil).CoinBaseAddress))
}

// DuplicateHeader mocks base method
func (m *MockAbstractHeader) DuplicateHeader() model.AbstractHeader {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DuplicateHeader")
	ret0, _ := ret[0].(model.AbstractHeader)
	return ret0
}

// DuplicateHeader indicates an expected call of DuplicateHeader
func (mr *MockAbstractHeaderMockRecorder) DuplicateHeader() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DuplicateHeader", reflect.TypeOf((*MockAbstractHeader)(nil).DuplicateHeader))
}

// EncodeRlpToBytes mocks base method
func (m *MockAbstractHeader) EncodeRlpToBytes() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EncodeRlpToBytes")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EncodeRlpToBytes indicates an expected call of EncodeRlpToBytes
func (mr *MockAbstractHeaderMockRecorder) EncodeRlpToBytes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EncodeRlpToBytes", reflect.TypeOf((*MockAbstractHeader)(nil).EncodeRlpToBytes))
}

// GetDifficulty mocks base method
func (m *MockAbstractHeader) GetDifficulty() common.Difficulty {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDifficulty")
	ret0, _ := ret[0].(common.Difficulty)
	return ret0
}

// GetDifficulty indicates an expected call of GetDifficulty
func (mr *MockAbstractHeaderMockRecorder) GetDifficulty() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDifficulty", reflect.TypeOf((*MockAbstractHeader)(nil).GetDifficulty))
}

// GetGasLimit mocks base method
func (m *MockAbstractHeader) GetGasLimit() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGasLimit")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetGasLimit indicates an expected call of GetGasLimit
func (mr *MockAbstractHeaderMockRecorder) GetGasLimit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGasLimit", reflect.TypeOf((*MockAbstractHeader)(nil).GetGasLimit))
}

// GetGasUsed mocks base method
func (m *MockAbstractHeader) GetGasUsed() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGasUsed")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetGasUsed indicates an expected call of GetGasUsed
func (mr *MockAbstractHeaderMockRecorder) GetGasUsed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGasUsed", reflect.TypeOf((*MockAbstractHeader)(nil).GetGasUsed))
}

// GetInterLinkRoot mocks base method
func (m *MockAbstractHeader) GetInterLinkRoot() common.Hash {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInterLinkRoot")
	ret0, _ := ret[0].(common.Hash)
	return ret0
}

// GetInterLinkRoot indicates an expected call of GetInterLinkRoot
func (mr *MockAbstractHeaderMockRecorder) GetInterLinkRoot() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInterLinkRoot", reflect.TypeOf((*MockAbstractHeader)(nil).GetInterLinkRoot))
}

// GetMinerPubKey mocks base method
func (m *MockAbstractHeader) GetMinerPubKey() *ecdsa.PublicKey {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMinerPubKey")
	ret0, _ := ret[0].(*ecdsa.PublicKey)
	return ret0
}

// GetMinerPubKey indicates an expected call of GetMinerPubKey
func (mr *MockAbstractHeaderMockRecorder) GetMinerPubKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMinerPubKey", reflect.TypeOf((*MockAbstractHeader)(nil).GetMinerPubKey))
}

// GetNumber mocks base method
func (m *MockAbstractHeader) GetNumber() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNumber")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetNumber indicates an expected call of GetNumber
func (mr *MockAbstractHeaderMockRecorder) GetNumber() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNumber", reflect.TypeOf((*MockAbstractHeader)(nil).GetNumber))
}

// GetPreHash mocks base method
func (m *MockAbstractHeader) GetPreHash() common.Hash {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPreHash")
	ret0, _ := ret[0].(common.Hash)
	return ret0
}

// GetPreHash indicates an expected call of GetPreHash
func (mr *MockAbstractHeaderMockRecorder) GetPreHash() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPreHash", reflect.TypeOf((*MockAbstractHeader)(nil).GetPreHash))
}

// GetProof mocks base method
func (m *MockAbstractHeader) GetProof() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProof")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// GetProof indicates an expected call of GetProof
func (mr *MockAbstractHeaderMockRecorder) GetProof() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProof", reflect.TypeOf((*MockAbstractHeader)(nil).GetProof))
}

// GetRegisterRoot mocks base method
func (m *MockAbstractHeader) GetRegisterRoot() common.Hash {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRegisterRoot")
	ret0, _ := ret[0].(common.Hash)
	return ret0
}

// GetRegisterRoot indicates an expected call of GetRegisterRoot
func (mr *MockAbstractHeaderMockRecorder) GetRegisterRoot() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRegisterRoot", reflect.TypeOf((*MockAbstractHeader)(nil).GetRegisterRoot))
}

// GetSeed mocks base method
func (m *MockAbstractHeader) GetSeed() common.Hash {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSeed")
	ret0, _ := ret[0].(common.Hash)
	return ret0
}

// GetSeed indicates an expected call of GetSeed
func (mr *MockAbstractHeaderMockRecorder) GetSeed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSeed", reflect.TypeOf((*MockAbstractHeader)(nil).GetSeed))
}

// GetStateRoot mocks base method
func (m *MockAbstractHeader) GetStateRoot() common.Hash {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStateRoot")
	ret0, _ := ret[0].(common.Hash)
	return ret0
}

// GetStateRoot indicates an expected call of GetStateRoot
func (mr *MockAbstractHeaderMockRecorder) GetStateRoot() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStateRoot", reflect.TypeOf((*MockAbstractHeader)(nil).GetStateRoot))
}

// GetTimeStamp mocks base method
func (m *MockAbstractHeader) GetTimeStamp() *big.Int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTimeStamp")
	ret0, _ := ret[0].(*big.Int)
	return ret0
}

// GetTimeStamp indicates an expected call of GetTimeStamp
func (mr *MockAbstractHeaderMockRecorder) GetTimeStamp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTimeStamp", reflect.TypeOf((*MockAbstractHeader)(nil).GetTimeStamp))
}

// Hash mocks base method
func (m *MockAbstractHeader) Hash() common.Hash {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hash")
	ret0, _ := ret[0].(common.Hash)
	return ret0
}

// Hash indicates an expected call of Hash
func (mr *MockAbstractHeaderMockRecorder) Hash() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hash", reflect.TypeOf((*MockAbstractHeader)(nil).Hash))
}

// IsEqual mocks base method
func (m *MockAbstractHeader) IsEqual(arg0 model.AbstractHeader) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsEqual", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsEqual indicates an expected call of IsEqual
func (mr *MockAbstractHeaderMockRecorder) IsEqual(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsEqual", reflect.TypeOf((*MockAbstractHeader)(nil).IsEqual), arg0)
}

// SetRegisterRoot mocks base method
func (m *MockAbstractHeader) SetRegisterRoot(arg0 common.Hash) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetRegisterRoot", arg0)
}

// SetRegisterRoot indicates an expected call of SetRegisterRoot
func (mr *MockAbstractHeaderMockRecorder) SetRegisterRoot(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRegisterRoot", reflect.TypeOf((*MockAbstractHeader)(nil).SetRegisterRoot), arg0)
}

// SetVerificationRoot mocks base method
func (m *MockAbstractHeader) SetVerificationRoot(arg0 common.Hash) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetVerificationRoot", arg0)
}

// SetVerificationRoot indicates an expected call of SetVerificationRoot
func (mr *MockAbstractHeaderMockRecorder) SetVerificationRoot(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetVerificationRoot", reflect.TypeOf((*MockAbstractHeader)(nil).SetVerificationRoot), arg0)
}