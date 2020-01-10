// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dipperin/dipperin-core/core/mine/minemsg (interfaces: Work)

// Package minemsg_mock is a generated GoMock package.
package minemsg_mock

import (
	common "github.com/dipperin/dipperin-core/common"
	model "github.com/dipperin/dipperin-core/core/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockWork is a mock of Work interface
type MockWork struct {
	ctrl     *gomock.Controller
	recorder *MockWorkMockRecorder
}

// MockWorkMockRecorder is the mock recorder for MockWork
type MockWorkMockRecorder struct {
	mock *MockWork
}

// NewMockWork creates a new mock instance
func NewMockWork(ctrl *gomock.Controller) *MockWork {
	mock := &MockWork{ctrl: ctrl}
	mock.recorder = &MockWorkMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWork) EXPECT() *MockWorkMockRecorder {
	return m.recorder
}

// FillSealResult mocks base method
func (m *MockWork) FillSealResult(arg0 model.AbstractBlock) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FillSealResult", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// FillSealResult indicates an expected call of FillSealResult
func (mr *MockWorkMockRecorder) FillSealResult(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FillSealResult", reflect.TypeOf((*MockWork)(nil).FillSealResult), arg0)
}

// GetWorkerCoinbaseAddress mocks base method
func (m *MockWork) GetWorkerCoinbaseAddress() common.Address {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkerCoinbaseAddress")
	ret0, _ := ret[0].(common.Address)
	return ret0
}

// GetWorkerCoinbaseAddress indicates an expected call of GetWorkerCoinbaseAddress
func (mr *MockWorkMockRecorder) GetWorkerCoinbaseAddress() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkerCoinbaseAddress", reflect.TypeOf((*MockWork)(nil).GetWorkerCoinbaseAddress))
}

// SetWorkerCoinbaseAddress mocks base method
func (m *MockWork) SetWorkerCoinbaseAddress(arg0 common.Address) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetWorkerCoinbaseAddress", arg0)
}

// SetWorkerCoinbaseAddress indicates an expected call of SetWorkerCoinbaseAddress
func (mr *MockWorkMockRecorder) SetWorkerCoinbaseAddress(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWorkerCoinbaseAddress", reflect.TypeOf((*MockWork)(nil).SetWorkerCoinbaseAddress), arg0)
}
