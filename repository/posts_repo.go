package repository

import (
	"golang-rest-api/model"
)

type PostRepository interface {
	FindAll() ([]model.Post, error)
	Save(post *model.Post) (*model.Post, error)
}
