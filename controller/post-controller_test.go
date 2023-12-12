package controller

import (
	"bytes"
	"encoding/json"
	"go-mux-crash-course/cache"
	"go-mux-crash-course/entity"
	"go-mux-crash-course/repository"
	"go-mux-crash-course/service"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	ID    int64  = 123
	TITLE string = "Title 1"
	TEXT  string = "Text 1"
)

var (
	postRepo       repository.PostRepository = repository.NewSQLiteRepository()
	postSrv        service.PostService       = service.NewPostService(postRepo)
	postCacheSrv   cache.PostCache           = cache.NewRedisCache("localhost:6379", 0, 10)
	postController PostController            = NewPostController(postSrv, postCacheSrv)
)

func TestAddPost(t *testing.T) {
	var jsonReq = []byte(`{"title": "` + TITLE + `", "text": "` + TEXT + `"}`)
	req, _ := http.NewRequest("POST", "/posts", bytes.NewBuffer(jsonReq))

	handler := http.HandlerFunc(postController.AddPost)

	response := httptest.NewRecorder()

	handler.ServeHTTP(response, req)

	status := response.Code

	if status != http.StatusOK {
		t.Errorf("Handler return a wrong status code: got %v want %v", status, http.StatusOK)
	}

	var posts []entity.Post
	json.NewDecoder(io.Reader(response.Body)).Decode(&posts)

	assert.NotNil(t, posts[0].ID)
	assert.Equal(t, TITLE, posts[0].Title)
	assert.Equal(t, TEXT, posts[0].Text)

	cleanUp(&posts[0])

}

func TestGetPosts(t *testing.T) {

	setup()

	req, _ := http.NewRequest("GET", "/posts", nil)

	handler := http.HandlerFunc(postController.GetPosts)

	response := httptest.NewRecorder()

	handler.ServeHTTP(response, req)

	status := response.Code
	if status != http.StatusOK {
		t.Errorf("Handler return a wrong status code: got %v want %v", status, http.StatusOK)
	}

	var posts []entity.Post
	json.NewDecoder(io.Reader(response.Body)).Decode(&posts)

	assert.NotNil(t, posts[0].ID)
	assert.Equal(t, TITLE, posts[0].Title)
	assert.Equal(t, TEXT, posts[0].Text)

	cleanUp(&posts[0])
}

func TestGetPostByID(t *testing.T) {

	setup()

	req, _ := http.NewRequest("GET", "/posts/"+strconv.FormatInt(ID, 10), nil)

	handler := http.HandlerFunc(postController.GetPostByID)

	response := httptest.NewRecorder()

	handler.ServeHTTP(response, req)

	status := response.Code
	if status != http.StatusOK {
		t.Errorf("Handler return a wrong status code: got %v want %v", status, http.StatusOK)
	}

	var post entity.Post
	json.NewDecoder(io.Reader(response.Body)).Decode(&post)

	assert.NotNil(t, post.ID)
	assert.Equal(t, TITLE, post.Title)
	assert.Equal(t, TEXT, post.Text)

	cleanUp(&post)
}

func setup() {
	var post entity.Post = entity.Post{
		ID:    int(ID),
		Title: TITLE,
		Text:  TEXT,
	}

	postRepo.Save(&post)
}

func cleanUp(post *entity.Post) {
	postRepo.Delete(string(post.ID))
}
