package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID    uuid.UUID `json:"user_id" db:"user_id" form:"user_id"`
	Name      string    `json:"name" db:"name" form:"name"`
	CreatedAt time.Time `json:"created_at" db:"created_at" form:"created_at"`
}

type Users []User

type UserCreate struct {
	UserID uuid.UUID `json:"user_id" db:"user_id"`
	TextID uuid.UUID `json:"text_id" db:"text_id"`
}

type UserCreates []UserCreate

type UserLike struct {
	UserID uuid.UUID `json:"user_id" db:"user_id"`
	TextID uuid.UUID `json:"text_id" db:"text_id"`
}

type UserLikes []UserLike
