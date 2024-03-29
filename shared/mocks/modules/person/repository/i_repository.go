// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/rzfhlv/doit/modules/person/model"
	mock "github.com/stretchr/testify/mock"

	param "github.com/rzfhlv/doit/utilities/param"
)

// IRepository is an autogenerated mock type for the IRepository type
type IRepository struct {
	mock.Mock
}

// Count provides a mock function with given fields: ctx
func (_m *IRepository) Count(ctx context.Context) (int64, error) {
	ret := _m.Called(ctx)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (int64, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) int64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: ctx, filter
func (_m *IRepository) GetAll(ctx context.Context, filter param.Param) ([]model.Person, error) {
	ret := _m.Called(ctx, filter)

	var r0 []model.Person
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, param.Param) ([]model.Person, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, param.Param) []model.Person); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Person)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, param.Param) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *IRepository) GetByID(ctx context.Context, id int64) (model.Person, error) {
	ret := _m.Called(ctx, id)

	var r0 model.Person
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (model.Person, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) model.Person); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.Person)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIRepository creates a new instance of IRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *IRepository {
	mock := &IRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
