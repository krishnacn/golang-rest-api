package main

import (
	"golang-rest-api/controller"
	router "golang-rest-api/http"
	"golang-rest-api/repository"
	"golang-rest-api/service"
)

var (
	httpRouter     router.Router             = router.NewMuxRouter()
	postRepo       repository.PostRepository = repository.NewMongoDBRepository()
	postService    service.PostService       = service.NewPostService(postRepo)
	postController controller.PostController = controller.NewPostController(postService)
)

func main() {
	const port = ":8080"

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)

	httpRouter.SERVE(port)
}
