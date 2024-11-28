package mapper

import (
	"dzhordano/132market/services/users/internal/application/command"
	"dzhordano/132market/services/users/pkg/pb/user_v1"
)

func CreateUserRequestToCommand(req *user_v1.CreateUserRequest) *command.CreateUserCommand {
	return &command.CreateUserCommand{
		Name:         req.Info.GetName(),
		Email:        req.Info.GetEmail(),
		PasswordHash: req.Info.GetPassword(),
	}
}

func UpdateUserRequestToCommand(req *user_v1.UpdateUserRequest) *command.UpdateUserCommand {
	return &command.UpdateUserCommand{
		ID:           req.GetId(),
		Name:         req.Info.GetName(),
		Email:        req.Info.GetEmail(),
		PasswordHash: req.Info.GetPassword(),
	}
}
