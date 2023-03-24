package database

import (
	"github.com/google/uuid"

	"github.com/Luftalian/writers_app/domain"
)

type TagRepository struct {
	DbHandler
}

func (repo *TagRepository) FindAll() (domain.Tags, error) {
	tags := domain.Tags{}
	err := repo.Select(&tags, "SELECT * FROM tags")
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (repo *TagRepository) FindByID(id uuid.UUID) (domain.Tag, error) {
	tag := domain.Tag{}
	err := repo.Select(&tag, "SELECT * FROM tags WHERE tag_id = ?", id)
	if err != nil {
		return domain.Tag{}, err
	}
	return tag, nil
}

func (repo *TagRepository) Store(tag domain.Tag) (domain.Tag, error) {
	tag.TagID = uuid.New()
	_, err := repo.Exec("INSERT INTO tags (tag_id, name) VALUES (?, ?)", tag.TagID, tag.TagName)
	if err != nil {
		return domain.Tag{}, err
	}
	return tag, nil
}
