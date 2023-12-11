package repository

import "go-mux-crash-course/entity"

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
	Delete(post *entity.Post) error
	FindByID(ID string) (*entity.Post, error)
}
