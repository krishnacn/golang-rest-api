package service

import (
	"errors"
	"golang-rest-api/model"
	"golang-rest-api/repository"

	"github.com/google/uuid"
)

type PostService interface {
	Validate(post *model.Post) error
	Create(post *model.Post) (*model.Post, error)
	FindAll() ([]model.Post, error)
}

type service struct{}

var repo repository.PostRepository

func NewPostService(postRepo repository.PostRepository) PostService {
	repo = postRepo
	return &service{}
}

func (*service) Validate(post *model.Post) error {
	if post == nil {
		return errors.New("the post is empty")
	}
	if post.Title == "" {
		return errors.New("the title of the post is empty")
	}
	return nil
}

func (*service) Create(post *model.Post) (*model.Post, error) {
	post.Id = uuid.New().String()
	return repo.Save(post)
}

func (*service) FindAll() ([]model.Post, error) {
	return repo.FindAll()
}
