package usecase

import (
	"github.com/google/uuid"

	"github.com/Luftalian/writers_app/domain"
)

type TagInteractor struct {
	TagRepository TagRepository
}

func (interactor *TagInteractor) Add(t domain.Tag) (domain.Tag, error) {
	tag, err := interactor.TagRepository.Store(t)
	if err != nil {
		return domain.Tag{}, err
	}
	return tag, nil
}

func (interactor *TagInteractor) Tags() (domain.Tags, error) {
	tags, err := interactor.TagRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (interactor *TagInteractor) TagByID(id uuid.UUID) (domain.Tag, error) {
	tag, err := interactor.TagRepository.FindByID(id)
	if err != nil {
		return domain.Tag{}, nil
	}
	return tag, nil
}
