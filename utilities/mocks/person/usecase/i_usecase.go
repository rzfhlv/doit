// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/rzfhlv/doit/modules/person/model"
	mock "github.com/stretchr/testify/mock"

	param "github.com/rzfhlv/doit/utilities/param"
)

// IUsecase is an autogenerated mock type for the IUsecase type
type IUsecase struct {
	mock.Mock
}

// GetAll provides a mock function with given fields: ctx, _a1
func (_m *IUsecase) GetAll(ctx context.Context, _a1 *param.Param) ([]model.Person, error) {
	ret := _m.Called(ctx, _a1)

	var r0 []model.Person
	if rf, ok := ret.Get(0).(func(context.Context, *param.Param) []model.Person); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Person)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *param.Param) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *IUsecase) GetByID(ctx context.Context, id int64) (model.Person, error) {
	ret := _m.Called(ctx, id)

	var r0 model.Person
	if rf, ok := ret.Get(0).(func(context.Context, int64) model.Person); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.Person)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewIUsecase creates a new instance of IUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIUsecase(t mockConstructorTestingTNewIUsecase) *IUsecase {
	mock := &IUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}