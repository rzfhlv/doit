// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// IHandler is an autogenerated mock type for the IHandler type
type IHandler struct {
	mock.Mock
}

// Generate provides a mock function with given fields: e
func (_m *IHandler) Generate(e echo.Context) error {
	ret := _m.Called(e)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: e
func (_m *IHandler) GetAll(e echo.Context) error {
	ret := _m.Called(e)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByID provides a mock function with given fields: e
func (_m *IHandler) GetByID(e echo.Context) error {
	ret := _m.Called(e)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Migrate provides a mock function with given fields: e
func (_m *IHandler) Migrate(e echo.Context) error {
	ret := _m.Called(e)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewIHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewIHandler creates a new instance of IHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIHandler(t mockConstructorTestingTNewIHandler) *IHandler {
	mock := &IHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
