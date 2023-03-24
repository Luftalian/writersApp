package usecase

import (
	"github.com/google/uuid"

	"github.com/Luftalian/writers_app/domain"
)

type TextRepository interface {
	Store(text domain.Text) (domain.Text, error)
	FindAll() (domain.Texts, error)
	FindByID(id uuid.UUID) (domain.Text, error)
	Update(text domain.Text) error
	DeleteByID(id uuid.UUID) error
}
