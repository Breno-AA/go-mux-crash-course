package service

import (
	"errors"
	"go-mux-crash-course/entity"
	"go-mux-crash-course/repository"
	"math/rand"
)

type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type service struct{}

var (
	repo repository.PostRepository
)

func NewPostService(repo repository.PostRepository) PostService {
	repo = repo
	return &service{}
}
func (*service) Validate(post *entity.Post) error {
	if post == nil {
		err := errors.New("The post is empty")
		return err
	}
	if post.Title == "" {
		err := errors.New("The post title empty")
		return err
	}
	if post.Text == "" {
		err := errors.New("The post text empty")
		return err
	}
	return nil
}
func (*service) Create(post *entity.Post) (*entity.Post, error) {
	post.ID = rand.Int()
	return repo.Save(post)
}
func (*service) FindAll() ([]entity.Post, error) {
	return repo.FindAll()
}