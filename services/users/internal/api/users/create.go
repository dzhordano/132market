package users

import (
	"context"
	"dzhordano/132market/services/users/internal/domain"
	"dzhordano/132market/services/users/pkg/users_v1"
	"fmt"
	"time"
)

func (i *Implementation) Create(ctx context.Context, req *users_v1.CreateUserRequest) (*users_v1.CreateUserResponse, error) {
	// TODO Правильно ли, что я все заполняю тут? (То, что не принимаю через запрос).
	user := domain.User{
		Name:     req.Info.GetName(),
		Email:    req.Info.GetEmail(),
		Password: req.Info.GetPassword(),
		Role:     "user", // TODO Consider refactoring this.
		Verified: false,
		AccountState: domain.AccountState{
			State: domain.AccountStateActive,
			Since: time.Now().UTC(),
		},
		CreatedAt: time.Now().UTC(),
		LastSeen:  time.Now().UTC(),
	}

	// TODO Сделать свой заворачиватель ошибок, чтобы можно было подряд все ошибки вывести в респонсе.
	var errs []error
	errs = append(errs, user.Validate()...)
	var allErrs error
	for _, err := range errs {
		if allErrs == nil {
			allErrs = err
		} else {
			allErrs = fmt.Errorf("%s\n%s", allErrs, err)
		}
	}
	if allErrs != nil {
		return nil, allErrs
	}

	uid, err := i.usersService.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return &users_v1.CreateUserResponse{
		Uuid: uid.String(),
	}, nil
}
