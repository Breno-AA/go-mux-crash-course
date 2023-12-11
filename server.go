package main

import (
	"fmt"
	"go-mux-crash-course/controller"
	router "go-mux-crash-course/http"
	"go-mux-crash-course/repository"
	"go-mux-crash-course/service"
	"net/http"
)

var (
	PostRepository repository.PostRepository = repository.NewSQLiteRepository()
	postService    service.PostService       = service.NewPostService(PostRepository)
	postController controller.PostController = controller.NewPostController(postService)
	httpRouter     router.Router             = router.NewMuxRouter()
)

func main() {
	const port string = ":8000"

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Server Running...")
	})

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)

	httpRouter.SERVE(port)

}
