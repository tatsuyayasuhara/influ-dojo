// Code generated by MockGen. DO NOT EDIT.
// Source: follower.go

// Package client is a generated GoMock package.
package client

import (
	gomock "github.com/golang/mock/gomock"
	model "influ-dojo/api/domain/model"
	reflect "reflect"
)

// MockFollower is a mock of Follower interface.
type MockFollower struct {
	ctrl     *gomock.Controller
	recorder *MockFollowerMockRecorder
}

// MockFollowerMockRecorder is the mock recorder for MockFollower.
type MockFollowerMockRecorder struct {
	mock *MockFollower
}

// NewMockFollower creates a new mock instance.
func NewMockFollower(ctrl *gomock.Controller) *MockFollower {
	mock := &MockFollower{ctrl: ctrl}
	mock.recorder = &MockFollowerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFollower) EXPECT() *MockFollowerMockRecorder {
	return m.recorder
}

// CountFollowers mocks base method.
func (m *MockFollower) CountFollowers() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountFollowers")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountFollowers indicates an expected call of CountFollowers.
func (mr *MockFollowerMockRecorder) CountFollowers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountFollowers", reflect.TypeOf((*MockFollower)(nil).CountFollowers))
}

// GetFollowers mocks base method.
func (m *MockFollower) GetFollowers() ([]*model.Follower, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFollowers")
	ret0, _ := ret[0].([]*model.Follower)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFollowers indicates an expected call of GetFollowers.
func (mr *MockFollowerMockRecorder) GetFollowers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFollowers", reflect.TypeOf((*MockFollower)(nil).GetFollowers))
}
