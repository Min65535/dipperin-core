// Copyright 2019, Keychain Foundation Ltd.
// This file is part of the dipperin-core library.
//
// The dipperin-core library is free software: you can redistribute
// it and/or modify it under the terms of the GNU Lesser General Public License
// as published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// The dipperin-core library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dipperin/dipperin-core/core/cs-chain (interfaces: TxPool)

// Package cs_chain is a generated GoMock package.
package cs_chain

import (
	model "github.com/dipperin/dipperin-core/core/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockTxPool is a mock of TxPool interface
type MockTxPool struct {
	ctrl     *gomock.Controller
	recorder *MockTxPoolMockRecorder
}

// MockTxPoolMockRecorder is the mock recorder for MockTxPool
type MockTxPoolMockRecorder struct {
	mock *MockTxPool
}

// NewMockTxPool creates a new mock instance
func NewMockTxPool(ctrl *gomock.Controller) *MockTxPool {
	mock := &MockTxPool{ctrl: ctrl}
	mock.recorder = &MockTxPoolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTxPool) EXPECT() *MockTxPoolMockRecorder {
	return m.recorder
}

// Reset mocks base method
func (m *MockTxPool) Reset(arg0, arg1 *model.Header) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Reset", arg0, arg1)
}

// Reset indicates an expected call of Reset
func (mr *MockTxPoolMockRecorder) Reset(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockTxPool)(nil).Reset), arg0, arg1)
}
