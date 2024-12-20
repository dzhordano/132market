// Code generated by MockGen. DO NOT EDIT.
// Source: internal/application/interfaces/user_service.go
//
// Generated by this command:
//
//	mockgen -source=internal/application/interfaces/user_service.go -destination=internal/interfaces/grpc/mocks/mocks.go
//

// Package mock_interfaces is a generated GoMock package.
package mock_interfaces

import (
	context "context"
	reflect "reflect"

	command "github.com/dzhordano/132market/services/users/internal/application/command"
	query "github.com/dzhordano/132market/services/users/internal/application/query"
	gomock "go.uber.org/mock/gomock"
)

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
	isgomock struct{}
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// CheckUserExists mocks base method.
func (m *MockUserService) CheckUserExists(ctx context.Context, email string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUserExists", ctx, email)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckUserExists indicates an expected call of CheckUserExists.
func (mr *MockUserServiceMockRecorder) CheckUserExists(ctx, email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserExists", reflect.TypeOf((*MockUserService)(nil).CheckUserExists), ctx, email)
}

// CreateUser mocks base method.
func (m *MockUserService) CreateUser(ctx context.Context, userCommand *command.CreateUserCommand) (*command.CreateUserCommandResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, userCommand)
	ret0, _ := ret[0].(*command.CreateUserCommandResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserServiceMockRecorder) CreateUser(ctx, userCommand any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserService)(nil).CreateUser), ctx, userCommand)
}

// DeleteUser mocks base method.
func (m *MockUserService) DeleteUser(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserServiceMockRecorder) DeleteUser(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserService)(nil).DeleteUser), ctx, id)
}

// FindUserByEmail mocks base method.
func (m *MockUserService) FindUserByEmail(ctx context.Context, email string) (*query.UserQueryResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByEmail", ctx, email)
	ret0, _ := ret[0].(*query.UserQueryResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByEmail indicates an expected call of FindUserByEmail.
func (mr *MockUserServiceMockRecorder) FindUserByEmail(ctx, email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByEmail", reflect.TypeOf((*MockUserService)(nil).FindUserByEmail), ctx, email)
}

// FindUserById mocks base method.
func (m *MockUserService) FindUserById(ctx context.Context, id string) (*query.UserQueryResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserById", ctx, id)
	ret0, _ := ret[0].(*query.UserQueryResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserById indicates an expected call of FindUserById.
func (mr *MockUserServiceMockRecorder) FindUserById(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserById", reflect.TypeOf((*MockUserService)(nil).FindUserById), ctx, id)
}

// ListUsers mocks base method.
func (m *MockUserService) ListUsers(ctx context.Context, offset, limit uint64, filters map[string]string) (*query.UserQueryListResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsers", ctx, offset, limit, filters)
	ret0, _ := ret[0].(*query.UserQueryListResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsers indicates an expected call of ListUsers.
func (mr *MockUserServiceMockRecorder) ListUsers(ctx, offset, limit, filters any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockUserService)(nil).ListUsers), ctx, offset, limit, filters)
}

// SetUserState mocks base method.
func (m *MockUserService) SetUserState(ctx context.Context, id, state string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUserState", ctx, id, state)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetUserState indicates an expected call of SetUserState.
func (mr *MockUserServiceMockRecorder) SetUserState(ctx, id, state any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUserState", reflect.TypeOf((*MockUserService)(nil).SetUserState), ctx, id, state)
}

// UpdateLastSeen mocks base method.
func (m *MockUserService) UpdateLastSeen(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLastSeen", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateLastSeen indicates an expected call of UpdateLastSeen.
func (mr *MockUserServiceMockRecorder) UpdateLastSeen(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLastSeen", reflect.TypeOf((*MockUserService)(nil).UpdateLastSeen), ctx, id)
}

// UpdateUser mocks base method.
func (m *MockUserService) UpdateUser(ctx context.Context, userCommand *command.UpdateUserCommand) (*command.UpdateUserCommandResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, userCommand)
	ret0, _ := ret[0].(*command.UpdateUserCommandResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserServiceMockRecorder) UpdateUser(ctx, userCommand any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserService)(nil).UpdateUser), ctx, userCommand)
}
