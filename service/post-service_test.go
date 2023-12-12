package service

import (
	"go-mux-crash-course/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) Save(post *entity.Post) (*entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Post), args.Error(1)
}

func (mock *MockRepository) FindAll() ([]entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Post), args.Error(1)
}

func (mock *MockRepository) Delete(ID string) error {
	args := mock.Called()
	return args.Error(1)
}

func (mock *MockRepository) FindByID(ID string) (*entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Post), args.Error(1)
}
func TestFindAll(t *testing.T) {
	mockRepo := new(MockRepository)

	identifier := 1

	post := entity.Post{ID: identifier, Title: "A", Text: "B"}

	mockRepo.On("FindAll").Return([]entity.Post{post}, nil)

	testService := NewPostService(mockRepo)

	result, _ := testService.FindAll()
	mockRepo.AssertExpectations(t)

	assert.Equal(t, identifier, result[0].ID)
	assert.Equal(t, "A", result[0].Title)
	assert.Equal(t, "B", result[0].Text)
}

func TestCreate(t *testing.T) {
	mockRepo := new(MockRepository)
	post := entity.Post{ID: 1, Title: "A", Text: "B"}

	mockRepo.On("Save").Return(&post, nil)

	testService := NewPostService(mockRepo)

	result, err := testService.Create(&post)

	mockRepo.AssertExpectations(t)

	assert.NotNil(t, result.ID)
	assert.Equal(t, "A", result.ID)
	assert.Equal(t, "B", result.ID)
	assert.Nil(t, err)
}

func TestValidateEmptyPost(t *testing.T) {
	testService := NewPostService(nil)

	err := testService.Validate(nil)

	assert.NotNil(t, err)

	assert.Equal(t, "The post is empty", err.Error())

}

func TestValidateEmptyPostTitle(t *testing.T) {
	post := entity.Post{ID: 1, Title: "A", Text: "B"}
	testService := NewPostService(nil)
	err := testService.Validate(&post)
	assert.NotNil(t, err)

	assert.Equal(t, "The post title is empty", err.Error())

}
