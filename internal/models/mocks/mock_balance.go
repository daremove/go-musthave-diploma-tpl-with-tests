// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/daremove/go-musthave-diploma-tpl/tree/master/internal/models (interfaces: BalanceService)

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	models "github.com/daremove/go-musthave-diploma-tpl/tree/master/internal/models"
	gomock "github.com/golang/mock/gomock"
)

// MockBalanceService is a mock of BalanceService interface.
type MockBalanceService struct {
	ctrl     *gomock.Controller
	recorder *MockBalanceServiceMockRecorder
}

// MockBalanceServiceMockRecorder is the mock recorder for MockBalanceService.
type MockBalanceServiceMockRecorder struct {
	mock *MockBalanceService
}

// NewMockBalanceService creates a new mock instance.
func NewMockBalanceService(ctrl *gomock.Controller) *MockBalanceService {
	mock := &MockBalanceService{ctrl: ctrl}
	mock.recorder = &MockBalanceServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBalanceService) EXPECT() *MockBalanceServiceMockRecorder {
	return m.recorder
}

// CreateWithdrawal mocks base method.
func (m *MockBalanceService) CreateWithdrawal(arg0 context.Context, arg1, arg2 string, arg3 float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWithdrawal", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateWithdrawal indicates an expected call of CreateWithdrawal.
func (mr *MockBalanceServiceMockRecorder) CreateWithdrawal(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWithdrawal", reflect.TypeOf((*MockBalanceService)(nil).CreateWithdrawal), arg0, arg1, arg2, arg3)
}

// GetUserBalance mocks base method.
func (m *MockBalanceService) GetUserBalance(arg0 context.Context, arg1 string) (models.Balance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserBalance", arg0, arg1)
	ret0, _ := ret[0].(models.Balance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserBalance indicates an expected call of GetUserBalance.
func (mr *MockBalanceServiceMockRecorder) GetUserBalance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserBalance", reflect.TypeOf((*MockBalanceService)(nil).GetUserBalance), arg0, arg1)
}

// GetWithdrawalFlow mocks base method.
func (m *MockBalanceService) GetWithdrawalFlow(arg0 context.Context, arg1 string) ([]models.WithdrawalFlowItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWithdrawalFlow", arg0, arg1)
	ret0, _ := ret[0].([]models.WithdrawalFlowItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWithdrawalFlow indicates an expected call of GetWithdrawalFlow.
func (mr *MockBalanceServiceMockRecorder) GetWithdrawalFlow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWithdrawalFlow", reflect.TypeOf((*MockBalanceService)(nil).GetWithdrawalFlow), arg0, arg1)
}
