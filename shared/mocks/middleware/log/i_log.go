// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// ILog is an autogenerated mock type for the ILog type
type ILog struct {
	mock.Mock
}

// Logrus provides a mock function with given fields: next
func (_m *ILog) Logrus(next echo.HandlerFunc) echo.HandlerFunc {
	ret := _m.Called(next)

	var r0 echo.HandlerFunc
	if rf, ok := ret.Get(0).(func(echo.HandlerFunc) echo.HandlerFunc); ok {
		r0 = rf(next)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(echo.HandlerFunc)
		}
	}

	return r0
}

// NewILog creates a new instance of ILog. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewILog(t interface {
	mock.TestingT
	Cleanup(func())
}) *ILog {
	mock := &ILog{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
