package domain

import "github.com/google/uuid"

type Tag struct {
	TagID   uuid.UUID `json:"tag_id" db:"tag_id"`
	TagName string    `json:"tag_name" db:"tag_name"`
}

type Tags []Tag

type TagList struct {
	TagID   uuid.UUID `json:"tag_id" db:"tag_id"`
	TagName string    `json:"tag_name" db:"name"`
	TextID  uuid.UUID `json:"text_id" db:"text_id"`
}

type TagLists []TagList
