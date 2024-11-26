package mapper

import (
	"dzhordano/132market/services/users/internal/application/model"
	"dzhordano/132market/services/users/internal/domain/entities"
)

func NewUserResultFromEntity(entity *entities.User) *model.UserResult {
	if entity == nil {
		return nil
	}

	return &model.UserResult{
		ID:         entity.ID,
		Name:       entity.Name,
		Email:      entity.Email,
		Roles:      entity.RolesToStrings(),
		Status:     entity.Status.String(),
		State:      entity.State.String(),
		LastSeenAt: entity.LastSeenAt,
		CreatedAt:  entity.CreatedAt,
	}
}

func NewUserResultListFromEntities(entities []*entities.User) []*model.UserResult {
	var userResults []*model.UserResult
	for _, entity := range entities {
		userResults = append(userResults, NewUserResultFromEntity(entity))
	}
	return userResults
}
