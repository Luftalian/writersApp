package usecase

import (
	"github.com/google/uuid"

	"github.com/Luftalian/writers_app/domain"
)

type UserCreateRepository interface {
	Store(userCreate domain.UserCreate) (domain.UserCreate, error)
	FindAll() (domain.UserCreates, error)
	FindByUserID(userID uuid.UUID) (domain.Texts, error)
	FindByTextID(textID uuid.UUID) (domain.UserCreate, error)
	Update(userCreate domain.UserCreate) error
	Delete(u domain.UserCreate) error
	DeleteByUserID(userID uuid.UUID) error
	DeleteByTextID(textID uuid.UUID) error
}
