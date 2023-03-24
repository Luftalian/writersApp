package usecase

import (
	"github.com/google/uuid"

	"github.com/Luftalian/writers_app/domain"
)

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u domain.User) (domain.User, error) {
	user, err := interactor.UserRepository.Store(u)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (interactor *UserInteractor) Users() (domain.Users, error) {
	users, err := interactor.UserRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (interactor *UserInteractor) UserByID(id uuid.UUID) (domain.User, error) {
	user, err := interactor.UserRepository.FindByID(id)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
