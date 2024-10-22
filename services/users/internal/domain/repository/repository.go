package repository

import "dzhordano/132market/services/users/internal/domain/entities"

type UserRepository interface {
	Save(user *entities.User) error
	Update(user *entities.User) error // TODO Нужна ли отдельная структура для обновления?
	Delete(id string) error
	FindById(id string) (*entities.User, error)
	FindAll() ([]*entities.User, error)
}
