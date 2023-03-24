package database

import (
	"github.com/google/uuid"

	"github.com/Luftalian/writers_app/domain"
)

type UserCreateRepository struct {
	DbHandler
}

func (repo *UserCreateRepository) FindAll() (domain.UserCreates, error) {
	userCreates := domain.UserCreates{}
	err := repo.Select(&userCreates, "SELECT * FROM userCreates")
	if err != nil {
		return nil, err
	}
	return userCreates, nil
}

func (repo *UserCreateRepository) FindByUserID(userID uuid.UUID) (domain.Texts, error) {
	texts := domain.Texts{}
	err := repo.Select(&texts, "SELECT texts.text_id, texts.title, texts.content, texts.user_name, texts.user_id, texts.created_at, texts.changed_at, texts.good_count, texts.bad_count FROM userCreates INNER JOIN texts ON userCreates.text_id = texts.text_id WHERE userCreates.user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	return texts, nil
}

func (repo *UserCreateRepository) FindByTextID(textID uuid.UUID) (domain.UserCreate, error) {
	userCreates := domain.UserCreates{}
	err := repo.Select(&userCreates, "SELECT * FROM userCreates WHERE text_id = ?", textID)
	if err != nil {
		return domain.UserCreate{}, err
	}
	return userCreates[0], nil
}

func (repo *UserCreateRepository) Store(userCreate domain.UserCreate) (domain.UserCreate, error) {
	_, err := repo.Exec("INSERT INTO userCreates (user_id, text_id) VALUES (?, ?)", userCreate.UserID, userCreate.TextID)
	if err != nil {
		return domain.UserCreate{}, err
	}
	return userCreate, nil
}

func (repo *UserCreateRepository) Update(userCreate domain.UserCreate) error {
	_, err := repo.Exec("UPDATE userCreates SET user_id = ?, text_id = ? WHERE user_id = ? AND text_id = ?", userCreate.UserID, userCreate.TextID, userCreate.UserID, userCreate.TextID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserCreateRepository) Delete(u domain.UserCreate) error {
	_, err := repo.Exec("DELETE FROM userCreates WHERE user_id = ? AND text_id = ?", u.UserID, u.TextID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserCreateRepository) DeleteByUserID(userID uuid.UUID) error {
	_, err := repo.Exec("DELETE FROM userCreates WHERE user_id = ?", userID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserCreateRepository) DeleteByTextID(textID uuid.UUID) error {
	_, err := repo.Exec("DELETE FROM userCreates WHERE text_id = ?", textID)
	if err != nil {
		return err
	}
	return nil
}
