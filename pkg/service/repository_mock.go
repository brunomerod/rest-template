// Code generated by MockGen. DO NOT EDIT.
// Source: rest-template/pkg/service (interfaces: Repository)

// Package service is a generated GoMock package.
package service

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// GetAccountById mocks base method.
func (m *MockRepository) GetAccountById(arg0 context.Context, arg1 primitive.ObjectID) ([]Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountById", arg0, arg1)
	ret0, _ := ret[0].([]Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountById indicates an expected call of GetAccountById.
func (mr *MockRepositoryMockRecorder) GetAccountById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountById", reflect.TypeOf((*MockRepository)(nil).GetAccountById), arg0, arg1)
}

// GetAccounts mocks base method.
func (m *MockRepository) GetAccounts(arg0 context.Context) ([]Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccounts", arg0)
	ret0, _ := ret[0].([]Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccounts indicates an expected call of GetAccounts.
func (mr *MockRepositoryMockRecorder) GetAccounts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccounts", reflect.TypeOf((*MockRepository)(nil).GetAccounts), arg0)
}

// InsertAccount mocks base method.
func (m *MockRepository) InsertAccount(arg0 context.Context, arg1 Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertAccount", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertAccount indicates an expected call of InsertAccount.
func (mr *MockRepositoryMockRecorder) InsertAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertAccount", reflect.TypeOf((*MockRepository)(nil).InsertAccount), arg0, arg1)
}

// InsertTransaction mocks base method.
func (m *MockRepository) InsertTransaction(arg0 context.Context, arg1 Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertTransaction", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertTransaction indicates an expected call of InsertTransaction.
func (mr *MockRepositoryMockRecorder) InsertTransaction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertTransaction", reflect.TypeOf((*MockRepository)(nil).InsertTransaction), arg0, arg1)
}
