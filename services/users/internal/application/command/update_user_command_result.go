package command

import "github.com/dzhordano/132market/services/users/internal/application/model"

// TODO Нужен ли этот ответ?? (именно модель)
type UpdateUserCommandResult struct {
	Result *model.UserResult
}
