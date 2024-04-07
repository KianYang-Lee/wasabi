// Code generated by mockery v2.42.1. DO NOT EDIT.

//go:build !compile

package mocks

import (
	wasabi "github.com/ksysoev/wasabi"
	mock "github.com/stretchr/testify/mock"
)

// MockRequestHandler is an autogenerated mock type for the RequestHandler type
type MockRequestHandler struct {
	mock.Mock
}

type MockRequestHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRequestHandler) EXPECT() *MockRequestHandler_Expecter {
	return &MockRequestHandler_Expecter{mock: &_m.Mock}
}

// Handle provides a mock function with given fields: conn, req
func (_m *MockRequestHandler) Handle(conn wasabi.Connection, req wasabi.Request) error {
	ret := _m.Called(conn, req)

	if len(ret) == 0 {
		panic("no return value specified for Handle")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(wasabi.Connection, wasabi.Request) error); ok {
		r0 = rf(conn, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRequestHandler_Handle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Handle'
type MockRequestHandler_Handle_Call struct {
	*mock.Call
}

// Handle is a helper method to define mock.On call
//   - conn wasabi.Connection
//   - req wasabi.Request
func (_e *MockRequestHandler_Expecter) Handle(conn interface{}, req interface{}) *MockRequestHandler_Handle_Call {
	return &MockRequestHandler_Handle_Call{Call: _e.mock.On("Handle", conn, req)}
}

func (_c *MockRequestHandler_Handle_Call) Run(run func(conn wasabi.Connection, req wasabi.Request)) *MockRequestHandler_Handle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(wasabi.Connection), args[1].(wasabi.Request))
	})
	return _c
}

func (_c *MockRequestHandler_Handle_Call) Return(_a0 error) *MockRequestHandler_Handle_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRequestHandler_Handle_Call) RunAndReturn(run func(wasabi.Connection, wasabi.Request) error) *MockRequestHandler_Handle_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRequestHandler creates a new instance of MockRequestHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRequestHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRequestHandler {
	mock := &MockRequestHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
