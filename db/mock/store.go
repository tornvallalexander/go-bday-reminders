// Code generated by MockGen. DO NOT EDIT.
// Source: go-bday-reminders/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	db "go-bday-reminders/db/sqlc"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateReminder mocks base method.
func (m *MockStore) CreateReminder(arg0 context.Context, arg1 db.CreateReminderParams) (db.Reminder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateReminder", arg0, arg1)
	ret0, _ := ret[0].(db.Reminder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateReminder indicates an expected call of CreateReminder.
func (mr *MockStoreMockRecorder) CreateReminder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateReminder", reflect.TypeOf((*MockStore)(nil).CreateReminder), arg0, arg1)
}

// CreateSession mocks base method.
func (m *MockStore) CreateSession(arg0 context.Context, arg1 db.CreateSessionParams) (db.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", arg0, arg1)
	ret0, _ := ret[0].(db.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockStoreMockRecorder) CreateSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockStore)(nil).CreateSession), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// DeleteReminder mocks base method.
func (m *MockStore) DeleteReminder(arg0 context.Context, arg1 int64) (db.Reminder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteReminder", arg0, arg1)
	ret0, _ := ret[0].(db.Reminder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteReminder indicates an expected call of DeleteReminder.
func (mr *MockStoreMockRecorder) DeleteReminder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteReminder", reflect.TypeOf((*MockStore)(nil).DeleteReminder), arg0, arg1)
}

// DeleteUser mocks base method.
func (m *MockStore) DeleteUser(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockStoreMockRecorder) DeleteUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockStore)(nil).DeleteUser), arg0, arg1)
}

// GetReminder mocks base method.
func (m *MockStore) GetReminder(arg0 context.Context, arg1 int64) (db.Reminder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReminder", arg0, arg1)
	ret0, _ := ret[0].(db.Reminder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReminder indicates an expected call of GetReminder.
func (mr *MockStoreMockRecorder) GetReminder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReminder", reflect.TypeOf((*MockStore)(nil).GetReminder), arg0, arg1)
}

// GetSession mocks base method.
func (m *MockStore) GetSession(arg0 context.Context, arg1 uuid.UUID) (db.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSession", arg0, arg1)
	ret0, _ := ret[0].(db.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSession indicates an expected call of GetSession.
func (mr *MockStoreMockRecorder) GetSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSession", reflect.TypeOf((*MockStore)(nil).GetSession), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockStore) GetUser(arg0 context.Context, arg1 string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockStoreMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockStore)(nil).GetUser), arg0, arg1)
}

// ListReminders mocks base method.
func (m *MockStore) ListReminders(arg0 context.Context, arg1 db.ListRemindersParams) ([]db.Reminder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListReminders", arg0, arg1)
	ret0, _ := ret[0].([]db.Reminder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListReminders indicates an expected call of ListReminders.
func (mr *MockStoreMockRecorder) ListReminders(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListReminders", reflect.TypeOf((*MockStore)(nil).ListReminders), arg0, arg1)
}

// UpdateReminder mocks base method.
func (m *MockStore) UpdateReminder(arg0 context.Context, arg1 db.UpdateReminderParams) (db.Reminder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateReminder", arg0, arg1)
	ret0, _ := ret[0].(db.Reminder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateReminder indicates an expected call of UpdateReminder.
func (mr *MockStoreMockRecorder) UpdateReminder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateReminder", reflect.TypeOf((*MockStore)(nil).UpdateReminder), arg0, arg1)
}
