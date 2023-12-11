package controller

import (
	"go-mux-crash-course/entity"
	"go-mux-crash-course/repository"
	"go-mux-crash-course/service"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	postRepo       repository.PostRepository = repository.NewSQLiteRepository()
	postSrv        service.PostService       = service.NewPostService(postRepo)
	postController PostController            = NewPostController(postSrv)
)

func TestAddPost(t *testing.T) {
	var json = []byte(`{"title": "Title 1", "text": "Text 1"}`)
	http.NewRequest("POST", "/posts", byter.NewBuffer(json))

	handler := http.HandlerFunc()

	response := httptest.NewRecorder()

	handler.ServeHTTP(response, req)

	status := response.Code
	if status != http.StatusOK {
		t.Errorf("Handler return a wrong status code: got %v want %v", status, http.StatusOK)

		var post entity.Post
		json.NewDecoder(io.Reader(response.Body)).Decode(&)

		assert.NotNil(t,post.ID)
	}
}

func TestGetPosts(t *testing.T) {

}
