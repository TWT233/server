// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/bangumi/server/internal/domain"
	mock "github.com/stretchr/testify/mock"

	model "github.com/bangumi/server/internal/model"
)

// TopicRepo is an autogenerated mock type for the TopicRepo type
type TopicRepo struct {
	mock.Mock
}

type TopicRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *TopicRepo) EXPECT() *TopicRepo_Expecter {
	return &TopicRepo_Expecter{mock: &_m.Mock}
}

// Count provides a mock function with given fields: ctx, topicType, id, displays
func (_m *TopicRepo) Count(ctx context.Context, topicType domain.TopicType, id uint32, displays []model.TopicDisplay) (int64, error) {
	ret := _m.Called(ctx, topicType, id, displays)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, domain.TopicType, uint32, []model.TopicDisplay) int64); ok {
		r0 = rf(ctx, topicType, id, displays)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.TopicType, uint32, []model.TopicDisplay) error); ok {
		r1 = rf(ctx, topicType, id, displays)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TopicRepo_Count_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Count'
type TopicRepo_Count_Call struct {
	*mock.Call
}

// Count is a helper method to define mock.On call
//   - ctx context.Context
//   - topicType domain.TopicType
//   - id uint32
//   - displays []model.TopicDisplay
func (_e *TopicRepo_Expecter) Count(ctx interface{}, topicType interface{}, id interface{}, displays interface{}) *TopicRepo_Count_Call {
	return &TopicRepo_Count_Call{Call: _e.mock.On("Count", ctx, topicType, id, displays)}
}

func (_c *TopicRepo_Count_Call) Run(run func(ctx context.Context, topicType domain.TopicType, id uint32, displays []model.TopicDisplay)) *TopicRepo_Count_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.TopicType), args[2].(uint32), args[3].([]model.TopicDisplay))
	})
	return _c
}

func (_c *TopicRepo_Count_Call) Return(_a0 int64, _a1 error) *TopicRepo_Count_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// CountReplies provides a mock function with given fields: ctx, commentType, id
func (_m *TopicRepo) CountReplies(ctx context.Context, commentType domain.CommentType, id model.TopicID) (int64, error) {
	ret := _m.Called(ctx, commentType, id)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, domain.CommentType, model.TopicID) int64); ok {
		r0 = rf(ctx, commentType, id)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.CommentType, model.TopicID) error); ok {
		r1 = rf(ctx, commentType, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TopicRepo_CountReplies_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CountReplies'
type TopicRepo_CountReplies_Call struct {
	*mock.Call
}

// CountReplies is a helper method to define mock.On call
//   - ctx context.Context
//   - commentType domain.CommentType
//   - id model.TopicID
func (_e *TopicRepo_Expecter) CountReplies(ctx interface{}, commentType interface{}, id interface{}) *TopicRepo_CountReplies_Call {
	return &TopicRepo_CountReplies_Call{Call: _e.mock.On("CountReplies", ctx, commentType, id)}
}

func (_c *TopicRepo_CountReplies_Call) Run(run func(ctx context.Context, commentType domain.CommentType, id model.TopicID)) *TopicRepo_CountReplies_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.CommentType), args[2].(model.TopicID))
	})
	return _c
}

func (_c *TopicRepo_CountReplies_Call) Return(_a0 int64, _a1 error) *TopicRepo_CountReplies_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Get provides a mock function with given fields: ctx, topicType, id
func (_m *TopicRepo) Get(ctx context.Context, topicType domain.TopicType, id model.TopicID) (model.Topic, error) {
	ret := _m.Called(ctx, topicType, id)

	var r0 model.Topic
	if rf, ok := ret.Get(0).(func(context.Context, domain.TopicType, model.TopicID) model.Topic); ok {
		r0 = rf(ctx, topicType, id)
	} else {
		r0 = ret.Get(0).(model.Topic)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.TopicType, model.TopicID) error); ok {
		r1 = rf(ctx, topicType, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TopicRepo_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type TopicRepo_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - topicType domain.TopicType
//   - id model.TopicID
func (_e *TopicRepo_Expecter) Get(ctx interface{}, topicType interface{}, id interface{}) *TopicRepo_Get_Call {
	return &TopicRepo_Get_Call{Call: _e.mock.On("Get", ctx, topicType, id)}
}

func (_c *TopicRepo_Get_Call) Run(run func(ctx context.Context, topicType domain.TopicType, id model.TopicID)) *TopicRepo_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.TopicType), args[2].(model.TopicID))
	})
	return _c
}

func (_c *TopicRepo_Get_Call) Return(_a0 model.Topic, _a1 error) *TopicRepo_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetTopicContent provides a mock function with given fields: ctx, topicType, id
func (_m *TopicRepo) GetTopicContent(ctx context.Context, topicType domain.TopicType, id model.TopicID) (model.Comment, error) {
	ret := _m.Called(ctx, topicType, id)

	var r0 model.Comment
	if rf, ok := ret.Get(0).(func(context.Context, domain.TopicType, model.TopicID) model.Comment); ok {
		r0 = rf(ctx, topicType, id)
	} else {
		r0 = ret.Get(0).(model.Comment)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.TopicType, model.TopicID) error); ok {
		r1 = rf(ctx, topicType, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TopicRepo_GetTopicContent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTopicContent'
type TopicRepo_GetTopicContent_Call struct {
	*mock.Call
}

// GetTopicContent is a helper method to define mock.On call
//   - ctx context.Context
//   - topicType domain.TopicType
//   - id model.TopicID
func (_e *TopicRepo_Expecter) GetTopicContent(ctx interface{}, topicType interface{}, id interface{}) *TopicRepo_GetTopicContent_Call {
	return &TopicRepo_GetTopicContent_Call{Call: _e.mock.On("GetTopicContent", ctx, topicType, id)}
}

func (_c *TopicRepo_GetTopicContent_Call) Run(run func(ctx context.Context, topicType domain.TopicType, id model.TopicID)) *TopicRepo_GetTopicContent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.TopicType), args[2].(model.TopicID))
	})
	return _c
}

func (_c *TopicRepo_GetTopicContent_Call) Return(_a0 model.Comment, _a1 error) *TopicRepo_GetTopicContent_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// List provides a mock function with given fields: ctx, topicType, id, displays, limit, offset
func (_m *TopicRepo) List(ctx context.Context, topicType domain.TopicType, id uint32, displays []model.TopicDisplay, limit int, offset int) ([]model.Topic, error) {
	ret := _m.Called(ctx, topicType, id, displays, limit, offset)

	var r0 []model.Topic
	if rf, ok := ret.Get(0).(func(context.Context, domain.TopicType, uint32, []model.TopicDisplay, int, int) []model.Topic); ok {
		r0 = rf(ctx, topicType, id, displays, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Topic)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.TopicType, uint32, []model.TopicDisplay, int, int) error); ok {
		r1 = rf(ctx, topicType, id, displays, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TopicRepo_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type TopicRepo_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - topicType domain.TopicType
//   - id uint32
//   - displays []model.TopicDisplay
//   - limit int
//   - offset int
func (_e *TopicRepo_Expecter) List(ctx interface{}, topicType interface{}, id interface{}, displays interface{}, limit interface{}, offset interface{}) *TopicRepo_List_Call {
	return &TopicRepo_List_Call{Call: _e.mock.On("List", ctx, topicType, id, displays, limit, offset)}
}

func (_c *TopicRepo_List_Call) Run(run func(ctx context.Context, topicType domain.TopicType, id uint32, displays []model.TopicDisplay, limit int, offset int)) *TopicRepo_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.TopicType), args[2].(uint32), args[3].([]model.TopicDisplay), args[4].(int), args[5].(int))
	})
	return _c
}

func (_c *TopicRepo_List_Call) Return(_a0 []model.Topic, _a1 error) *TopicRepo_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// ListReplies provides a mock function with given fields: ctx, commentType, id, limit, offset
func (_m *TopicRepo) ListReplies(ctx context.Context, commentType domain.CommentType, id model.TopicID, limit int, offset int) ([]model.Comment, error) {
	ret := _m.Called(ctx, commentType, id, limit, offset)

	var r0 []model.Comment
	if rf, ok := ret.Get(0).(func(context.Context, domain.CommentType, model.TopicID, int, int) []model.Comment); ok {
		r0 = rf(ctx, commentType, id, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Comment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.CommentType, model.TopicID, int, int) error); ok {
		r1 = rf(ctx, commentType, id, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TopicRepo_ListReplies_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListReplies'
type TopicRepo_ListReplies_Call struct {
	*mock.Call
}

// ListReplies is a helper method to define mock.On call
//   - ctx context.Context
//   - commentType domain.CommentType
//   - id model.TopicID
//   - limit int
//   - offset int
func (_e *TopicRepo_Expecter) ListReplies(ctx interface{}, commentType interface{}, id interface{}, limit interface{}, offset interface{}) *TopicRepo_ListReplies_Call {
	return &TopicRepo_ListReplies_Call{Call: _e.mock.On("ListReplies", ctx, commentType, id, limit, offset)}
}

func (_c *TopicRepo_ListReplies_Call) Run(run func(ctx context.Context, commentType domain.CommentType, id model.TopicID, limit int, offset int)) *TopicRepo_ListReplies_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.CommentType), args[2].(model.TopicID), args[3].(int), args[4].(int))
	})
	return _c
}

func (_c *TopicRepo_ListReplies_Call) Return(_a0 []model.Comment, _a1 error) *TopicRepo_ListReplies_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewTopicRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewTopicRepo creates a new instance of TopicRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTopicRepo(t mockConstructorTestingTNewTopicRepo) *TopicRepo {
	mock := &TopicRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
