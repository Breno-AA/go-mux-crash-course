package controller

import (
	"encoding/json"
	"fmt"
	"go-mux-crash-course/cache"
	"go-mux-crash-course/entity"
	"go-mux-crash-course/errors"
	"go-mux-crash-course/service"
	"net/http"
	"strings"
)

type controller struct {
	postService service.PostService
	postCache   cache.PostCache
}

type PostController interface {
	GetPosts(response http.ResponseWriter, request *http.Request)
	AddPost(response http.ResponseWriter, request *http.Request)
	GetPostByID(response http.ResponseWriter, request *http.Request)
}

func NewPostController(service service.PostService, cache cache.PostCache) PostController {
	return &controller{
		postService: service,
		postCache:   cache,
	}
}

func (c *controller) GetPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")

	posts, err := c.postService.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error getting the posts"})
		return
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

	err = c.postService.Validate(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: err.Error()})
		return
	}

	postID, err := c.postService.Create(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error saving the post"})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode([]byte(fmt.Sprint(postID)))
}

func (c *controller) GetPostByID(response http.ResponseWriter, request *http.Request) {
	postID := strings.Split(request.URL.Path, "/")[2]

	response.Header().Set("Content-type", "application/json")
	cached := c.postCache.Get(postID)

	if cached != nil {
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(cached)
		return
	}

	post, err := c.postService.FindByID(postID)
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "No posts found!"})
		return
	}
	c.postCache.Set(postID, &post)

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(post)
}
