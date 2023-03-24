package database

import (
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/Luftalian/writers_app/domain"
)

type TextRepository struct {
	DbHandler
}

func (repo *TextRepository) FindAll() (domain.Texts, error) {
	texts := domain.Texts{}
	err := repo.Select(&texts, "SELECT * FROM texts")
	if err != nil {
		return nil, err
	}
	return texts, nil
}

func (repo *TextRepository) FindByID(id uuid.UUID) (domain.Text, error) {
	texts := domain.Texts{}
	err := repo.Select(&texts, "SELECT * FROM texts WHERE text_id = ?", id)
	if err != nil {
		return domain.Text{}, err
	}
	if len(texts) == 0 {
		return domain.Text{}, errors.New("not found")
	}
	return texts[0], nil
}

func (repo *TextRepository) Store(text domain.Text) (domain.Text, error) {
	text.TextID = uuid.New()
	text.CreatedAt = time.Now()
	text.ChangedAt = text.CreatedAt
	text.GoodCount = 0
	text.BadCount = 0
	_, err := repo.Exec("INSERT INTO texts (text_id, title, content, user_name, user_id, created_at, changed_at, good_count, bad_count) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", text.TextID, text.Title, text.Content, text.UserName, text.UserID, text.CreatedAt, text.ChangedAt, text.GoodCount, text.BadCount)
	if err != nil {
		return domain.Text{}, err
	}
	return text, nil
}

func (repo *TextRepository) Update(text domain.Text) error {
	text.ChangedAt = time.Now()
	_, err := repo.Exec("UPDATE texts SET title = ?, content = ?, changed_at = ?, good_count = ?, bad_count = ? WHERE text_id = ?", text.Title, text.Content, text.ChangedAt, text.GoodCount, text.BadCount, text.TextID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *TextRepository) DeleteByID(id uuid.UUID) error {
	_, err := repo.Exec("DELETE FROM texts WHERE text_id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
