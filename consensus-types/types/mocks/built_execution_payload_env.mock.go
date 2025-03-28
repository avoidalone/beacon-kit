// Code generated by mockery v2.49.0. DO NOT EDIT.

package mocks

import (
	types "github.com/berachain/beacon-kit/consensus-types/types"
	uint256 "github.com/holiman/uint256"
	mock "github.com/stretchr/testify/mock"
)

// BuiltExecutionPayloadEnv is an autogenerated mock type for the BuiltExecutionPayloadEnv type
type BuiltExecutionPayloadEnv struct {
	mock.Mock
}

type BuiltExecutionPayloadEnv_Expecter struct {
	mock *mock.Mock
}

func (_m *BuiltExecutionPayloadEnv) EXPECT() *BuiltExecutionPayloadEnv_Expecter {
	return &BuiltExecutionPayloadEnv_Expecter{mock: &_m.Mock}
}

// GetBlobsBundle provides a mock function with given fields:
func (_m *BuiltExecutionPayloadEnv) GetBlobsBundle() types.BlobsBundle {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetBlobsBundle")
	}

	var r0 types.BlobsBundle
	if rf, ok := ret.Get(0).(func() types.BlobsBundle); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.BlobsBundle)
		}
	}

	return r0
}

// BuiltExecutionPayloadEnv_GetBlobsBundle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBlobsBundle'
type BuiltExecutionPayloadEnv_GetBlobsBundle_Call struct {
	*mock.Call
}

// GetBlobsBundle is a helper method to define mock.On call
func (_e *BuiltExecutionPayloadEnv_Expecter) GetBlobsBundle() *BuiltExecutionPayloadEnv_GetBlobsBundle_Call {
	return &BuiltExecutionPayloadEnv_GetBlobsBundle_Call{Call: _e.mock.On("GetBlobsBundle")}
}

func (_c *BuiltExecutionPayloadEnv_GetBlobsBundle_Call) Run(run func()) *BuiltExecutionPayloadEnv_GetBlobsBundle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BuiltExecutionPayloadEnv_GetBlobsBundle_Call) Return(_a0 types.BlobsBundle) *BuiltExecutionPayloadEnv_GetBlobsBundle_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BuiltExecutionPayloadEnv_GetBlobsBundle_Call) RunAndReturn(run func() types.BlobsBundle) *BuiltExecutionPayloadEnv_GetBlobsBundle_Call {
	_c.Call.Return(run)
	return _c
}

// GetExecutionPayload provides a mock function with given fields:
func (_m *BuiltExecutionPayloadEnv) GetExecutionPayload() *types.ExecutionPayload {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetExecutionPayload")
	}

	var r0 *types.ExecutionPayload
	if rf, ok := ret.Get(0).(func() *types.ExecutionPayload); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ExecutionPayload)
		}
	}

	return r0
}

// BuiltExecutionPayloadEnv_GetExecutionPayload_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetExecutionPayload'
type BuiltExecutionPayloadEnv_GetExecutionPayload_Call struct {
	*mock.Call
}

// GetExecutionPayload is a helper method to define mock.On call
func (_e *BuiltExecutionPayloadEnv_Expecter) GetExecutionPayload() *BuiltExecutionPayloadEnv_GetExecutionPayload_Call {
	return &BuiltExecutionPayloadEnv_GetExecutionPayload_Call{Call: _e.mock.On("GetExecutionPayload")}
}

func (_c *BuiltExecutionPayloadEnv_GetExecutionPayload_Call) Run(run func()) *BuiltExecutionPayloadEnv_GetExecutionPayload_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BuiltExecutionPayloadEnv_GetExecutionPayload_Call) Return(_a0 *types.ExecutionPayload) *BuiltExecutionPayloadEnv_GetExecutionPayload_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BuiltExecutionPayloadEnv_GetExecutionPayload_Call) RunAndReturn(run func() *types.ExecutionPayload) *BuiltExecutionPayloadEnv_GetExecutionPayload_Call {
	_c.Call.Return(run)
	return _c
}

// GetValue provides a mock function with given fields:
func (_m *BuiltExecutionPayloadEnv) GetValue() *uint256.Int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetValue")
	}

	var r0 *uint256.Int
	if rf, ok := ret.Get(0).(func() *uint256.Int); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*uint256.Int)
		}
	}

	return r0
}

// BuiltExecutionPayloadEnv_GetValue_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetValue'
type BuiltExecutionPayloadEnv_GetValue_Call struct {
	*mock.Call
}

// GetValue is a helper method to define mock.On call
func (_e *BuiltExecutionPayloadEnv_Expecter) GetValue() *BuiltExecutionPayloadEnv_GetValue_Call {
	return &BuiltExecutionPayloadEnv_GetValue_Call{Call: _e.mock.On("GetValue")}
}

func (_c *BuiltExecutionPayloadEnv_GetValue_Call) Run(run func()) *BuiltExecutionPayloadEnv_GetValue_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BuiltExecutionPayloadEnv_GetValue_Call) Return(_a0 *uint256.Int) *BuiltExecutionPayloadEnv_GetValue_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BuiltExecutionPayloadEnv_GetValue_Call) RunAndReturn(run func() *uint256.Int) *BuiltExecutionPayloadEnv_GetValue_Call {
	_c.Call.Return(run)
	return _c
}

// ShouldOverrideBuilder provides a mock function with given fields:
func (_m *BuiltExecutionPayloadEnv) ShouldOverrideBuilder() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ShouldOverrideBuilder")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// BuiltExecutionPayloadEnv_ShouldOverrideBuilder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ShouldOverrideBuilder'
type BuiltExecutionPayloadEnv_ShouldOverrideBuilder_Call struct {
	*mock.Call
}

// ShouldOverrideBuilder is a helper method to define mock.On call
func (_e *BuiltExecutionPayloadEnv_Expecter) ShouldOverrideBuilder() *BuiltExecutionPayloadEnv_ShouldOverrideBuilder_Call {
	return &BuiltExecutionPayloadEnv_ShouldOverrideBuilder_Call{Call: _e.mock.On("ShouldOverrideBuilder")}
}

func (_c *BuiltExecutionPayloadEnv_ShouldOverrideBuilder_Call) Run(run func()) *BuiltExecutionPayloadEnv_ShouldOverrideBuilder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BuiltExecutionPayloadEnv_ShouldOverrideBuilder_Call) Return(_a0 bool) *BuiltExecutionPayloadEnv_ShouldOverrideBuilder_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BuiltExecutionPayloadEnv_ShouldOverrideBuilder_Call) RunAndReturn(run func() bool) *BuiltExecutionPayloadEnv_ShouldOverrideBuilder_Call {
	_c.Call.Return(run)
	return _c
}

// NewBuiltExecutionPayloadEnv creates a new instance of BuiltExecutionPayloadEnv. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBuiltExecutionPayloadEnv(t interface {
	mock.TestingT
	Cleanup(func())
}) *BuiltExecutionPayloadEnv {
	mock := &BuiltExecutionPayloadEnv{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
