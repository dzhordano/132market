package query

import (
	"github.com/dzhordano/132market/services/users/internal/application/model"
)

type UserQueryResult struct {
	Result *model.UserResult
}

type UserQueryListResult struct {
	Result []*model.UserResult
	Count  uint64
}
