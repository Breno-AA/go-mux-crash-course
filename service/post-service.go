package service

import (
	"errors"
	"go-mux-crash-course/entity"
	"go-mux-crash-course/repository"
	"math/rand"
)

type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (int64, error)
	FindAll() ([]entity.Post, error)
	FindByID(ID string) (entity.Post, error)
}

type service struct{}

var (
	repo repository.PostRepository
)

func NewPostService(newRepo repository.PostRepository) PostService {
	repo = newRepo
	return &service{}
}
func (s *service) Validate(post *entity.Post) error {
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
func (s *service) Create(post *entity.Post) (int64, error) {
	post.ID = rand.Int63()
	return repo.Save(post)
}
func (s *service) FindAll() ([]entity.Post, error) {
	return repo.FindAll()
}

func (s *service) FindByID(ID string) (entity.Post, error) {
	return repo.FindByID(ID)
}
