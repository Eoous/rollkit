// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	types "github.com/cometbft/cometbft/abci/types"
	mock "github.com/stretchr/testify/mock"
)

// Application is an autogenerated mock type for the Application type
type Application struct {
	mock.Mock
}

// ApplySnapshotChunk provides a mock function with given fields: _a0, _a1
func (_m *Application) ApplySnapshotChunk(_a0 context.Context, _a1 *types.RequestApplySnapshotChunk) (*types.ResponseApplySnapshotChunk, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.ResponseApplySnapshotChunk
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.RequestApplySnapshotChunk) (*types.ResponseApplySnapshotChunk, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.RequestApplySnapshotChunk) *types.ResponseApplySnapshotChunk); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseApplySnapshotChunk)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.RequestApplySnapshotChunk) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BeginBlock provides a mock function with given fields: _a0, _a1
func (_m *Application) BeginBlock(_a0 context.Context, _a1 *types.RequestBeginBlock) (*types.ResponseBeginBlock, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.ResponseBeginBlock
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.RequestBeginBlock) (*types.ResponseBeginBlock, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.RequestBeginBlock) *types.ResponseBeginBlock); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseBeginBlock)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.RequestBeginBlock) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckTx provides a mock function with given fields: _a0, _a1
func (_m *Application) CheckTx(_a0 context.Context, _a1 *types.RequestCheckTx) (*types.ResponseCheckTx, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.ResponseCheckTx
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.RequestCheckTx) (*types.ResponseCheckTx, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.RequestCheckTx) *types.ResponseCheckTx); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseCheckTx)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.RequestCheckTx) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Commit provides a mock function with given fields: _a0, _a1
func (_m *Application) Commit(_a0 context.Context, _a1 *types.RequestCommit) (*types.ResponseCommit, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.ResponseCommit
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.RequestCommit) (*types.ResponseCommit, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.RequestCommit) *types.ResponseCommit); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseCommit)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.RequestCommit) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeliverTx provides a mock function with given fields: _a0, _a1
func (_m *Application) DeliverTx(_a0 context.Context, _a1 *types.RequestDeliverTx) (*types.ResponseDeliverTx, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.ResponseDeliverTx
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.RequestDeliverTx) (*types.ResponseDeliverTx, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.RequestDeliverTx) *types.ResponseDeliverTx); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseDeliverTx)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.RequestDeliverTx) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EndBlock provides a mock function with given fields: _a0, _a1
func (_m *Application) EndBlock(_a0 context.Context, _a1 *types.RequestEndBlock) (*types.ResponseEndBlock, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.ResponseEndBlock
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.RequestEndBlock) (*types.ResponseEndBlock, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.RequestEndBlock) *types.ResponseEndBlock); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseEndBlock)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.RequestEndBlock) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExtendVote provides a mock function with given fields: _a0, _a1
func (_m *Application) ExtendVote(_a0 context.Context, _a1 *types.RequestExtendVote) (*types.ResponseExtendVote, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.ResponseExtendVote
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.RequestExtendVote) (*types.ResponseExtendVote, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.RequestExtendVote) *types.ResponseExtendVote); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseExtendVote)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.RequestExtendVote) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FinalizeBlock provides a mock function with given fields: _a0, _a1
func (_m *Application) FinalizeBlock(_a0 context.Context, _a1 *types.RequestFinalizeBlock) (*types.ResponseFinalizeBlock, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.ResponseFinalizeBlock
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.RequestFinalizeBlock) (*types.ResponseFinalizeBlock, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.RequestFinalizeBlock) *types.ResponseFinalizeBlock); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseFinalizeBlock)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.RequestFinalizeBlock) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateFraudProof provides a mock function with given fields: _a0, _a1
func (_m *Application) GenerateFraudProof(_a0 context.Context, _a1 *types.RequestGenerateFraudProof) (*types.ResponseGenerateFraudProof, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.ResponseGenerateFraudProof
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.RequestGenerateFraudProof) (*types.ResponseGenerateFraudProof, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.RequestGenerateFraudProof) *types.ResponseGenerateFraudProof); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseGenerateFraudProof)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.RequestGenerateFraudProof) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAppHash provides a mock function with given fields: _a0, _a1
func (_m *Application) GetAppHash(_a0 context.Context, _a1 *types.RequestGetAppHash) (*types.ResponseGetAppHash, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.ResponseGetAppHash
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.RequestGetAppHash) (*types.ResponseGetAppHash, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.RequestGetAppHash) *types.ResponseGetAppHash); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseGetAppHash)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.RequestGetAppHash) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Info provides a mock function with given fields: _a0, _a1
func (_m *Application) Info(_a0 context.Context, _a1 *types.RequestInfo) (*types.ResponseInfo, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.ResponseInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.RequestInfo) (*types.ResponseInfo, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.RequestInfo) *types.ResponseInfo); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.RequestInfo) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InitChain provides a mock function with given fields: _a0, _a1
func (_m *Application) InitChain(_a0 context.Context, _a1 *types.RequestInitChain) (*types.ResponseInitChain, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.ResponseInitChain
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.RequestInitChain) (*types.ResponseInitChain, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.RequestInitChain) *types.ResponseInitChain); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseInitChain)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.RequestInitChain) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListSnapshots provides a mock function with given fields: _a0, _a1
func (_m *Application) ListSnapshots(_a0 context.Context, _a1 *types.RequestListSnapshots) (*types.ResponseListSnapshots, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.ResponseListSnapshots
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.RequestListSnapshots) (*types.ResponseListSnapshots, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.RequestListSnapshots) *types.ResponseListSnapshots); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseListSnapshots)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.RequestListSnapshots) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoadSnapshotChunk provides a mock function with given fields: _a0, _a1
func (_m *Application) LoadSnapshotChunk(_a0 context.Context, _a1 *types.RequestLoadSnapshotChunk) (*types.ResponseLoadSnapshotChunk, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.ResponseLoadSnapshotChunk
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.RequestLoadSnapshotChunk) (*types.ResponseLoadSnapshotChunk, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.RequestLoadSnapshotChunk) *types.ResponseLoadSnapshotChunk); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseLoadSnapshotChunk)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.RequestLoadSnapshotChunk) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OfferSnapshot provides a mock function with given fields: _a0, _a1
func (_m *Application) OfferSnapshot(_a0 context.Context, _a1 *types.RequestOfferSnapshot) (*types.ResponseOfferSnapshot, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.ResponseOfferSnapshot
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.RequestOfferSnapshot) (*types.ResponseOfferSnapshot, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.RequestOfferSnapshot) *types.ResponseOfferSnapshot); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseOfferSnapshot)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.RequestOfferSnapshot) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PrepareProposal provides a mock function with given fields: _a0, _a1
func (_m *Application) PrepareProposal(_a0 context.Context, _a1 *types.RequestPrepareProposal) (*types.ResponsePrepareProposal, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.ResponsePrepareProposal
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.RequestPrepareProposal) (*types.ResponsePrepareProposal, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.RequestPrepareProposal) *types.ResponsePrepareProposal); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponsePrepareProposal)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.RequestPrepareProposal) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProcessProposal provides a mock function with given fields: _a0, _a1
func (_m *Application) ProcessProposal(_a0 context.Context, _a1 *types.RequestProcessProposal) (*types.ResponseProcessProposal, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.ResponseProcessProposal
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.RequestProcessProposal) (*types.ResponseProcessProposal, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.RequestProcessProposal) *types.ResponseProcessProposal); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseProcessProposal)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.RequestProcessProposal) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Query provides a mock function with given fields: _a0, _a1
func (_m *Application) Query(_a0 context.Context, _a1 *types.RequestQuery) (*types.ResponseQuery, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.ResponseQuery
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.RequestQuery) (*types.ResponseQuery, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.RequestQuery) *types.ResponseQuery); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseQuery)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.RequestQuery) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifyFraudProof provides a mock function with given fields: _a0, _a1
func (_m *Application) VerifyFraudProof(_a0 context.Context, _a1 *types.RequestVerifyFraudProof) (*types.ResponseVerifyFraudProof, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.ResponseVerifyFraudProof
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.RequestVerifyFraudProof) (*types.ResponseVerifyFraudProof, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.RequestVerifyFraudProof) *types.ResponseVerifyFraudProof); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseVerifyFraudProof)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.RequestVerifyFraudProof) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifyVoteExtension provides a mock function with given fields: _a0, _a1
func (_m *Application) VerifyVoteExtension(_a0 context.Context, _a1 *types.RequestVerifyVoteExtension) (*types.ResponseVerifyVoteExtension, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.ResponseVerifyVoteExtension
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.RequestVerifyVoteExtension) (*types.ResponseVerifyVoteExtension, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.RequestVerifyVoteExtension) *types.ResponseVerifyVoteExtension); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseVerifyVoteExtension)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.RequestVerifyVoteExtension) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewApplication interface {
	mock.TestingT
	Cleanup(func())
}

// NewApplication creates a new instance of Application. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewApplication(t mockConstructorTestingTNewApplication) *Application {
	mock := &Application{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
