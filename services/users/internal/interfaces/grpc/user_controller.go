package grpc

import (
	"context"

	"github.com/dzhordano/132market/services/users/internal/application/interfaces"
	"github.com/dzhordano/132market/services/users/internal/interfaces/grpc/dto/mapper"
	"github.com/dzhordano/132market/services/users/pkg/pb/user_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserController struct {
	userService interfaces.UserService
	user_v1.UnimplementedUserServiceV1Server
}

func NewUserController(userService interfaces.UserService) *UserController {
	return &UserController{userService: userService}
}

func (u *UserController) CreateUser(ctx context.Context, request *user_v1.CreateUserRequest) (*user_v1.CreateUserResponse, error) {
	userCommand := mapper.CreateUserRequestToCommand(request)

	commandResp, err := u.userService.CreateUser(ctx, userCommand)
	if err != nil {
		return nil, err
	}

	response := mapper.ToUserResponse(commandResp.Result)

	return &user_v1.CreateUserResponse{User: response}, nil
}

func (u *UserController) UpdateUser(ctx context.Context, request *user_v1.UpdateUserRequest) (*user_v1.UpdateUserResponse, error) {
	userCommand := mapper.UpdateUserRequestToCommand(request)

	commandResp, err := u.userService.UpdateUser(ctx, userCommand)
	if err != nil {
		return nil, err
	}

	response := mapper.ToUserResponse(commandResp.Result)

	return &user_v1.UpdateUserResponse{User: response}, nil
}

func (u *UserController) DeleteUser(ctx context.Context, request *user_v1.DeleteUserRequest) (*emptypb.Empty, error) {
	err := u.userService.DeleteUser(ctx, request.GetId())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (u *UserController) FindUserById(ctx context.Context, request *user_v1.FindUserByIdRequest) (*user_v1.FindUserByIdResponse, error) {
	queryResp, err := u.userService.FindUserById(ctx, request.GetId())
	if err != nil {
		return nil, err
	}

	response := mapper.ToUserResponse(queryResp.Result)

	return &user_v1.FindUserByIdResponse{User: response}, nil
}

func (u *UserController) FindUserByEmail(ctx context.Context, request *user_v1.FindUserByEmailRequest) (*user_v1.FindUserByEmailResponse, error) {
	queryResp, err := u.userService.FindUserByEmail(ctx, request.GetEmail())
	if err != nil {
		return nil, err
	}

	response := mapper.ToUserResponse(queryResp.Result)

	return &user_v1.FindUserByEmailResponse{User: response}, nil
}

func (u *UserController) ListUsers(ctx context.Context, request *user_v1.ListUsersRequest) (*user_v1.ListUsersResponse, error) {
	queryResp, err := u.userService.ListUsers(ctx, request.GetOffset(), request.GetLimit(), request.GetFilters())
	if err != nil {
		return nil, err
	}

	response := mapper.ToUserListResponse(queryResp.Result)

	return &user_v1.ListUsersResponse{Users: response}, nil
}

// FIXME последним сделать.
func (u *UserController) SearchUsers(ctx context.Context, request *user_v1.SearchUsersRequest) (*user_v1.SearchUsersResponse, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented. waiting for searching service implementation")
}

func (u *UserController) SetUserState(ctx context.Context, request *user_v1.SetUserStateRequest) (*emptypb.Empty, error) {
	err := u.userService.SetUserState(ctx, request.GetId(), request.GetState())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (u *UserController) UpdateLastSeen(ctx context.Context, request *user_v1.UpdateLastSeenRequest) (*emptypb.Empty, error) {
	err := u.userService.UpdateLastSeen(ctx, request.GetId())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (u *UserController) CheckUserExists(ctx context.Context, request *user_v1.CheckUserExistsRequest) (*user_v1.CheckUserExistsResponse, error) {
	exists, err := u.userService.CheckUserExists(ctx, request.GetEmail())
	if err != nil {
		return nil, err
	}

	return &user_v1.CheckUserExistsResponse{Exists: exists}, nil
}
