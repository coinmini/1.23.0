// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/filecoin-project/lotus/storage/pipeline (interfaces: CommitBatcherApi)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	address "github.com/filecoin-project/go-address"
	abi "github.com/filecoin-project/go-state-types/abi"
	big "github.com/filecoin-project/go-state-types/big"
	miner "github.com/filecoin-project/go-state-types/builtin/v8/miner"
	network "github.com/filecoin-project/go-state-types/network"
	api "github.com/filecoin-project/lotus/api"
	types "github.com/filecoin-project/lotus/chain/types"
	gomock "github.com/golang/mock/gomock"
	cid "github.com/ipfs/go-cid"
)

// MockCommitBatcherApi is a mock of CommitBatcherApi interface.
type MockCommitBatcherApi struct {
	ctrl     *gomock.Controller
	recorder *MockCommitBatcherApiMockRecorder
}

// MockCommitBatcherApiMockRecorder is the mock recorder for MockCommitBatcherApi.
type MockCommitBatcherApiMockRecorder struct {
	mock *MockCommitBatcherApi
}

// NewMockCommitBatcherApi creates a new mock instance.
func NewMockCommitBatcherApi(ctrl *gomock.Controller) *MockCommitBatcherApi {
	mock := &MockCommitBatcherApi{ctrl: ctrl}
	mock.recorder = &MockCommitBatcherApiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommitBatcherApi) EXPECT() *MockCommitBatcherApiMockRecorder {
	return m.recorder
}

// ChainHead mocks base method.
func (m *MockCommitBatcherApi) ChainHead(arg0 context.Context) (*types.TipSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChainHead", arg0)
	ret0, _ := ret[0].(*types.TipSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChainHead indicates an expected call of ChainHead.
func (mr *MockCommitBatcherApiMockRecorder) ChainHead(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChainHead", reflect.TypeOf((*MockCommitBatcherApi)(nil).ChainHead), arg0)
}

// SendMsg mocks base method.
func (m *MockCommitBatcherApi) SendMsg(arg0 context.Context, arg1, arg2 address.Address, arg3 abi.MethodNum, arg4, arg5 big.Int, arg6 []byte) (cid.Cid, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(cid.Cid)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockCommitBatcherApiMockRecorder) SendMsg(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockCommitBatcherApi)(nil).SendMsg), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// StateMinerAvailableBalance mocks base method.
func (m *MockCommitBatcherApi) StateMinerAvailableBalance(arg0 context.Context, arg1 address.Address, arg2 types.TipSetKey) (big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateMinerAvailableBalance", arg0, arg1, arg2)
	ret0, _ := ret[0].(big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateMinerAvailableBalance indicates an expected call of StateMinerAvailableBalance.
func (mr *MockCommitBatcherApiMockRecorder) StateMinerAvailableBalance(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateMinerAvailableBalance", reflect.TypeOf((*MockCommitBatcherApi)(nil).StateMinerAvailableBalance), arg0, arg1, arg2)
}

// StateMinerInfo mocks base method.
func (m *MockCommitBatcherApi) StateMinerInfo(arg0 context.Context, arg1 address.Address, arg2 types.TipSetKey) (api.MinerInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateMinerInfo", arg0, arg1, arg2)
	ret0, _ := ret[0].(api.MinerInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateMinerInfo indicates an expected call of StateMinerInfo.
func (mr *MockCommitBatcherApiMockRecorder) StateMinerInfo(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateMinerInfo", reflect.TypeOf((*MockCommitBatcherApi)(nil).StateMinerInfo), arg0, arg1, arg2)
}

// StateMinerInitialPledgeCollateral mocks base method.
func (m *MockCommitBatcherApi) StateMinerInitialPledgeCollateral(arg0 context.Context, arg1 address.Address, arg2 miner.SectorPreCommitInfo, arg3 types.TipSetKey) (big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateMinerInitialPledgeCollateral", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateMinerInitialPledgeCollateral indicates an expected call of StateMinerInitialPledgeCollateral.
func (mr *MockCommitBatcherApiMockRecorder) StateMinerInitialPledgeCollateral(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateMinerInitialPledgeCollateral", reflect.TypeOf((*MockCommitBatcherApi)(nil).StateMinerInitialPledgeCollateral), arg0, arg1, arg2, arg3)
}

// StateNetworkVersion mocks base method.
func (m *MockCommitBatcherApi) StateNetworkVersion(arg0 context.Context, arg1 types.TipSetKey) (network.Version, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateNetworkVersion", arg0, arg1)
	ret0, _ := ret[0].(network.Version)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateNetworkVersion indicates an expected call of StateNetworkVersion.
func (mr *MockCommitBatcherApiMockRecorder) StateNetworkVersion(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateNetworkVersion", reflect.TypeOf((*MockCommitBatcherApi)(nil).StateNetworkVersion), arg0, arg1)
}

// StateSectorPreCommitInfo mocks base method.
func (m *MockCommitBatcherApi) StateSectorPreCommitInfo(arg0 context.Context, arg1 address.Address, arg2 abi.SectorNumber, arg3 types.TipSetKey) (*miner.SectorPreCommitOnChainInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateSectorPreCommitInfo", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*miner.SectorPreCommitOnChainInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateSectorPreCommitInfo indicates an expected call of StateSectorPreCommitInfo.
func (mr *MockCommitBatcherApiMockRecorder) StateSectorPreCommitInfo(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateSectorPreCommitInfo", reflect.TypeOf((*MockCommitBatcherApi)(nil).StateSectorPreCommitInfo), arg0, arg1, arg2, arg3)
}
