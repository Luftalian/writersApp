package database

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"

	"github.com/Luftalian/writers_app/domain"
)

type UserLikeRepository struct {
	DbHandler
}

func (repo *UserLikeRepository) FindAll() (domain.UserLikes, error) {
	userLikes := domain.UserLikes{}
	err := repo.Select(&userLikes, "SELECT * FROM userLikes")
	if err != nil {
		return nil, err
	}
	return userLikes, nil
}

func (repo *UserLikeRepository) FindByUserID(userID uuid.UUID) (domain.Texts, error) {
	texts := domain.Texts{}
	err := repo.Select(&texts, "SELECT texts.text_id, texts.title, texts.content, texts.user_name, texts.user_id, texts.created_at, texts.changed_at, texts.good_count, texts.bad_count FROM userLikes INNER JOIN texts ON userLikes.text_id = texts.text_id WHERE userLikes.user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	return texts, nil
}

func (repo *UserLikeRepository) FindByTextID(textID uuid.UUID) (domain.UserLike, error) {
	userLikes := domain.UserLikes{}
	err := repo.Select(&userLikes, "SELECT * FROM userLikes WHERE text_id = ?", textID)
	if err != nil {
		return domain.UserLike{}, err
	}
	return userLikes[0], nil
}

func (repo *UserLikeRepository) Store(userLike domain.UserLike) (domain.UserLike, error) {
    err := repo.Get(&userLike, "SELECT * FROM userLikes WHERE user_id = ? AND text_id = ?", userLike.UserID, userLike.TextID)
    if err == nil {
		return userLike, errors.New("UserLike already exists")
	}
	if err != sql.ErrNoRows {
		return userLike, err
    }
	_, err = repo.Exec("INSERT INTO userLikes (user_id, text_id) VALUES (?, ?)", userLike.UserID, userLike.TextID)
	if err != nil {
		return domain.UserLike{}, err
	}
	return userLike, nil
}

func (repo *UserLikeRepository) CheckLikeUser(userLike domain.UserLike) (bool, error) {
	err := repo.Get(&userLike, "SELECT * FROM userLikes WHERE user_id = ? AND text_id = ?", userLike.UserID, userLike.TextID)
	if err == sql.ErrNoRows {
		return false, err
    }
	if err != nil {
		return false, errors.New("UserLike already exists")
	}
	return true, nil
}

func (repo *UserLikeRepository) Update(userLike domain.UserLike) error {
	_, err := repo.Exec("UPDATE userLikes SET user_id = ?, text_id = ? WHERE user_id = ? AND text_id = ?", userLike.UserID, userLike.TextID, userLike.UserID, userLike.TextID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserLikeRepository) Delete(u domain.UserLike) error {
	_, err := repo.Exec("DELETE FROM userLikes WHERE user_id = ? AND text_id = ?", u.UserID, u.TextID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserLikeRepository) DeleteByUserID(userID uuid.UUID) error {
	_, err := repo.Exec("DELETE FROM userLikes WHERE user_id = ?", userID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserLikeRepository) DeleteByTextID(textID uuid.UUID) error {
	_, err := repo.Exec("DELETE FROM userLikes WHERE text_id = ?", textID)
	if err != nil {
		return err
	}
	return nil
}
