// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_servies is a generated GoMock package.
package mock_servies

import (
	models "cab-management-portal/app/models"
	servies "cab-management-portal/app/servies"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	sqlx "github.com/jmoiron/sqlx"
)

// MockServices is a mock of Services interface.
type MockServices struct {
	ctrl     *gomock.Controller
	recorder *MockServicesMockRecorder
}

// MockServicesMockRecorder is the mock recorder for MockServices.
type MockServicesMockRecorder struct {
	mock *MockServices
}

// NewMockServices creates a new mock instance.
func NewMockServices(ctrl *gomock.Controller) *MockServices {
	mock := &MockServices{ctrl: ctrl}
	mock.recorder = &MockServicesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServices) EXPECT() *MockServicesMockRecorder {
	return m.recorder
}

// BookCabTxn mocks base method.
func (m *MockServices) BookCabTxn(ctx context.Context, cityId int) (*models.Cab, *models.Ride, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BookCabTxn", ctx, cityId)
	ret0, _ := ret[0].(*models.Cab)
	ret1, _ := ret[1].(*models.Ride)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// BookCabTxn indicates an expected call of BookCabTxn.
func (mr *MockServicesMockRecorder) BookCabTxn(ctx, cityId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BookCabTxn", reflect.TypeOf((*MockServices)(nil).BookCabTxn), ctx, cityId)
}

// CreateCab mocks base method.
func (m *MockServices) CreateCab(ctx context.Context, cab *models.Cab) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCab", ctx, cab)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCab indicates an expected call of CreateCab.
func (mr *MockServicesMockRecorder) CreateCab(ctx, cab interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCab", reflect.TypeOf((*MockServices)(nil).CreateCab), ctx, cab)
}

// CreateCity mocks base method.
func (m *MockServices) CreateCity(ctx context.Context, city *models.City) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCity", ctx, city)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCity indicates an expected call of CreateCity.
func (mr *MockServicesMockRecorder) CreateCity(ctx, city interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCity", reflect.TypeOf((*MockServices)(nil).CreateCity), ctx, city)
}

// CreateRideRequest mocks base method.
func (m *MockServices) CreateRideRequest(ctx context.Context, rideRequest *models.RideRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRideRequest", ctx, rideRequest)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRideRequest indicates an expected call of CreateRideRequest.
func (mr *MockServicesMockRecorder) CreateRideRequest(ctx, rideRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRideRequest", reflect.TypeOf((*MockServices)(nil).CreateRideRequest), ctx, rideRequest)
}

// FinishRide mocks base method.
func (m *MockServices) FinishRide(ctx context.Context, rideId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinishRide", ctx, rideId)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinishRide indicates an expected call of FinishRide.
func (mr *MockServicesMockRecorder) FinishRide(ctx, rideId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinishRide", reflect.TypeOf((*MockServices)(nil).FinishRide), ctx, rideId)
}

// GetAllCabs mocks base method.
func (m *MockServices) GetAllCabs(ctx context.Context) ([]models.Cab, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCabs", ctx)
	ret0, _ := ret[0].([]models.Cab)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCabs indicates an expected call of GetAllCabs.
func (mr *MockServicesMockRecorder) GetAllCabs(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCabs", reflect.TypeOf((*MockServices)(nil).GetAllCabs), ctx)
}

// GetAllCities mocks base method.
func (m *MockServices) GetAllCities(ctx context.Context) ([]models.City, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCities", ctx)
	ret0, _ := ret[0].([]models.City)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCities indicates an expected call of GetAllCities.
func (mr *MockServicesMockRecorder) GetAllCities(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCities", reflect.TypeOf((*MockServices)(nil).GetAllCities), ctx)
}

// GetAllRides mocks base method.
func (m *MockServices) GetAllRides(ctx context.Context) ([]models.Ride, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllRides", ctx)
	ret0, _ := ret[0].([]models.Ride)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllRides indicates an expected call of GetAllRides.
func (mr *MockServicesMockRecorder) GetAllRides(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllRides", reflect.TypeOf((*MockServices)(nil).GetAllRides), ctx)
}

// GetCab mocks base method.
func (m *MockServices) GetCab(ctx context.Context, id int) (*models.Cab, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCab", ctx, id)
	ret0, _ := ret[0].(*models.Cab)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCab indicates an expected call of GetCab.
func (mr *MockServicesMockRecorder) GetCab(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCab", reflect.TypeOf((*MockServices)(nil).GetCab), ctx, id)
}

// GetCabActivities mocks base method.
func (m *MockServices) GetCabActivities(ctx context.Context, id int) ([]models.CabAudit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCabActivities", ctx, id)
	ret0, _ := ret[0].([]models.CabAudit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCabActivities indicates an expected call of GetCabActivities.
func (mr *MockServicesMockRecorder) GetCabActivities(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCabActivities", reflect.TypeOf((*MockServices)(nil).GetCabActivities), ctx, id)
}

// GetCabForUpdate mocks base method.
func (m *MockServices) GetCabForUpdate(ctx context.Context, id int, tx *sqlx.Tx) (*models.Cab, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCabForUpdate", ctx, id, tx)
	ret0, _ := ret[0].(*models.Cab)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCabForUpdate indicates an expected call of GetCabForUpdate.
func (mr *MockServicesMockRecorder) GetCabForUpdate(ctx, id, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCabForUpdate", reflect.TypeOf((*MockServices)(nil).GetCabForUpdate), ctx, id, tx)
}

// GetCity mocks base method.
func (m *MockServices) GetCity(ctx context.Context, id int) (*models.City, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCity", ctx, id)
	ret0, _ := ret[0].(*models.City)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCity indicates an expected call of GetCity.
func (mr *MockServicesMockRecorder) GetCity(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCity", reflect.TypeOf((*MockServices)(nil).GetCity), ctx, id)
}

// GetCityWiseRideInsight mocks base method.
func (m *MockServices) GetCityWiseRideInsight(ctx context.Context) ([]servies.RideInsight, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCityWiseRideInsight", ctx)
	ret0, _ := ret[0].([]servies.RideInsight)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCityWiseRideInsight indicates an expected call of GetCityWiseRideInsight.
func (mr *MockServicesMockRecorder) GetCityWiseRideInsight(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCityWiseRideInsight", reflect.TypeOf((*MockServices)(nil).GetCityWiseRideInsight), ctx)
}

// GetMostIdleCabOfCity mocks base method.
func (m *MockServices) GetMostIdleCabOfCity(ctx context.Context, cityId int, tx *sqlx.Tx) ([]models.Cab, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMostIdleCabOfCity", ctx, cityId, tx)
	ret0, _ := ret[0].([]models.Cab)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMostIdleCabOfCity indicates an expected call of GetMostIdleCabOfCity.
func (mr *MockServicesMockRecorder) GetMostIdleCabOfCity(ctx, cityId, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMostIdleCabOfCity", reflect.TypeOf((*MockServices)(nil).GetMostIdleCabOfCity), ctx, cityId, tx)
}

// UpdateCab mocks base method.
func (m *MockServices) UpdateCab(ctx context.Context, cab *models.Cab, tx *sqlx.Tx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCab", ctx, cab, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCab indicates an expected call of UpdateCab.
func (mr *MockServicesMockRecorder) UpdateCab(ctx, cab, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCab", reflect.TypeOf((*MockServices)(nil).UpdateCab), ctx, cab, tx)
}

// UpdateCabCity mocks base method.
func (m *MockServices) UpdateCabCity(ctx context.Context, cab *models.Cab, tx *sqlx.Tx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCabCity", ctx, cab, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCabCity indicates an expected call of UpdateCabCity.
func (mr *MockServicesMockRecorder) UpdateCabCity(ctx, cab, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCabCity", reflect.TypeOf((*MockServices)(nil).UpdateCabCity), ctx, cab, tx)
}

// UpdateCabCityTxn mocks base method.
func (m *MockServices) UpdateCabCityTxn(ctx context.Context, CabId, CurrentCityId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCabCityTxn", ctx, CabId, CurrentCityId)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCabCityTxn indicates an expected call of UpdateCabCityTxn.
func (mr *MockServicesMockRecorder) UpdateCabCityTxn(ctx, CabId, CurrentCityId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCabCityTxn", reflect.TypeOf((*MockServices)(nil).UpdateCabCityTxn), ctx, CabId, CurrentCityId)
}

// UpdateCabState mocks base method.
func (m *MockServices) UpdateCabState(ctx context.Context, cab *models.Cab, tx *sqlx.Tx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCabState", ctx, cab, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCabState indicates an expected call of UpdateCabState.
func (mr *MockServicesMockRecorder) UpdateCabState(ctx, cab, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCabState", reflect.TypeOf((*MockServices)(nil).UpdateCabState), ctx, cab, tx)
}

// UpdateRide mocks base method.
func (m *MockServices) UpdateRide(ctx context.Context, ride *models.Ride, tx *sqlx.Tx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRide", ctx, ride, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRide indicates an expected call of UpdateRide.
func (mr *MockServicesMockRecorder) UpdateRide(ctx, ride, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRide", reflect.TypeOf((*MockServices)(nil).UpdateRide), ctx, ride, tx)
}
