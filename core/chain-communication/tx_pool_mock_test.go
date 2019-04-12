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

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/dipperin/dipperin-core/core/chain-communication (interfaces: TxPool)

package chain_communication

import (
	common "github.com/dipperin/dipperin-core/common"
	bloom "github.com/dipperin/dipperin-core/core/bloom"
	model "github.com/dipperin/dipperin-core/core/model"
	gomock "github.com/golang/mock/gomock"
)

// Mock of TxPool interface
type MockTxPool struct {
	ctrl     *gomock.Controller
	recorder *_MockTxPoolRecorder
}

// Recorder for MockTxPool (not exported)
type _MockTxPoolRecorder struct {
	mock *MockTxPool
}

func NewMockTxPool(ctrl *gomock.Controller) *MockTxPool {
	mock := &MockTxPool{ctrl: ctrl}
	mock.recorder = &_MockTxPoolRecorder{mock}
	return mock
}

func (_m *MockTxPool) EXPECT() *_MockTxPoolRecorder {
	return _m.recorder
}

func (_m *MockTxPool) AddLocal(_param0 model.AbstractTransaction) error {
	ret := _m.ctrl.Call(_m, "AddLocal", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockTxPoolRecorder) AddLocal(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddLocal", arg0)
}

func (_m *MockTxPool) AddLocals(_param0 []model.AbstractTransaction) []error {
	ret := _m.ctrl.Call(_m, "AddLocals", _param0)
	ret0, _ := ret[0].([]error)
	return ret0
}

func (_mr *_MockTxPoolRecorder) AddLocals(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddLocals", arg0)
}

func (_m *MockTxPool) AddRemote(_param0 model.AbstractTransaction) error {
	ret := _m.ctrl.Call(_m, "AddRemote", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockTxPoolRecorder) AddRemote(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddRemote", arg0)
}

func (_m *MockTxPool) AddRemotes(_param0 []model.AbstractTransaction) []error {
	ret := _m.ctrl.Call(_m, "AddRemotes", _param0)
	ret0, _ := ret[0].([]error)
	return ret0
}

func (_mr *_MockTxPoolRecorder) AddRemotes(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddRemotes", arg0)
}

func (_m *MockTxPool) ConvertPoolToMap() map[common.Hash]model.AbstractTransaction {
	ret := _m.ctrl.Call(_m, "ConvertPoolToMap")
	ret0, _ := ret[0].(map[common.Hash]model.AbstractTransaction)
	return ret0
}

func (_mr *_MockTxPoolRecorder) ConvertPoolToMap() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ConvertPoolToMap")
}

func (_m *MockTxPool) GetTxsEstimator(_param0 *bloom.Bloom) *bloom.HybridEstimator {
	ret := _m.ctrl.Call(_m, "GetTxsEstimator", _param0)
	ret0, _ := ret[0].(*bloom.HybridEstimator)
	return ret0
}

func (_mr *_MockTxPoolRecorder) GetTxsEstimator(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTxsEstimator", arg0)
}

func (_m *MockTxPool) Pending() (map[common.Address][]model.AbstractTransaction, error) {
	ret := _m.ctrl.Call(_m, "Pending")
	ret0, _ := ret[0].(map[common.Address][]model.AbstractTransaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTxPoolRecorder) Pending() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Pending")
}

func (_m *MockTxPool) Queueing() (map[common.Address][]model.AbstractTransaction, error) {
	ret := _m.ctrl.Call(_m, "Queueing")
	ret0, _ := ret[0].(map[common.Address][]model.AbstractTransaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTxPoolRecorder) Queueing() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Queueing")
}

func (_m *MockTxPool) Stats() (int, int) {
	ret := _m.ctrl.Call(_m, "Stats")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(int)
	return ret0, ret1
}

func (_mr *_MockTxPoolRecorder) Stats() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Stats")
}
