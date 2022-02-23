// Code generated by mockery v2.10.0. DO NOT EDIT.

package domain

import (
	context "context"

	model "github.com/bangumi/server/model"
	mock "github.com/stretchr/testify/mock"
)

// MockSubjectRepo is an autogenerated mock type for the SubjectRepo type
type MockSubjectRepo struct {
	mock.Mock
}

type MockSubjectRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSubjectRepo) EXPECT() *MockSubjectRepo_Expecter {
	return &MockSubjectRepo_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: ctx, id
func (_m *MockSubjectRepo) Get(ctx context.Context, id uint32) (model.Subject, error) {
	ret := _m.Called(ctx, id)

	var r0 model.Subject
	if rf, ok := ret.Get(0).(func(context.Context, uint32) model.Subject); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.Subject)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint32) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSubjectRepo_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockSubjectRepo_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//  - ctx context.Context
//  - id uint32
func (_e *MockSubjectRepo_Expecter) Get(ctx interface{}, id interface{}) *MockSubjectRepo_Get_Call {
	return &MockSubjectRepo_Get_Call{Call: _e.mock.On("Get", ctx, id)}
}

func (_c *MockSubjectRepo_Get_Call) Run(run func(ctx context.Context, id uint32)) *MockSubjectRepo_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32))
	})
	return _c
}

func (_c *MockSubjectRepo_Get_Call) Return(_a0 model.Subject, _a1 error) *MockSubjectRepo_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetPersonRelated provides a mock function with given fields: ctx, personID
func (_m *MockSubjectRepo) GetPersonRelated(ctx context.Context, personID uint32) ([]model.Subject, []model.PersonSubjectRelation, error) {
	ret := _m.Called(ctx, personID)

	var r0 []model.Subject
	if rf, ok := ret.Get(0).(func(context.Context, uint32) []model.Subject); ok {
		r0 = rf(ctx, personID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Subject)
		}
	}

	var r1 []model.PersonSubjectRelation
	if rf, ok := ret.Get(1).(func(context.Context, uint32) []model.PersonSubjectRelation); ok {
		r1 = rf(ctx, personID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]model.PersonSubjectRelation)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, uint32) error); ok {
		r2 = rf(ctx, personID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockSubjectRepo_GetPersonRelated_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPersonRelated'
type MockSubjectRepo_GetPersonRelated_Call struct {
	*mock.Call
}

// GetPersonRelated is a helper method to define mock.On call
//  - ctx context.Context
//  - personID uint32
func (_e *MockSubjectRepo_Expecter) GetPersonRelated(ctx interface{}, personID interface{}) *MockSubjectRepo_GetPersonRelated_Call {
	return &MockSubjectRepo_GetPersonRelated_Call{Call: _e.mock.On("GetPersonRelated", ctx, personID)}
}

func (_c *MockSubjectRepo_GetPersonRelated_Call) Run(run func(ctx context.Context, personID uint32)) *MockSubjectRepo_GetPersonRelated_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32))
	})
	return _c
}

func (_c *MockSubjectRepo_GetPersonRelated_Call) Return(_a0 []model.Subject, _a1 []model.PersonSubjectRelation, _a2 error) *MockSubjectRepo_GetPersonRelated_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}