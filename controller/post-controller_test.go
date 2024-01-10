package controller_test

import (
	"bytes"
	"go-mux-crash-course/controller"
	"go-mux-crash-course/repository"
	"go-mux-crash-course/service"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	postRepo       repository.PostRepository = repository.NewMySQLRepository()
	postSrv        service.PostService       = service.NewPostService(postRepo)
	postController controller.PostController = controller.NewPostController(postSrv)
)

func TestAddPost(t *testing.T) {
	book := readTestData(t, "addPost.json")
	bookReader := bytes.NewReader(book)

	req := httptest.NewRequest(http.MethodPost, "/posts", bookReader)
	w := httptest.NewRecorder()
	postController.AddPost(w, req)

	res := w.Result()
	defer res.Body.Close()
	assert.Equal(t, 200, res.StatusCode)
}

func TestGetPosts(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/posts", nil)
	w := httptest.NewRecorder()
	postController.GetPosts(w, req)

	res := w.Result()
	defer res.Body.Close()
	assert.Equal(t, 200, res.StatusCode)
}

func readTestData(t *testing.T, name string) []byte {
	t.Helper()
	content, err := os.ReadFile("../" + name)
	if err != nil {
		t.Errorf("Could not read %v", name)
	}

	return content
}
