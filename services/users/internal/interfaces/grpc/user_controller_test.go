package grpc

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/dzhordano/132market/services/users/internal/application/command"
	"github.com/dzhordano/132market/services/users/internal/application/mapper"
	"github.com/dzhordano/132market/services/users/internal/application/model"
	"github.com/dzhordano/132market/services/users/internal/application/query"
	"github.com/dzhordano/132market/services/users/internal/domain/entities"
	mock_interfaces "github.com/dzhordano/132market/services/users/internal/interfaces/grpc/mocks"
	"github.com/dzhordano/132market/services/users/pkg/pb/user_v1"
	"github.com/go-playground/assert/v2"
	"github.com/google/uuid"
	"go.uber.org/mock/gomock"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Test_FindUserById(t *testing.T) {
	type mockBehavior func(s *mock_interfaces.MockUserService, id string)

	testUser_1 := &entities.User{
		ID:         uuid.UUID{},
		Name:       "test",
		Email:      "test@mail.ru",
		Roles:      []entities.Role{entities.RoleUser},
		Status:     entities.StatusOffline,
		State:      entities.StateActive,
		CreatedAt:  time.Now(),
		LastSeenAt: time.Now(),
		DeletedAt:  time.Time{},
	}

	tests := []struct {
		name           string
		inpReq         *user_v1.FindUserByIdRequest
		mockBehavior   mockBehavior
		expectedResult *user_v1.FindUserByIdResponse
		expectedErr    error
	}{
		{
			name: "OK",
			inpReq: &user_v1.FindUserByIdRequest{
				Id: uuid.UUID{}.String(),
			},
			mockBehavior: func(s *mock_interfaces.MockUserService, id string) {
				s.EXPECT().FindUserById(gomock.Any(), id).Return(&query.UserQueryResult{
					Result: mapper.NewUserResultFromEntity(testUser_1),
				}, nil)
			},
			expectedResult: &user_v1.FindUserByIdResponse{
				User: &user_v1.User{
					Id:         uuid.UUID{}.String(),
					Name:       "test",
					Email:      "test@mail.ru",
					Roles:      []string{entities.RoleUser.String()},
					Status:     "offline",
					State:      "active",
					CreatedAt:  timestamppb.New(time.Now()),
					LastSeenAt: timestamppb.New(time.Now()),
				},
			},
			expectedErr: nil,
		},
		{
			name: "Not found",
			inpReq: &user_v1.FindUserByIdRequest{
				Id: uuid.UUID{}.String(),
			},
			mockBehavior: func(s *mock_interfaces.MockUserService, id string) {
				s.EXPECT().FindUserById(gomock.Any(), id).Return(nil, errors.New("not found"))
			},
			expectedResult: nil,
			expectedErr:    errors.New("not found"),
		},
		{
			name: "Internal Failure",
			inpReq: &user_v1.FindUserByIdRequest{
				Id: uuid.UUID{}.String(),
			},
			mockBehavior: func(s *mock_interfaces.MockUserService, id string) {
				s.EXPECT().FindUserById(gomock.Any(), id).Return(nil, errors.New("internal failure"))
			},
			expectedResult: nil,
			expectedErr:    errors.New("internal failure"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			s := mock_interfaces.NewMockUserService(c)
			test.mockBehavior(s, test.inpReq.Id)

			ctrl := NewUserController(s)
			res, err := ctrl.FindUserById(context.Background(), test.inpReq)

			if test.expectedErr != nil && test.expectedErr.Error() != err.Error() {
				t.Errorf("expected error %v got %v", test.expectedErr, err)
			}

			assert.Equal(t, test.expectedResult, res)
		})
	}

}
func Test_FindUserByCredentials(t *testing.T) {
	type mockBehavior func(s *mock_interfaces.MockUserService, email string)

	testUser_1 := &entities.User{
		ID:         uuid.UUID{},
		Name:       "test",
		Email:      "test@mail.ru",
		Roles:      []entities.Role{entities.RoleUser},
		Status:     entities.StatusOffline,
		State:      entities.StateActive,
		CreatedAt:  time.Now(),
		LastSeenAt: time.Now(),
		DeletedAt:  time.Time{},
	}

	tests := []struct {
		name           string
		inpReq         *user_v1.FindUserByEmailRequest
		mockBehavior   mockBehavior
		expectedResult *user_v1.FindUserByEmailResponse
		expectedErr    error
	}{
		{
			name: "OK",
			inpReq: &user_v1.FindUserByEmailRequest{
				Email: "test@mail.ru",
			},
			mockBehavior: func(s *mock_interfaces.MockUserService, email string) {
				s.EXPECT().FindUserByEmail(context.Background(), email).Return(&query.UserQueryResult{
					Result: mapper.NewUserResultFromEntity(testUser_1),
				}, nil)
			},
			expectedResult: &user_v1.FindUserByEmailResponse{
				User: &user_v1.User{
					Id:         uuid.UUID{}.String(),
					Name:       "test",
					Email:      "test@mail.ru",
					Roles:      []string{entities.RoleUser.String()},
					Status:     "offline",
					State:      "active",
					CreatedAt:  timestamppb.New(time.Now()),
					LastSeenAt: timestamppb.New(time.Now()),
				},
			},
			expectedErr: nil,
		},
		{
			name: "Not found",
			inpReq: &user_v1.FindUserByEmailRequest{
				Email: "test@mail.ru",
			},
			mockBehavior: func(s *mock_interfaces.MockUserService, email string) {
				s.EXPECT().FindUserByEmail(context.Background(), email).Return(nil, errors.New("not found"))
			},
			expectedResult: nil,
			expectedErr:    errors.New("not found"),
		},
		{
			name: "Internal Failure",
			inpReq: &user_v1.FindUserByEmailRequest{
				Email: "test@mail.ru",
			},
			mockBehavior: func(s *mock_interfaces.MockUserService, email string) {
				s.EXPECT().FindUserByEmail(context.Background(), email).Return(nil, errors.New("internal failure"))
			},
			expectedResult: nil,
			expectedErr:    errors.New("internal failure"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			s := mock_interfaces.NewMockUserService(c)
			test.mockBehavior(s, test.inpReq.Email)

			ctrl := NewUserController(s)
			res, err := ctrl.FindUserByEmail(context.Background(), test.inpReq)

			if test.expectedErr != nil && test.expectedErr.Error() != err.Error() {
				t.Errorf("expected error %v got %v", test.expectedErr, err)
			}

			assert.Equal(t, test.expectedResult, res)
		})
	}
}
func Test_FindAllUsers(t *testing.T) {
	type mockBehavior func(s *mock_interfaces.MockUserService, offset, limit uint64, filters map[string]string)

	testUser_1 := &entities.User{
		ID:         uuid.UUID{},
		Name:       "test1",
		Email:      "test1@mail.ru",
		Roles:      []entities.Role{entities.RoleUser},
		Status:     entities.StatusOffline,
		State:      entities.StateActive,
		CreatedAt:  time.Now(),
		LastSeenAt: time.Now(),
		DeletedAt:  time.Time{},
	}
	testUser_2 := &entities.User{
		ID:         uuid.UUID{},
		Name:       "test2",
		Email:      "test2@mail.ru",
		Roles:      []entities.Role{entities.RoleUser},
		Status:     entities.StatusOffline,
		State:      entities.StateActive,
		CreatedAt:  time.Now(),
		LastSeenAt: time.Now(),
		DeletedAt:  time.Time{},
	}
	testUser_3 := &entities.User{
		ID:         uuid.UUID{},
		Name:       "test3",
		Email:      "test3@mail.ru",
		Roles:      []entities.Role{entities.RoleUser},
		Status:     entities.StatusOffline,
		State:      entities.StateActive,
		CreatedAt:  time.Now(),
		LastSeenAt: time.Now(),
		DeletedAt:  time.Time{},
	}

	tests := []struct {
		name           string
		inpReq         *user_v1.ListUsersRequest
		mockBehavior   mockBehavior
		expectedResult *user_v1.ListUsersResponse
		expectedErr    error
	}{
		{
			name: "OK", inpReq: &user_v1.ListUsersRequest{
				Offset: 0,
				Limit:  10,
			}, mockBehavior: func(s *mock_interfaces.MockUserService, offset, limit uint64, filters map[string]string) {
				s.EXPECT().ListUsers(context.Background(), offset, limit, filters).Return(&query.UserQueryListResult{
					Result: mapper.NewUserResultListFromEntities([]*entities.User{
						testUser_1, testUser_2, testUser_3})},
					nil)
			}, expectedResult: &user_v1.ListUsersResponse{
				Users: []*user_v1.User{
					{
						Id:         uuid.UUID{}.String(),
						Name:       "test1",
						Email:      "test1@mail.ru",
						Roles:      []string{entities.RoleUser.String()},
						Status:     "offline",
						State:      "active",
						CreatedAt:  timestamppb.New(time.Now()),
						LastSeenAt: timestamppb.New(time.Now()),
					},
					{
						Id:         uuid.UUID{}.String(),
						Name:       "test2",
						Email:      "test2@mail.ru",
						Roles:      []string{entities.RoleUser.String()},
						Status:     "offline",
						State:      "active",
						CreatedAt:  timestamppb.New(time.Now()),
						LastSeenAt: timestamppb.New(time.Now()),
					},
					{
						Id:         uuid.UUID{}.String(),
						Name:       "test3",
						Email:      "test3@mail.ru",
						Roles:      []string{entities.RoleUser.String()},
						Status:     "offline",
						State:      "active",
						CreatedAt:  timestamppb.New(time.Now()),
						LastSeenAt: timestamppb.New(time.Now()),
					},
				},
			}, expectedErr: nil},
		{
			name: "Retrieve with limit 2",
			inpReq: &user_v1.ListUsersRequest{
				Offset: 0,
				Limit:  2,
			},
			mockBehavior: func(s *mock_interfaces.MockUserService, offset, limit uint64, filters map[string]string) {
				s.EXPECT().ListUsers(context.Background(), offset, limit, filters).Return(&query.UserQueryListResult{
					Result: mapper.NewUserResultListFromEntities([]*entities.User{
						testUser_1, testUser_2})},
					nil)
			},
			expectedResult: &user_v1.ListUsersResponse{
				Users: []*user_v1.User{
					{
						Id:         uuid.UUID{}.String(),
						Name:       "test1",
						Email:      "test1@mail.ru",
						Roles:      []string{entities.RoleUser.String()},
						Status:     "offline",
						State:      "active",
						CreatedAt:  timestamppb.New(time.Now()),
						LastSeenAt: timestamppb.New(time.Now()),
					},
					{
						Id:         uuid.UUID{}.String(),
						Name:       "test2",
						Email:      "test2@mail.ru",
						Roles:      []string{entities.RoleUser.String()},
						Status:     "offline",
						State:      "active",
						CreatedAt:  timestamppb.New(time.Now()),
						LastSeenAt: timestamppb.New(time.Now()),
					},
				},
			},
			expectedErr: nil,
		},
		{
			name: "Retrieve with offset 2",
			inpReq: &user_v1.ListUsersRequest{
				Offset: 2,
				Limit:  10,
			},
			mockBehavior: func(s *mock_interfaces.MockUserService, offset, limit uint64, filters map[string]string) {
				s.EXPECT().ListUsers(context.Background(), offset, limit, filters).Return(&query.UserQueryListResult{
					Result: mapper.NewUserResultListFromEntities([]*entities.User{
						testUser_3})},
					nil)
			},
			expectedResult: &user_v1.ListUsersResponse{Users: []*user_v1.User{
				{
					Id:         uuid.UUID{}.String(),
					Name:       "test3",
					Email:      "test3@mail.ru",
					Roles:      []string{entities.RoleUser.String()},
					Status:     "offline",
					State:      "active",
					CreatedAt:  timestamppb.New(time.Now()),
					LastSeenAt: timestamppb.New(time.Now()),
				},
			},
			},
			expectedErr: nil,
		},
		{
			name: "Retrieve with limit 2 and offset 2",
			inpReq: &user_v1.ListUsersRequest{
				Offset: 2,
				Limit:  2,
			},
			mockBehavior: func(s *mock_interfaces.MockUserService, offset, limit uint64, filters map[string]string) {
				s.EXPECT().ListUsers(context.Background(), offset, limit, filters).Return(&query.UserQueryListResult{
					Result: mapper.NewUserResultListFromEntities([]*entities.User{
						testUser_3})},
					nil)
			},
			expectedResult: &user_v1.ListUsersResponse{
				Users: []*user_v1.User{
					{
						Id:         uuid.UUID{}.String(),
						Name:       "test3",
						Email:      "test3@mail.ru",
						Roles:      []string{entities.RoleUser.String()},
						Status:     "offline",
						State:      "active",
						CreatedAt:  timestamppb.New(time.Now()),
						LastSeenAt: timestamppb.New(time.Now()),
					},
				},
			},
			expectedErr: nil},
		{
			name: "Internal failure",
			inpReq: &user_v1.ListUsersRequest{
				Offset: 2,
				Limit:  2,
			},
			mockBehavior: func(s *mock_interfaces.MockUserService, offset, limit uint64, filters map[string]string) {
				s.EXPECT().ListUsers(context.Background(), offset, limit, filters).Return(nil, errors.New("internal failure"))
			},
			expectedResult: nil,
			expectedErr:    errors.New("internal failure"),
		},
		{
			name: "Not found",
			inpReq: &user_v1.ListUsersRequest{
				Offset: 3,
				Limit:  10,
			},
			mockBehavior: func(s *mock_interfaces.MockUserService, offset, limit uint64, filters map[string]string) {
				s.EXPECT().ListUsers(context.Background(), offset, limit, filters).Return(nil, errors.New("not found"))
			},
			expectedErr: errors.New("not found"),
		},
		{
			// FIXME ADD TEST TO TEST FILTERS
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			s := mock_interfaces.NewMockUserService(c)
			test.mockBehavior(s, test.inpReq.Offset, test.inpReq.Limit, test.inpReq.Filters)

			ctrl := NewUserController(s)
			res, err := ctrl.ListUsers(context.Background(), test.inpReq)

			if test.expectedErr != nil && test.expectedErr.Error() != err.Error() {
				t.Errorf("expected error %v got %v", test.expectedErr, err)
			}

			assert.Equal(t, test.expectedResult, res)
		})
	}
}

// FIXME Когда будут хеши, пароли на хеши.
func Test_CreateUser(t *testing.T) {
	type mockBehavior func(s *mock_interfaces.MockUserService, cmd *command.CreateUserCommand)

	testUser_1 := &user_v1.User{
		Id:         uuid.UUID{}.String(),
		Name:       "test",
		Email:      "test@mail.ru",
		Roles:      []string{entities.RoleUser.String()},
		Status:     "offline",
		State:      "active",
		CreatedAt:  timestamppb.New(time.Now()),
		LastSeenAt: timestamppb.New(time.Now()),
	}

	tests := []struct {
		name           string
		inpReq         *user_v1.CreateUserRequest
		inpCmd         *command.CreateUserCommand
		mockBehavior   mockBehavior
		expectedResult *user_v1.CreateUserResponse
		expectedErr    error
	}{
		{
			name: "OK",
			inpReq: &user_v1.CreateUserRequest{
				Info: &user_v1.UserInfo{
					Name:  testUser_1.Name,
					Email: testUser_1.Email,
				},
			},
			inpCmd: &command.CreateUserCommand{
				Name:  testUser_1.Name,
				Email: testUser_1.Email,
			},
			mockBehavior: func(s *mock_interfaces.MockUserService, cmd *command.CreateUserCommand) {
				s.EXPECT().CreateUser(context.Background(), cmd).Return(&command.CreateUserCommandResult{
					Result: &model.UserResult{
						ID:         uuid.UUID{},
						Name:       testUser_1.Name,
						Email:      testUser_1.Email,
						Roles:      testUser_1.Roles,
						Status:     testUser_1.Status,
						State:      testUser_1.State,
						CreatedAt:  testUser_1.CreatedAt.AsTime(),
						LastSeenAt: testUser_1.LastSeenAt.AsTime(),
					},
				}, nil)
			},
			expectedResult: &user_v1.CreateUserResponse{
				User: testUser_1,
			},
			expectedErr: nil,
		},
		{
			name: "Internal Failure",
			inpReq: &user_v1.CreateUserRequest{
				Info: &user_v1.UserInfo{
					Name:  testUser_1.Name,
					Email: testUser_1.Email,
				},
			},
			inpCmd: &command.CreateUserCommand{
				Name:  testUser_1.Name,
				Email: testUser_1.Email,
			},
			mockBehavior: func(s *mock_interfaces.MockUserService, cmd *command.CreateUserCommand) {
				s.EXPECT().CreateUser(context.Background(), cmd).Return(nil, errors.New("internal failure"))
			},
			expectedResult: nil,
			expectedErr:    errors.New("internal failure"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			s := mock_interfaces.NewMockUserService(c)
			test.mockBehavior(s, test.inpCmd)

			ctrl := NewUserController(s)

			resp, err := ctrl.CreateUser(context.Background(), test.inpReq)

			if test.expectedResult != nil {
				assert.Equal(t, test.expectedResult, resp)
			}

			assert.Equal(t, test.expectedErr, err)
		})
	}
}

func Test_UpdateUser(t *testing.T) {
	type mockBehavior func(s *mock_interfaces.MockUserService, cmd *command.UpdateUserCommand)

	testUser_1 := &user_v1.User{
		Id:         uuid.UUID{}.String(),
		Name:       "test",
		Email:      "test@mail.ru",
		Roles:      []string{entities.RoleUser.String()},
		Status:     "offline",
		State:      "active",
		CreatedAt:  timestamppb.New(time.Now()),
		LastSeenAt: timestamppb.New(time.Now()),
	}

	tests := []struct {
		name           string
		inpReq         *user_v1.UpdateUserRequest
		inpCmd         *command.UpdateUserCommand
		mockBehavior   mockBehavior
		expectedResult *user_v1.UpdateUserResponse
		expectedErr    error
	}{
		{
			name: "OK",
			inpReq: &user_v1.UpdateUserRequest{
				Info: &user_v1.UserInfo{
					Name:  testUser_1.Name,
					Email: "new@mail.ru",
				},
			},
			inpCmd: &command.UpdateUserCommand{
				Name:  testUser_1.Name,
				Email: "new@mail.ru",
			},
			mockBehavior: func(s *mock_interfaces.MockUserService, cmd *command.UpdateUserCommand) {
				s.EXPECT().UpdateUser(context.Background(), cmd).Return(&command.UpdateUserCommandResult{
					Result: &model.UserResult{
						ID:         uuid.UUID{},
						Name:       testUser_1.Name,
						Email:      "new@mail.ru",
						Roles:      testUser_1.Roles,
						Status:     testUser_1.Status,
						State:      testUser_1.State,
						CreatedAt:  testUser_1.CreatedAt.AsTime(),
						LastSeenAt: testUser_1.LastSeenAt.AsTime(),
					},
				}, nil)
			},
			expectedResult: &user_v1.UpdateUserResponse{
				User: &user_v1.User{
					Id:         uuid.UUID{}.String(),
					Name:       testUser_1.Name,
					Email:      "new@mail.ru",
					Roles:      testUser_1.Roles,
					Status:     testUser_1.Status,
					State:      testUser_1.State,
					CreatedAt:  timestamppb.New(testUser_1.CreatedAt.AsTime()),
					LastSeenAt: timestamppb.New(testUser_1.LastSeenAt.AsTime()),
				},
			},
			expectedErr: nil,
		},
		{
			name: "User Not Found",
			inpReq: &user_v1.UpdateUserRequest{
				Info: &user_v1.UserInfo{
					Name:  "test",
					Email: "new@mail.ru",
				},
			},
			inpCmd: &command.UpdateUserCommand{
				Name:  "test",
				Email: "new@mail.ru",
			},
			mockBehavior: func(s *mock_interfaces.MockUserService, cmd *command.UpdateUserCommand) {
				s.EXPECT().UpdateUser(context.Background(), cmd).Return(nil, errors.New("not found"))
			},
			expectedResult: nil,
			expectedErr:    errors.New("not found"),
		},
		{
			name: "Internal Failure",
			inpReq: &user_v1.UpdateUserRequest{
				Info: &user_v1.UserInfo{
					Name:  testUser_1.Name,
					Email: "new@mail.ru",
				},
			},
			inpCmd: &command.UpdateUserCommand{
				Name:  testUser_1.Name,
				Email: "new@mail.ru",
			},
			mockBehavior: func(s *mock_interfaces.MockUserService, cmd *command.UpdateUserCommand) {
				s.EXPECT().UpdateUser(context.Background(), cmd).Return(nil, errors.New("internal failure"))
			},
			expectedResult: nil,
			expectedErr:    errors.New("internal failure"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			s := mock_interfaces.NewMockUserService(c)
			test.mockBehavior(s, test.inpCmd)

			ctrl := NewUserController(s)

			resp, err := ctrl.UpdateUser(context.Background(), test.inpReq)

			if test.expectedResult != nil {
				assert.Equal(t, test.expectedResult, resp)
			}

			assert.Equal(t, test.expectedErr, err)
		})
	}
}

func Test_DeleteUser(t *testing.T) {
	type mockBehavior func(s *mock_interfaces.MockUserService, id string)

	tests := []struct {
		name         string
		inpReq       *user_v1.DeleteUserRequest
		mockBehavior mockBehavior
		expectedErr  error
	}{
		{
			name: "OK",
			inpReq: &user_v1.DeleteUserRequest{
				Id: uuid.NewString(),
			},
			mockBehavior: func(s *mock_interfaces.MockUserService, id string) {
				s.EXPECT().DeleteUser(gomock.Any(), id).Return(nil).AnyTimes()
			},
			expectedErr: nil,
		},
		{
			name: "User Not Found",
			inpReq: &user_v1.DeleteUserRequest{
				Id: uuid.NewString(),
			},
			mockBehavior: func(s *mock_interfaces.MockUserService, id string) {
				s.EXPECT().DeleteUser(gomock.Any(), id).Return(errors.New("not found")).AnyTimes()
			},
			expectedErr: errors.New("not found"),
		},
		{
			name: "Internal Failure",
			inpReq: &user_v1.DeleteUserRequest{
				Id: uuid.NewString(),
			},
			mockBehavior: func(s *mock_interfaces.MockUserService, id string) {
				s.EXPECT().DeleteUser(gomock.Any(), id).Return(errors.New("internal failure")).AnyTimes()
			},
			expectedErr: errors.New("internal failure"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			s := mock_interfaces.NewMockUserService(c)
			test.mockBehavior(s, test.inpReq.Id)

			ctrl := NewUserController(s)

			resp, err := ctrl.DeleteUser(context.Background(), test.inpReq)

			if resp != nil {
				assert.Equal(t, resp, &emptypb.Empty{})
			}

			if test.expectedErr != nil {
				assert.Equal(t, test.expectedErr.Error(), err.Error())
			}
		})
	}
}
