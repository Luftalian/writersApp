package domain

import (
	"time"

	"github.com/google/uuid"
)

type Text struct {
	TextID    uuid.UUID `json:"text_id" db:"text_id"`
	Title     string    `json:"title" db:"title"`
	Content   string    `json:"content" db:"content"`
	UserName  string    `json:"user_name" db:"user_name"`
	UserID    uuid.UUID `json:"user_id" db:"user_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	ChangedAt time.Time `json:"changed_at" db:"changed_at"`
	GoodCount int       `json:"good_count" db:"good_count"`
	BadCount  int       `json:"bad_count" db:"bad_count"`
}

type Texts []Text
