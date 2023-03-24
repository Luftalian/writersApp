package usecase

import (
	"github.com/google/uuid"

	"github.com/Luftalian/writers_app/domain"
)

type TextInteractor struct {
	TextRepository TextRepository
}

func (interactor *TextInteractor) Add(t domain.Text) (domain.Text, error) {
	text, err := interactor.TextRepository.Store(t)
	if err != nil {
		return domain.Text{}, err
	}
	return text, nil
}

func (interactor *TextInteractor) Texts() (domain.Texts, error) {
	texts, err := interactor.TextRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return texts, nil
}

func (interactor *TextInteractor) TextByID(id uuid.UUID) (domain.Text, error) {
	text, err := interactor.TextRepository.FindByID(id)
	if err != nil {
		return domain.Text{}, err
	}
	return text, nil
}

func (interactor *TextInteractor) TextChange(t domain.Text) error {
	err := interactor.TextRepository.Update(t)
	if err != nil {
		return err
	}
	return nil
}

func (interactor *TextInteractor) TextDelete(id uuid.UUID) error {
	err := interactor.TextRepository.DeleteByID(id)
	if err != nil {
		return err
	}
	return nil
}
