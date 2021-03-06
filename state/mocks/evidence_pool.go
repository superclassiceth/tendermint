// Code generated by mockery v2.3.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	abcitypes "github.com/tendermint/tendermint/abci/types"

	state "github.com/tendermint/tendermint/state"

	types "github.com/tendermint/tendermint/types"
)

// EvidencePool is an autogenerated mock type for the EvidencePool type
type EvidencePool struct {
	mock.Mock
}

// ABCIEvidence provides a mock function with given fields: _a0, _a1
func (_m *EvidencePool) ABCIEvidence(_a0 int64, _a1 []types.Evidence) []abcitypes.Evidence {
	ret := _m.Called(_a0, _a1)

	var r0 []abcitypes.Evidence
	if rf, ok := ret.Get(0).(func(int64, []types.Evidence) []abcitypes.Evidence); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]abcitypes.Evidence)
		}
	}

	return r0
}

// AddEvidence provides a mock function with given fields: _a0
func (_m *EvidencePool) AddEvidence(_a0 types.Evidence) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Evidence) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CheckEvidence provides a mock function with given fields: _a0
func (_m *EvidencePool) CheckEvidence(_a0 types.EvidenceList) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.EvidenceList) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PendingEvidence provides a mock function with given fields: _a0
func (_m *EvidencePool) PendingEvidence(_a0 int64) ([]types.Evidence, int64) {
	ret := _m.Called(_a0)

	var r0 []types.Evidence
	if rf, ok := ret.Get(0).(func(int64) []types.Evidence); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Evidence)
		}
	}

	var r1 int64
	if rf, ok := ret.Get(1).(func(int64) int64); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Get(1).(int64)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0
func (_m *EvidencePool) Update(_a0 state.State) {
	_m.Called(_a0)
}
