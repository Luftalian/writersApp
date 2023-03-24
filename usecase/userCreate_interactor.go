package usecase

import (
	"github.com/google/uuid"

	"github.com/Luftalian/writers_app/domain"
)

type UserCreateInteractor struct {
	UserCreateRepository UserCreateRepository
}

func (interactor *UserCreateInteractor) Add(u domain.UserCreate) (domain.UserCreate, error) {
	userCreate, err := interactor.UserCreateRepository.Store(u)
	if err != nil {
		return domain.UserCreate{}, err
	}
	return userCreate, nil
}

func (interactor *UserCreateInteractor) UserCreates() (domain.UserCreates, error) {
	userCreates, err := interactor.UserCreateRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return userCreates, nil
}

func (interactor *UserCreateInteractor) UserCreateByUserID(userID uuid.UUID) (domain.Texts, error) {
	texts, err := interactor.UserCreateRepository.FindByUserID(userID)
	if err != nil {
		return nil, err
	}
	return texts, nil
}

func (interactor *UserCreateInteractor) UserCreateByTextID(textID uuid.UUID) (domain.UserCreate, error) {
	userCreate, err := interactor.UserCreateRepository.FindByTextID(textID)
	if err != nil {
		return domain.UserCreate{}, err
	}
	return userCreate, nil
}

func (interactor *UserCreateInteractor) Change(u domain.UserCreate) error {
	err := interactor.UserCreateRepository.Update(u)
	if err != nil {
		return err
	}
	return nil
}

func (interactor *UserCreateInteractor) Delete(u domain.UserCreate) error {
	err := interactor.UserCreateRepository.Delete(u)
	if err != nil {
		return err
	}
	return nil
}

func (interactor *UserCreateInteractor) DeleteByUserID(userID uuid.UUID) error {
	err := interactor.UserCreateRepository.DeleteByUserID(userID)
	if err != nil {
		return err
	}
	return nil
}

func (interactor *UserCreateInteractor) DeleteByTextID(textID uuid.UUID) error {
	err := interactor.UserCreateRepository.DeleteByTextID(textID)
	if err != nil {
		return err
	}
	return nil
}
