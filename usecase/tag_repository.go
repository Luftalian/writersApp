package usecase

import (
	"github.com/google/uuid"

	"github.com/Luftalian/writers_app/domain"
)

type TagRepository interface {
	Store(tag domain.Tag) (domain.Tag, error)
	FindAll() (domain.Tags, error)
	FindByID(id uuid.UUID) (domain.Tag, error)
}
