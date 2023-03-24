package usecase

import (
	"github.com/google/uuid"

	"github.com/Luftalian/writers_app/domain"
)

type TagListInteractor struct {
	TagListRepository TagListRepository
}

func (interactor *TagListInteractor) Add(t domain.TagList) (domain.TagList, error) {
	tag, err := interactor.TagListRepository.Store(t)
	if err != nil {
		return domain.TagList{}, err
	}
	return tag, nil
}

func (interactor *TagListInteractor) TagLists() (domain.TagLists, error) {
	tags, err := interactor.TagListRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (interactor *TagListInteractor) TagListByTextID(id uuid.UUID) (domain.TagLists, error) {
	tags, err := interactor.TagListRepository.FindByTextID(id)
	if err != nil {
		return nil, nil
	}
	return tags, nil
}

func (interactor *TagListInteractor) TagListByTagID(id uuid.UUID) (domain.TagLists, error) {
	tags, err := interactor.TagListRepository.FindByTagID(id)
	if err != nil {
		return nil, nil
	}
	return tags, nil
}

func (interactor *TagListInteractor) TagListByName(name string) (domain.Texts, error) {
	texts, err := interactor.TagListRepository.FindByName(name)
    if err!= nil {
        return nil, nil
    }
    return texts, nil
}

func (interactor *TagListInteractor) Update(t domain.TagList) (domain.TagList, error) {
	tag, err := interactor.TagListRepository.Update(t)
	if err != nil {
		return domain.TagList{}, err
	}
	return tag, nil
}

func (interactor *TagListInteractor) Delete(id uuid.UUID) error {
	err := interactor.TagListRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
