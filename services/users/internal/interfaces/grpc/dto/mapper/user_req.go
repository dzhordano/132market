package mapper

import (
	"github.com/dzhordano/132market/services/users/internal/application/command"
	"github.com/dzhordano/132market/services/users/pkg/pb/user_v1"
)

func CreateUserRequestToCommand(req *user_v1.CreateUserRequest) *command.CreateUserCommand {
	return &command.CreateUserCommand{
		Name:  req.Info.GetName(),
		Email: req.Info.GetEmail(),
	}
}

func UpdateUserRequestToCommand(req *user_v1.UpdateUserRequest) *command.UpdateUserCommand {
	return &command.UpdateUserCommand{
		ID:    req.GetId(),
		Name:  req.Info.GetName(),
		Email: req.Info.GetEmail(),
	}
}
