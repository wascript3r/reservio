// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/features/reservation/repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	models "github.com/wascript3r/reservio/pkg/features/reservation/models"
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

// Delete mocks base method.
func (m *MockRepository) Delete(ctx context.Context, companyID, serviceID, reservationID, clientID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, companyID, serviceID, reservationID, clientID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRepositoryMockRecorder) Delete(ctx, companyID, serviceID, reservationID, clientID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), ctx, companyID, serviceID, reservationID, clientID)
}

// DeleteByCompany mocks base method.
func (m *MockRepository) DeleteByCompany(ctx context.Context, companyID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByCompany", ctx, companyID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByCompany indicates an expected call of DeleteByCompany.
func (mr *MockRepositoryMockRecorder) DeleteByCompany(ctx, companyID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByCompany", reflect.TypeOf((*MockRepository)(nil).DeleteByCompany), ctx, companyID)
}

// DeleteByService mocks base method.
func (m *MockRepository) DeleteByService(ctx context.Context, serviceID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByService", ctx, serviceID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByService indicates an expected call of DeleteByService.
func (mr *MockRepositoryMockRecorder) DeleteByService(ctx, serviceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByService", reflect.TypeOf((*MockRepository)(nil).DeleteByService), ctx, serviceID)
}

// Exists mocks base method.
func (m *MockRepository) Exists(ctx context.Context, companyID, serviceID string, date time.Time) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, companyID, serviceID, date)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockRepositoryMockRecorder) Exists(ctx, companyID, serviceID, date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockRepository)(nil).Exists), ctx, companyID, serviceID, date)
}

// Get mocks base method.
func (m *MockRepository) Get(ctx context.Context, companyID, serviceID, reservationID string, clientID *string, onlyApprovedCompany bool) (*models.FullReservation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, companyID, serviceID, reservationID, clientID, onlyApprovedCompany)
	ret0, _ := ret[0].(*models.FullReservation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockRepositoryMockRecorder) Get(ctx, companyID, serviceID, reservationID, clientID, onlyApprovedCompany interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRepository)(nil).Get), ctx, companyID, serviceID, reservationID, clientID, onlyApprovedCompany)
}

// GetAll mocks base method.
func (m *MockRepository) GetAll(ctx context.Context, companyID, serviceID string, onlyApprovedCompany bool) ([]*models.FullReservation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx, companyID, serviceID, onlyApprovedCompany)
	ret0, _ := ret[0].([]*models.FullReservation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockRepositoryMockRecorder) GetAll(ctx, companyID, serviceID, onlyApprovedCompany interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockRepository)(nil).GetAll), ctx, companyID, serviceID, onlyApprovedCompany)
}

// GetAllByClient mocks base method.
func (m *MockRepository) GetAllByClient(ctx context.Context, clientID string) ([]*models.ClientReservation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllByClient", ctx, clientID)
	ret0, _ := ret[0].([]*models.ClientReservation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllByClient indicates an expected call of GetAllByClient.
func (mr *MockRepositoryMockRecorder) GetAllByClient(ctx, clientID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllByClient", reflect.TypeOf((*MockRepository)(nil).GetAllByClient), ctx, clientID)
}

// GetAllByCompany mocks base method.
func (m *MockRepository) GetAllByCompany(ctx context.Context, companyID string, onlyApprovedCompany bool) ([]*models.FullReservation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllByCompany", ctx, companyID, onlyApprovedCompany)
	ret0, _ := ret[0].([]*models.FullReservation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllByCompany indicates an expected call of GetAllByCompany.
func (mr *MockRepositoryMockRecorder) GetAllByCompany(ctx, companyID, onlyApprovedCompany interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllByCompany", reflect.TypeOf((*MockRepository)(nil).GetAllByCompany), ctx, companyID, onlyApprovedCompany)
}

// Insert mocks base method.
func (m *MockRepository) Insert(ctx context.Context, rs *models.Reservation) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, rs)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockRepositoryMockRecorder) Insert(ctx, rs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockRepository)(nil).Insert), ctx, rs)
}

// Update mocks base method.
func (m *MockRepository) Update(ctx context.Context, companyID, serviceID, reservationID, clientID string, ru *models.ReservationUpdate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, companyID, serviceID, reservationID, clientID, ru)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockRepositoryMockRecorder) Update(ctx, companyID, serviceID, reservationID, clientID, ru interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRepository)(nil).Update), ctx, companyID, serviceID, reservationID, clientID, ru)
}