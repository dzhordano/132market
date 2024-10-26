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
		Id:         entity.ID,
		Username:   entity.Name,
		Email:      entity.Email,
		Roles:      entity.RolesToStrings(),
		Status:     entity.Status.String(),
		State:      entity.State.String(),
		LastSeenAt: entity.LastSeenAt,
		CreatedAt:  entity.CreatedAt,
	}
}
