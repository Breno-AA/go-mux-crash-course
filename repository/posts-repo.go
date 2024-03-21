package repository

import "go-mux-crash-course/entity"

type PostRepository interface {
	Save(post *entity.Post) (int64, error)
	FindAll() ([]entity.Post, error)
	Delete(ID string) error
	FindByID(ID string) (entity.Post, error)
}
