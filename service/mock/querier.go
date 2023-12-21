// Code generated by MockGen. DO NOT EDIT.
// Source: service/store/querier.go
//
// Generated by this command:
//
//	mockgen -source=service/store/querier.go -destination=service/mock/querier.go -package=mock
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	store "pnb/service/store"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockQuerier is a mock of Querier interface.
type MockQuerier struct {
	ctrl     *gomock.Controller
	recorder *MockQuerierMockRecorder
}

// MockQuerierMockRecorder is the mock recorder for MockQuerier.
type MockQuerierMockRecorder struct {
	mock *MockQuerier
}

// NewMockQuerier creates a new mock instance.
func NewMockQuerier(ctrl *gomock.Controller) *MockQuerier {
	mock := &MockQuerier{ctrl: ctrl}
	mock.recorder = &MockQuerierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuerier) EXPECT() *MockQuerierMockRecorder {
	return m.recorder
}

// CreateHealth mocks base method.
func (m *MockQuerier) CreateHealth(ctx context.Context, arg store.CreateHealthParams) (store.Health, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateHealth", ctx, arg)
	ret0, _ := ret[0].(store.Health)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateHealth indicates an expected call of CreateHealth.
func (mr *MockQuerierMockRecorder) CreateHealth(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateHealth", reflect.TypeOf((*MockQuerier)(nil).CreateHealth), ctx, arg)
}

// CreateSource mocks base method.
func (m *MockQuerier) CreateSource(ctx context.Context, arg store.CreateSourceParams) (store.Source, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSource", ctx, arg)
	ret0, _ := ret[0].(store.Source)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSource indicates an expected call of CreateSource.
func (mr *MockQuerierMockRecorder) CreateSource(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSource", reflect.TypeOf((*MockQuerier)(nil).CreateSource), ctx, arg)
}

// DeleteSource mocks base method.
func (m *MockQuerier) DeleteSource(ctx context.Context, id string) (store.Source, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSource", ctx, id)
	ret0, _ := ret[0].(store.Source)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSource indicates an expected call of DeleteSource.
func (mr *MockQuerierMockRecorder) DeleteSource(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSource", reflect.TypeOf((*MockQuerier)(nil).DeleteSource), ctx, id)
}

// GetSource mocks base method.
func (m *MockQuerier) GetSource(ctx context.Context, id string) (store.Source, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSource", ctx, id)
	ret0, _ := ret[0].(store.Source)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSource indicates an expected call of GetSource.
func (mr *MockQuerierMockRecorder) GetSource(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSource", reflect.TypeOf((*MockQuerier)(nil).GetSource), ctx, id)
}

// ListSource mocks base method.
func (m *MockQuerier) ListSource(ctx context.Context, arg store.ListSourceParams) ([]store.Source, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSource", ctx, arg)
	ret0, _ := ret[0].([]store.Source)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSource indicates an expected call of ListSource.
func (mr *MockQuerierMockRecorder) ListSource(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSource", reflect.TypeOf((*MockQuerier)(nil).ListSource), ctx, arg)
}

// UpdateSource mocks base method.
func (m *MockQuerier) UpdateSource(ctx context.Context, arg store.UpdateSourceParams) (store.Source, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSource", ctx, arg)
	ret0, _ := ret[0].(store.Source)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSource indicates an expected call of UpdateSource.
func (mr *MockQuerierMockRecorder) UpdateSource(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSource", reflect.TypeOf((*MockQuerier)(nil).UpdateSource), ctx, arg)
}
