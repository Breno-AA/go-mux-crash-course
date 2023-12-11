package controller

import (
	"encoding/json"
	"go-mux-crash-course/cache"
	"go-mux-crash-course/entity"
	"go-mux-crash-course/errors"
	"go-mux-crash-course/service"
	"net/http"
	"strings"
)

type controller struct{}

var (
	postService service.PostService
	postCache   cache.PostCache
)

type PostController interface {
	GetPosts(response http.ResponseWriter, request *http.Request)
	AddPost(response http.ResponseWriter, request *http.Request)
	GetPostByID(response http.ResponseWriter, request *http.Request)
}

func NewPostController(service service.PostService, cache cache.PostCache) PostController {
	postService = service
	postCache = cache
	return &controller{}
}

func (c *controller) GetPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	posts, err := postService.FindAll()

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error getting the posts"})

	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(posts)
}

func (c *controller) AddPost(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	var post entity.Post
	err := json.NewDecoder(request.Body).Decode(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error unmarshalling data"})
		return
	}

	err1 := postService.Validate(&post)
	if err1 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: err1.Error()})

	}

	result, err2 := postService.Create(&post)
	if err2 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error saving the post"})

	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)

}

func (c *controller) GetPostByID(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	postID := strings.Split(request.URL.Path, "/")[2]
	var post *entity.Post = postCache.Get(postID)
	if post == nil {
		post, err := postService.FindByID(postID)
		if err != nil {
			response.WriteHeader(http.StatusNotFound)
			json.NewEncoder(response).Encode(errors.ServiceError{Message: "No posts found!"})
			return
		}
		postCache.Set(postID, post)
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(post)
	} else {
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(post)
	}
}
