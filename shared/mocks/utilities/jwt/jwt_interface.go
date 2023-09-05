// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	jwt "github.com/rzfhlv/doit/utilities/jwt"
	mock "github.com/stretchr/testify/mock"
)

// JWTInterface is an autogenerated mock type for the JWTInterface type
type JWTInterface struct {
	mock.Mock
}

// Generate provides a mock function with given fields: id, username, email
func (_m *JWTInterface) Generate(id int64, username string, email string) (string, error) {
	ret := _m.Called(id, username, email)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(int64, string, string) (string, error)); ok {
		return rf(id, username, email)
	}
	if rf, ok := ret.Get(0).(func(int64, string, string) string); ok {
		r0 = rf(id, username, email)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(int64, string, string) error); ok {
		r1 = rf(id, username, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateToken provides a mock function with given fields: signedToken
func (_m *JWTInterface) ValidateToken(signedToken string) (*jwt.JWTClaim, error) {
	ret := _m.Called(signedToken)

	var r0 *jwt.JWTClaim
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*jwt.JWTClaim, error)); ok {
		return rf(signedToken)
	}
	if rf, ok := ret.Get(0).(func(string) *jwt.JWTClaim); ok {
		r0 = rf(signedToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*jwt.JWTClaim)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(signedToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewJWTInterface creates a new instance of JWTInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewJWTInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *JWTInterface {
	mock := &JWTInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}