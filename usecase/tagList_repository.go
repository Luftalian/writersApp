package usecase

import (
	"github.com/google/uuid"

	"github.com/Luftalian/writers_app/domain"
)

type TagListRepository interface {
	Store(tag domain.TagList) (domain.TagList, error)
	FindAll() (domain.TagLists, error)
	FindByTextID(textID uuid.UUID) (domain.TagLists, error)
	FindByTagID(tagID uuid.UUID) (domain.TagLists, error)
	FindByName(name string) (domain.Texts, error)
	Update(tag domain.TagList) (domain.TagList, error)
	Delete(id uuid.UUID) error
}
