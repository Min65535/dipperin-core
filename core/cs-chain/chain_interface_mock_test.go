// code generated by MockGen. DO NOT EDIT.
// Source: github.com/dipperin/dipperin-core/core/cs-chain/chain-writer/middleware (interfaces: ChainInterface)

// Package chain_writer is a generated GoMock package.
package cs_chain

import (
	common "github.com/dipperin/dipperin-core/common"
	chain "github.com/dipperin/dipperin-core/core/chain"
	chain_config "github.com/dipperin/dipperin-core/core/chain-config"
	chaindb "github.com/dipperin/dipperin-core/core/chain/chaindb"
	registerdb "github.com/dipperin/dipperin-core/core/chain/registerdb"
	state_processor "github.com/dipperin/dipperin-core/core/chain/state-processor"
	economy_model "github.com/dipperin/dipperin-core/core/economy-model"
	model "github.com/dipperin/dipperin-core/core/model"
	model0 "github.com/dipperin/dipperin-core/core/vm/model"
	rlp "github.com/ethereum/go-ethereum/rlp"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockChainInterface is a mock of ChainInterface interface
type MockChainInterface struct {
	ctrl     *gomock.Controller
	recorder *MockChainInterfaceMockRecorder
}

// MockChainInterfaceMockRecorder is the mock recorder for MockChainInterface
type MockChainInterfaceMockRecorder struct {
	mock *MockChainInterface
}

// NewMockChainInterface creates a new mock instance
func NewMockChainInterface(ctrl *gomock.Controller) *MockChainInterface {
	mock := &MockChainInterface{ctrl: ctrl}
	mock.recorder = &MockChainInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChainInterface) EXPECT() *MockChainInterfaceMockRecorder {
	return m.recorder
}

// AccountStateDB mocks base method
func (m *MockChainInterface) AccountStateDB(arg0 common.Hash) (*state_processor.AccountStateDB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccountStateDB", arg0)
	ret0, _ := ret[0].(*state_processor.AccountStateDB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AccountStateDB indicates an expected call of AccountStateDB
func (mr *MockChainInterfaceMockRecorder) AccountStateDB(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountStateDB", reflect.TypeOf((*MockChainInterface)(nil).AccountStateDB), arg0)
}

// BlockProcessor mocks base method
func (m *MockChainInterface) BlockProcessor(arg0 common.Hash) (*chain.BlockProcessor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockProcessor", arg0)
	ret0, _ := ret[0].(*chain.BlockProcessor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockProcessor indicates an expected call of BlockProcessor
func (mr *MockChainInterfaceMockRecorder) BlockProcessor(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockProcessor", reflect.TypeOf((*MockChainInterface)(nil).BlockProcessor), arg0)
}

// BlockProcessorByNumber mocks base method
func (m *MockChainInterface) BlockProcessorByNumber(arg0 uint64) (*chain.BlockProcessor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockProcessorByNumber", arg0)
	ret0, _ := ret[0].(*chain.BlockProcessor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockProcessorByNumber indicates an expected call of BlockProcessorByNumber
func (mr *MockChainInterfaceMockRecorder) BlockProcessorByNumber(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockProcessorByNumber", reflect.TypeOf((*MockChainInterface)(nil).BlockProcessorByNumber), arg0)
}

// BuildRegisterProcessor mocks base method
func (m *MockChainInterface) BuildRegisterProcessor(arg0 common.Hash) (*registerdb.RegisterDB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildRegisterProcessor", arg0)
	ret0, _ := ret[0].(*registerdb.RegisterDB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildRegisterProcessor indicates an expected call of BuildRegisterProcessor
func (mr *MockChainInterfaceMockRecorder) BuildRegisterProcessor(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildRegisterProcessor", reflect.TypeOf((*MockChainInterface)(nil).BuildRegisterProcessor), arg0)
}

// BuildStateProcessor mocks base method
func (m *MockChainInterface) BuildStateProcessor(arg0 common.Hash) (*state_processor.AccountStateDB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildStateProcessor", arg0)
	ret0, _ := ret[0].(*state_processor.AccountStateDB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildStateProcessor indicates an expected call of BuildStateProcessor
func (mr *MockChainInterfaceMockRecorder) BuildStateProcessor(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildStateProcessor", reflect.TypeOf((*MockChainInterface)(nil).BuildStateProcessor), arg0)
}

// CurrentBlock mocks base method
func (m *MockChainInterface) CurrentBlock() model.AbstractBlock {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentBlock")
	ret0, _ := ret[0].(model.AbstractBlock)
	return ret0
}

// CurrentBlock indicates an expected call of CurrentBlock
func (mr *MockChainInterfaceMockRecorder) CurrentBlock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentBlock", reflect.TypeOf((*MockChainInterface)(nil).CurrentBlock))
}

// CurrentHeader mocks base method
func (m *MockChainInterface) CurrentHeader() model.AbstractHeader {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentHeader")
	ret0, _ := ret[0].(model.AbstractHeader)
	return ret0
}

// CurrentHeader indicates an expected call of CurrentHeader
func (mr *MockChainInterfaceMockRecorder) CurrentHeader() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentHeader", reflect.TypeOf((*MockChainInterface)(nil).CurrentHeader))
}

// CurrentSeed mocks base method
func (m *MockChainInterface) CurrentSeed() (common.Hash, uint64) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentSeed")
	ret0, _ := ret[0].(common.Hash)
	ret1, _ := ret[1].(uint64)
	return ret0, ret1
}

// CurrentSeed indicates an expected call of CurrentSeed
func (mr *MockChainInterfaceMockRecorder) CurrentSeed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentSeed", reflect.TypeOf((*MockChainInterface)(nil).CurrentSeed))
}

// CurrentState mocks base method
func (m *MockChainInterface) CurrentState() (*state_processor.AccountStateDB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentState")
	ret0, _ := ret[0].(*state_processor.AccountStateDB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CurrentState indicates an expected call of CurrentState
func (mr *MockChainInterfaceMockRecorder) CurrentState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentState", reflect.TypeOf((*MockChainInterface)(nil).CurrentState))
}

// Genesis mocks base method
func (m *MockChainInterface) Genesis() model.AbstractBlock {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Genesis")
	ret0, _ := ret[0].(model.AbstractBlock)
	return ret0
}

// Genesis indicates an expected call of Genesis
func (mr *MockChainInterfaceMockRecorder) Genesis() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Genesis", reflect.TypeOf((*MockChainInterface)(nil).Genesis))
}

// GetBlock mocks base method
func (m *MockChainInterface) GetBlock(arg0 common.Hash, arg1 uint64) model.AbstractBlock {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlock", arg0, arg1)
	ret0, _ := ret[0].(model.AbstractBlock)
	return ret0
}

// GetBlock indicates an expected call of GetBlock
func (mr *MockChainInterfaceMockRecorder) GetBlock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlock", reflect.TypeOf((*MockChainInterface)(nil).GetBlock), arg0, arg1)
}

// GetBlockByHash mocks base method
func (m *MockChainInterface) GetBlockByHash(arg0 common.Hash) model.AbstractBlock {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockByHash", arg0)
	ret0, _ := ret[0].(model.AbstractBlock)
	return ret0
}

// GetBlockByHash indicates an expected call of GetBlockByHash
func (mr *MockChainInterfaceMockRecorder) GetBlockByHash(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockByHash", reflect.TypeOf((*MockChainInterface)(nil).GetBlockByHash), arg0)
}

// GetBlockByNumber mocks base method
func (m *MockChainInterface) GetBlockByNumber(arg0 uint64) model.AbstractBlock {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockByNumber", arg0)
	ret0, _ := ret[0].(model.AbstractBlock)
	return ret0
}

// GetBlockByNumber indicates an expected call of GetBlockByNumber
func (mr *MockChainInterfaceMockRecorder) GetBlockByNumber(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockByNumber", reflect.TypeOf((*MockChainInterface)(nil).GetBlockByNumber), arg0)
}

// GetBlockNumber mocks base method
func (m *MockChainInterface) GetBlockNumber(arg0 common.Hash) *uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockNumber", arg0)
	ret0, _ := ret[0].(*uint64)
	return ret0
}

// GetBlockNumber indicates an expected call of GetBlockNumber
func (mr *MockChainInterfaceMockRecorder) GetBlockNumber(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockNumber", reflect.TypeOf((*MockChainInterface)(nil).GetBlockNumber), arg0)
}

// GetBloomBits mocks base method
func (m *MockChainInterface) GetBloomBits(arg0 common.Hash, arg1 uint, arg2 uint64) []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBloomBits", arg0, arg1, arg2)
	ret0, _ := ret[0].([]byte)
	return ret0
}

// GetBloomBits indicates an expected call of GetBloomBits
func (mr *MockChainInterfaceMockRecorder) GetBloomBits(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBloomBits", reflect.TypeOf((*MockChainInterface)(nil).GetBloomBits), arg0, arg1, arg2)
}

// GetBloomLog mocks base method
func (m *MockChainInterface) GetBloomLog(arg0 common.Hash, arg1 uint64) model0.Bloom {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBloomLog", arg0, arg1)
	ret0, _ := ret[0].(model0.Bloom)
	return ret0
}

// GetBloomLog indicates an expected call of GetBloomLog
func (mr *MockChainInterfaceMockRecorder) GetBloomLog(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBloomLog", reflect.TypeOf((*MockChainInterface)(nil).GetBloomLog), arg0, arg1)
}

// GetBody mocks base method
func (m *MockChainInterface) GetBody(arg0 common.Hash) model.AbstractBody {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBody", arg0)
	ret0, _ := ret[0].(model.AbstractBody)
	return ret0
}

// GetBody indicates an expected call of GetBody
func (mr *MockChainInterfaceMockRecorder) GetBody(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBody", reflect.TypeOf((*MockChainInterface)(nil).GetBody), arg0)
}

// GetBodyRLP mocks base method
func (m *MockChainInterface) GetBodyRLP(arg0 common.Hash) rlp.RawValue {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBodyRLP", arg0)
	ret0, _ := ret[0].(rlp.RawValue)
	return ret0
}

// GetBodyRLP indicates an expected call of GetBodyRLP
func (mr *MockChainInterfaceMockRecorder) GetBodyRLP(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBodyRLP", reflect.TypeOf((*MockChainInterface)(nil).GetBodyRLP), arg0)
}

// GetChainConfig mocks base method
func (m *MockChainInterface) GetChainConfig() *chain_config.ChainConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChainConfig")
	ret0, _ := ret[0].(*chain_config.ChainConfig)
	return ret0
}

// GetChainConfig indicates an expected call of GetChainConfig
func (mr *MockChainInterfaceMockRecorder) GetChainConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChainConfig", reflect.TypeOf((*MockChainInterface)(nil).GetChainConfig))
}

// GetChainDB mocks base method
func (m *MockChainInterface) GetChainDB() chaindb.Database {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChainDB")
	ret0, _ := ret[0].(chaindb.Database)
	return ret0
}

// GetChainDB indicates an expected call of GetChainDB
func (mr *MockChainInterfaceMockRecorder) GetChainDB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChainDB", reflect.TypeOf((*MockChainInterface)(nil).GetChainDB))
}

// GetCurrVerifiers mocks base method
func (m *MockChainInterface) GetCurrVerifiers() []common.Address {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrVerifiers")
	ret0, _ := ret[0].([]common.Address)
	return ret0
}

// GetCurrVerifiers indicates an expected call of GetCurrVerifiers
func (mr *MockChainInterfaceMockRecorder) GetCurrVerifiers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrVerifiers", reflect.TypeOf((*MockChainInterface)(nil).GetCurrVerifiers))
}

// GetEconomyModel mocks base method
func (m *MockChainInterface) GetEconomyModel() economy_model.EconomyModel {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEconomyModel")
	ret0, _ := ret[0].(economy_model.EconomyModel)
	return ret0
}

// GetEconomyModel indicates an expected call of GetEconomyModel
func (mr *MockChainInterfaceMockRecorder) GetEconomyModel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEconomyModel", reflect.TypeOf((*MockChainInterface)(nil).GetEconomyModel))
}

// GetHeader mocks base method
func (m *MockChainInterface) GetHeader(arg0 common.Hash, arg1 uint64) model.AbstractHeader {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHeader", arg0, arg1)
	ret0, _ := ret[0].(model.AbstractHeader)
	return ret0
}

// GetHeader indicates an expected call of GetHeader
func (mr *MockChainInterfaceMockRecorder) GetHeader(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeader", reflect.TypeOf((*MockChainInterface)(nil).GetHeader), arg0, arg1)
}

// GetHeaderByHash mocks base method
func (m *MockChainInterface) GetHeaderByHash(arg0 common.Hash) model.AbstractHeader {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHeaderByHash", arg0)
	ret0, _ := ret[0].(model.AbstractHeader)
	return ret0
}

// GetHeaderByHash indicates an expected call of GetHeaderByHash
func (mr *MockChainInterfaceMockRecorder) GetHeaderByHash(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeaderByHash", reflect.TypeOf((*MockChainInterface)(nil).GetHeaderByHash), arg0)
}

// GetHeaderByNumber mocks base method
func (m *MockChainInterface) GetHeaderByNumber(arg0 uint64) model.AbstractHeader {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHeaderByNumber", arg0)
	ret0, _ := ret[0].(model.AbstractHeader)
	return ret0
}

// GetHeaderByNumber indicates an expected call of GetHeaderByNumber
func (mr *MockChainInterfaceMockRecorder) GetHeaderByNumber(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeaderByNumber", reflect.TypeOf((*MockChainInterface)(nil).GetHeaderByNumber), arg0)
}

// GetHeaderRLP mocks base method
func (m *MockChainInterface) GetHeaderRLP(arg0 common.Hash) rlp.RawValue {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHeaderRLP", arg0)
	ret0, _ := ret[0].(rlp.RawValue)
	return ret0
}

// GetHeaderRLP indicates an expected call of GetHeaderRLP
func (mr *MockChainInterfaceMockRecorder) GetHeaderRLP(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeaderRLP", reflect.TypeOf((*MockChainInterface)(nil).GetHeaderRLP), arg0)
}

// GetLastChangePoint mocks base method
func (m *MockChainInterface) GetLastChangePoint(arg0 model.AbstractBlock) *uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastChangePoint", arg0)
	ret0, _ := ret[0].(*uint64)
	return ret0
}

// GetLastChangePoint indicates an expected call of GetLastChangePoint
func (mr *MockChainInterfaceMockRecorder) GetLastChangePoint(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastChangePoint", reflect.TypeOf((*MockChainInterface)(nil).GetLastChangePoint), arg0)
}

// GetLatestNormalBlock mocks base method
func (m *MockChainInterface) GetLatestNormalBlock() model.AbstractBlock {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestNormalBlock")
	ret0, _ := ret[0].(model.AbstractBlock)
	return ret0
}

// GetLatestNormalBlock indicates an expected call of GetLatestNormalBlock
func (mr *MockChainInterfaceMockRecorder) GetLatestNormalBlock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestNormalBlock", reflect.TypeOf((*MockChainInterface)(nil).GetLatestNormalBlock))
}

// GetNextVerifiers mocks base method
func (m *MockChainInterface) GetNextVerifiers() []common.Address {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextVerifiers")
	ret0, _ := ret[0].([]common.Address)
	return ret0
}

// GetNextVerifiers indicates an expected call of GetNextVerifiers
func (mr *MockChainInterfaceMockRecorder) GetNextVerifiers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextVerifiers", reflect.TypeOf((*MockChainInterface)(nil).GetNextVerifiers))
}

// GetReceipts mocks base method
func (m *MockChainInterface) GetReceipts(arg0 common.Hash, arg1 uint64) model0.Receipts {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReceipts", arg0, arg1)
	ret0, _ := ret[0].(model0.Receipts)
	return ret0
}

// GetReceipts indicates an expected call of GetReceipts
func (mr *MockChainInterfaceMockRecorder) GetReceipts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReceipts", reflect.TypeOf((*MockChainInterface)(nil).GetReceipts), arg0, arg1)
}

// GetSlot mocks base method
func (m *MockChainInterface) GetSlot(arg0 model.AbstractBlock) *uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSlot", arg0)
	ret0, _ := ret[0].(*uint64)
	return ret0
}

// GetSlot indicates an expected call of GetSlot
func (mr *MockChainInterfaceMockRecorder) GetSlot(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSlot", reflect.TypeOf((*MockChainInterface)(nil).GetSlot), arg0)
}

// GetSlotByNum mocks base method
func (m *MockChainInterface) GetSlotByNum(arg0 uint64) *uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSlotByNum", arg0)
	ret0, _ := ret[0].(*uint64)
	return ret0
}

// GetSlotByNum indicates an expected call of GetSlotByNum
func (mr *MockChainInterfaceMockRecorder) GetSlotByNum(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSlotByNum", reflect.TypeOf((*MockChainInterface)(nil).GetSlotByNum), arg0)
}

// GetStateStorage mocks base method
func (m *MockChainInterface) GetStateStorage() state_processor.StateStorage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStateStorage")
	ret0, _ := ret[0].(state_processor.StateStorage)
	return ret0
}

// GetStateStorage indicates an expected call of GetStateStorage
func (mr *MockChainInterfaceMockRecorder) GetStateStorage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStateStorage", reflect.TypeOf((*MockChainInterface)(nil).GetStateStorage))
}

// GetTransaction mocks base method
func (m *MockChainInterface) GetTransaction(arg0 common.Hash) (model.AbstractTransaction, common.Hash, uint64, uint64) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransaction", arg0)
	ret0, _ := ret[0].(model.AbstractTransaction)
	ret1, _ := ret[1].(common.Hash)
	ret2, _ := ret[2].(uint64)
	ret3, _ := ret[3].(uint64)
	return ret0, ret1, ret2, ret3
}

// GetTransaction indicates an expected call of GetTransaction
func (mr *MockChainInterfaceMockRecorder) GetTransaction(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransaction", reflect.TypeOf((*MockChainInterface)(nil).GetTransaction), arg0)
}

// GetVerifiers mocks base method
func (m *MockChainInterface) GetVerifiers(arg0 uint64) []common.Address {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVerifiers", arg0)
	ret0, _ := ret[0].([]common.Address)
	return ret0
}

// GetVerifiers indicates an expected call of GetVerifiers
func (mr *MockChainInterfaceMockRecorder) GetVerifiers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVerifiers", reflect.TypeOf((*MockChainInterface)(nil).GetVerifiers), arg0)
}

// HasBlock mocks base method
func (m *MockChainInterface) HasBlock(arg0 common.Hash, arg1 uint64) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasBlock", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasBlock indicates an expected call of HasBlock
func (mr *MockChainInterfaceMockRecorder) HasBlock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasBlock", reflect.TypeOf((*MockChainInterface)(nil).HasBlock), arg0, arg1)
}

// HasHeader mocks base method
func (m *MockChainInterface) HasHeader(arg0 common.Hash, arg1 uint64) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasHeader", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasHeader indicates an expected call of HasHeader
func (mr *MockChainInterfaceMockRecorder) HasHeader(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasHeader", reflect.TypeOf((*MockChainInterface)(nil).HasHeader), arg0, arg1)
}

// IsChangePoint mocks base method
func (m *MockChainInterface) IsChangePoint(arg0 model.AbstractBlock, arg1 bool) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsChangePoint", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsChangePoint indicates an expected call of IsChangePoint
func (mr *MockChainInterfaceMockRecorder) IsChangePoint(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsChangePoint", reflect.TypeOf((*MockChainInterface)(nil).IsChangePoint), arg0, arg1)
}

// NumBeforeLastBySlot mocks base method
func (m *MockChainInterface) NumBeforeLastBySlot(arg0 uint64) *uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NumBeforeLastBySlot", arg0)
	ret0, _ := ret[0].(*uint64)
	return ret0
}

// NumBeforeLastBySlot indicates an expected call of NumBeforeLastBySlot
func (mr *MockChainInterfaceMockRecorder) NumBeforeLastBySlot(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NumBeforeLastBySlot", reflect.TypeOf((*MockChainInterface)(nil).NumBeforeLastBySlot), arg0)
}

// Rollback mocks base method
func (m *MockChainInterface) Rollback(arg0 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rollback", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Rollback indicates an expected call of Rollback
func (mr *MockChainInterfaceMockRecorder) Rollback(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockChainInterface)(nil).Rollback), arg0)
}

// SaveBlock mocks base method
func (m *MockChainInterface) SaveBlock(arg0 model.AbstractBlock) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveBlock", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveBlock indicates an expected call of SaveBlock
func (mr *MockChainInterfaceMockRecorder) SaveBlock(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveBlock", reflect.TypeOf((*MockChainInterface)(nil).SaveBlock), arg0)
}

// StateAtByBlockNumber mocks base method
func (m *MockChainInterface) StateAtByBlockNumber(arg0 uint64) (*state_processor.AccountStateDB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateAtByBlockNumber", arg0)
	ret0, _ := ret[0].(*state_processor.AccountStateDB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateAtByBlockNumber indicates an expected call of StateAtByBlockNumber
func (mr *MockChainInterfaceMockRecorder) StateAtByBlockNumber(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateAtByBlockNumber", reflect.TypeOf((*MockChainInterface)(nil).StateAtByBlockNumber), arg0)
}

// StateAtByStateRoot mocks base method
func (m *MockChainInterface) StateAtByStateRoot(arg0 common.Hash) (*state_processor.AccountStateDB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateAtByStateRoot", arg0)
	ret0, _ := ret[0].(*state_processor.AccountStateDB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateAtByStateRoot indicates an expected call of StateAtByStateRoot
func (mr *MockChainInterfaceMockRecorder) StateAtByStateRoot(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateAtByStateRoot", reflect.TypeOf((*MockChainInterface)(nil).StateAtByStateRoot), arg0)
}
