// Code generated by MockGen. DO NOT EDIT.
// Source: bot.go

// Package client is a generated GoMock package.
package client

import (
	gomock "github.com/golang/mock/gomock"
	model "influ-dojo/api/domain/model"
	reflect "reflect"
)

// MockBot is a mock of Bot interface.
type MockBot struct {
	ctrl     *gomock.Controller
	recorder *MockBotMockRecorder
}

// MockBotMockRecorder is the mock recorder for MockBot.
type MockBotMockRecorder struct {
	mock *MockBot
}

// NewMockBot creates a new mock instance.
func NewMockBot(ctrl *gomock.Controller) *MockBot {
	mock := &MockBot{ctrl: ctrl}
	mock.recorder = &MockBotMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBot) EXPECT() *MockBotMockRecorder {
	return m.recorder
}

// Tweet mocks base method.
func (m *MockBot) Tweet(work []*model.Work, result []*model.Result, path string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tweet", work, result, path)
	ret0, _ := ret[0].(error)
	return ret0
}

// Tweet indicates an expected call of Tweet.
func (mr *MockBotMockRecorder) Tweet(work, result, path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tweet", reflect.TypeOf((*MockBot)(nil).Tweet), work, result, path)
}

// Favorite mocks base method.
func (m *MockBot) Favorite() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Favorite")
	ret0, _ := ret[0].(error)
	return ret0
}

// Favorite indicates an expected call of Favorite.
func (mr *MockBotMockRecorder) Favorite() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Favorite", reflect.TypeOf((*MockBot)(nil).Favorite))
}
