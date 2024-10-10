package users

import (
	"dzhordano/132market/services/users/internal/service"
	"dzhordano/132market/services/users/pkg/users_v1"
)

type Implementation struct {
	users_v1.UnimplementedUsersV1Server
	usersService service.UsersService
}

func New(usersService service.UsersService) *Implementation {
	return &Implementation{
		usersService: usersService,
	}
}
