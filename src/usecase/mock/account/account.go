// Code generated by MockGen. DO NOT EDIT.
// Source: src/usecase/account/account.go

// Package mock_account is a generated GoMock package.
package mock_account

import (
	reflect "reflect"

	model "github.com/achwanyusuf/carrent-accountsvc/src/model"
	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
)

// MockAccountInterface is a mock of AccountInterface interface.
type MockAccountInterface struct {
	ctrl     *gomock.Controller
	recorder *MockAccountInterfaceMockRecorder
}

// MockAccountInterfaceMockRecorder is the mock recorder for MockAccountInterface.
type MockAccountInterfaceMockRecorder struct {
	mock *MockAccountInterface
}

// NewMockAccountInterface creates a new mock instance.
func NewMockAccountInterface(ctrl *gomock.Controller) *MockAccountInterface {
	mock := &MockAccountInterface{ctrl: ctrl}
	mock.recorder = &MockAccountInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountInterface) EXPECT() *MockAccountInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAccountInterface) Create(ctx *gin.Context, v model.Register) (model.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, v)
	ret0, _ := ret[0].(model.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockAccountInterfaceMockRecorder) Create(ctx, v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAccountInterface)(nil).Create), ctx, v)
}

// DeleteByID mocks base method.
func (m *MockAccountInterface) DeleteByID(ctx *gin.Context, id int64, isHardDelete bool, vid int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", ctx, id, isHardDelete, vid)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockAccountInterfaceMockRecorder) DeleteByID(ctx, id, isHardDelete, vid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockAccountInterface)(nil).DeleteByID), ctx, id, isHardDelete, vid)
}

// GetByID mocks base method.
func (m *MockAccountInterface) GetByID(ctx *gin.Context, cacheControl string, id int64) (model.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, cacheControl, id)
	ret0, _ := ret[0].(model.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockAccountInterfaceMockRecorder) GetByID(ctx, cacheControl, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockAccountInterface)(nil).GetByID), ctx, cacheControl, id)
}

// GetByParam mocks base method.
func (m *MockAccountInterface) GetByParam(ctx *gin.Context, cacheControl string, v model.GetAccountsByParam) ([]model.Account, model.Pagination, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByParam", ctx, cacheControl, v)
	ret0, _ := ret[0].([]model.Account)
	ret1, _ := ret[1].(model.Pagination)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetByParam indicates an expected call of GetByParam.
func (mr *MockAccountInterfaceMockRecorder) GetByParam(ctx, cacheControl, v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByParam", reflect.TypeOf((*MockAccountInterface)(nil).GetByParam), ctx, cacheControl, v)
}

// Oauth2 mocks base method.
func (m *MockAccountInterface) Oauth2(ctx *gin.Context, v model.Login) (model.Auth, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Oauth2", ctx, v)
	ret0, _ := ret[0].(model.Auth)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Oauth2 indicates an expected call of Oauth2.
func (mr *MockAccountInterfaceMockRecorder) Oauth2(ctx, v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Oauth2", reflect.TypeOf((*MockAccountInterface)(nil).Oauth2), ctx, v)
}

// UpdateByID mocks base method.
func (m *MockAccountInterface) UpdateByID(ctx *gin.Context, id int64, v model.UpdateAccountData) (model.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateByID", ctx, id, v)
	ret0, _ := ret[0].(model.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateByID indicates an expected call of UpdateByID.
func (mr *MockAccountInterfaceMockRecorder) UpdateByID(ctx, id, v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateByID", reflect.TypeOf((*MockAccountInterface)(nil).UpdateByID), ctx, id, v)
}

// UpdatePasswordByID mocks base method.
func (m *MockAccountInterface) UpdatePasswordByID(ctx *gin.Context, id int64, v model.UpdatePasswordData) (model.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePasswordByID", ctx, id, v)
	ret0, _ := ret[0].(model.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePasswordByID indicates an expected call of UpdatePasswordByID.
func (mr *MockAccountInterfaceMockRecorder) UpdatePasswordByID(ctx, id, v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePasswordByID", reflect.TypeOf((*MockAccountInterface)(nil).UpdatePasswordByID), ctx, id, v)
}
