// Code generated by mockery v2.10.0. DO NOT EDIT.

package domain

import (
	context "context"

	model "github.com/bangumi/server/model"
	mock "github.com/stretchr/testify/mock"
)

// MockEpisodeRepo is an autogenerated mock type for the EpisodeRepo type
type MockEpisodeRepo struct {
	mock.Mock
}

type MockEpisodeRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *MockEpisodeRepo) EXPECT() *MockEpisodeRepo_Expecter {
	return &MockEpisodeRepo_Expecter{mock: &_m.Mock}
}

// Count provides a mock function with given fields: ctx, subjectID
func (_m *MockEpisodeRepo) Count(ctx context.Context, subjectID uint32) (int64, error) {
	ret := _m.Called(ctx, subjectID)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, uint32) int64); ok {
		r0 = rf(ctx, subjectID)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint32) error); ok {
		r1 = rf(ctx, subjectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEpisodeRepo_Count_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Count'
type MockEpisodeRepo_Count_Call struct {
	*mock.Call
}

// Count is a helper method to define mock.On call
//  - ctx context.Context
//  - subjectID uint32
func (_e *MockEpisodeRepo_Expecter) Count(ctx interface{}, subjectID interface{}) *MockEpisodeRepo_Count_Call {
	return &MockEpisodeRepo_Count_Call{Call: _e.mock.On("Count", ctx, subjectID)}
}

func (_c *MockEpisodeRepo_Count_Call) Run(run func(ctx context.Context, subjectID uint32)) *MockEpisodeRepo_Count_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32))
	})
	return _c
}

func (_c *MockEpisodeRepo_Count_Call) Return(_a0 int64, _a1 error) *MockEpisodeRepo_Count_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// CountByType provides a mock function with given fields: ctx, subjectID, epType
func (_m *MockEpisodeRepo) CountByType(ctx context.Context, subjectID uint32, epType uint8) (int64, error) {
	ret := _m.Called(ctx, subjectID, epType)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, uint32, uint8) int64); ok {
		r0 = rf(ctx, subjectID, epType)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint32, uint8) error); ok {
		r1 = rf(ctx, subjectID, epType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEpisodeRepo_CountByType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CountByType'
type MockEpisodeRepo_CountByType_Call struct {
	*mock.Call
}

// CountByType is a helper method to define mock.On call
//  - ctx context.Context
//  - subjectID uint32
//  - epType uint8
func (_e *MockEpisodeRepo_Expecter) CountByType(ctx interface{}, subjectID interface{}, epType interface{}) *MockEpisodeRepo_CountByType_Call {
	return &MockEpisodeRepo_CountByType_Call{Call: _e.mock.On("CountByType", ctx, subjectID, epType)}
}

func (_c *MockEpisodeRepo_CountByType_Call) Run(run func(ctx context.Context, subjectID uint32, epType uint8)) *MockEpisodeRepo_CountByType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32), args[2].(uint8))
	})
	return _c
}

func (_c *MockEpisodeRepo_CountByType_Call) Return(_a0 int64, _a1 error) *MockEpisodeRepo_CountByType_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Get provides a mock function with given fields: ctx, episodeID
func (_m *MockEpisodeRepo) Get(ctx context.Context, episodeID uint32) (model.Episode, error) {
	ret := _m.Called(ctx, episodeID)

	var r0 model.Episode
	if rf, ok := ret.Get(0).(func(context.Context, uint32) model.Episode); ok {
		r0 = rf(ctx, episodeID)
	} else {
		r0 = ret.Get(0).(model.Episode)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint32) error); ok {
		r1 = rf(ctx, episodeID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEpisodeRepo_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockEpisodeRepo_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//  - ctx context.Context
//  - episodeID uint32
func (_e *MockEpisodeRepo_Expecter) Get(ctx interface{}, episodeID interface{}) *MockEpisodeRepo_Get_Call {
	return &MockEpisodeRepo_Get_Call{Call: _e.mock.On("Get", ctx, episodeID)}
}

func (_c *MockEpisodeRepo_Get_Call) Run(run func(ctx context.Context, episodeID uint32)) *MockEpisodeRepo_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32))
	})
	return _c
}

func (_c *MockEpisodeRepo_Get_Call) Return(_a0 model.Episode, _a1 error) *MockEpisodeRepo_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// List provides a mock function with given fields: ctx, subjectID, limit, offset
func (_m *MockEpisodeRepo) List(ctx context.Context, subjectID uint32, limit int, offset int) ([]model.Episode, error) {
	ret := _m.Called(ctx, subjectID, limit, offset)

	var r0 []model.Episode
	if rf, ok := ret.Get(0).(func(context.Context, uint32, int, int) []model.Episode); ok {
		r0 = rf(ctx, subjectID, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Episode)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint32, int, int) error); ok {
		r1 = rf(ctx, subjectID, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEpisodeRepo_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockEpisodeRepo_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//  - ctx context.Context
//  - subjectID uint32
//  - limit int
//  - offset int
func (_e *MockEpisodeRepo_Expecter) List(ctx interface{}, subjectID interface{}, limit interface{}, offset interface{}) *MockEpisodeRepo_List_Call {
	return &MockEpisodeRepo_List_Call{Call: _e.mock.On("List", ctx, subjectID, limit, offset)}
}

func (_c *MockEpisodeRepo_List_Call) Run(run func(ctx context.Context, subjectID uint32, limit int, offset int)) *MockEpisodeRepo_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32), args[2].(int), args[3].(int))
	})
	return _c
}

func (_c *MockEpisodeRepo_List_Call) Return(_a0 []model.Episode, _a1 error) *MockEpisodeRepo_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// ListByType provides a mock function with given fields: ctx, subjectID, epType, limit, offset
func (_m *MockEpisodeRepo) ListByType(ctx context.Context, subjectID uint32, epType uint8, limit int, offset int) ([]model.Episode, error) {
	ret := _m.Called(ctx, subjectID, epType, limit, offset)

	var r0 []model.Episode
	if rf, ok := ret.Get(0).(func(context.Context, uint32, uint8, int, int) []model.Episode); ok {
		r0 = rf(ctx, subjectID, epType, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Episode)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint32, uint8, int, int) error); ok {
		r1 = rf(ctx, subjectID, epType, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEpisodeRepo_ListByType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListByType'
type MockEpisodeRepo_ListByType_Call struct {
	*mock.Call
}

// ListByType is a helper method to define mock.On call
//  - ctx context.Context
//  - subjectID uint32
//  - epType uint8
//  - limit int
//  - offset int
func (_e *MockEpisodeRepo_Expecter) ListByType(ctx interface{}, subjectID interface{}, epType interface{}, limit interface{}, offset interface{}) *MockEpisodeRepo_ListByType_Call {
	return &MockEpisodeRepo_ListByType_Call{Call: _e.mock.On("ListByType", ctx, subjectID, epType, limit, offset)}
}

func (_c *MockEpisodeRepo_ListByType_Call) Run(run func(ctx context.Context, subjectID uint32, epType uint8, limit int, offset int)) *MockEpisodeRepo_ListByType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32), args[2].(uint8), args[3].(int), args[4].(int))
	})
	return _c
}

func (_c *MockEpisodeRepo_ListByType_Call) Return(_a0 []model.Episode, _a1 error) *MockEpisodeRepo_ListByType_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}