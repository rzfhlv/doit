// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// IHandler is an autogenerated mock type for the IHandler type
type IHandler struct {
	mock.Mock
}

// Login provides a mock function with given fields: e
func (_m *IHandler) Login(e echo.Context) error {
	ret := _m.Called(e)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Logout provides a mock function with given fields: e
func (_m *IHandler) Logout(e echo.Context) error {
	ret := _m.Called(e)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Register provides a mock function with given fields: e
func (_m *IHandler) Register(e echo.Context) error {
	ret := _m.Called(e)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Validate provides a mock function with given fields: e
func (_m *IHandler) Validate(e echo.Context) error {
	ret := _m.Called(e)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIHandler creates a new instance of IHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *IHandler {
	mock := &IHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
