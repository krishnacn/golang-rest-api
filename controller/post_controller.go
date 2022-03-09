package controller

import (
	"encoding/json"
	"golang-rest-api/errors"
	"golang-rest-api/model"
	"golang-rest-api/service"
	"net/http"
)

type PostController interface {
	GetPosts(response http.ResponseWriter, request *http.Request)
	AddPost(response http.ResponseWriter, request *http.Request)
}

type controller struct{}

var postService service.PostService

func NewPostController(service service.PostService) PostController {
	postService = service
	return &controller{}
}

func (*controller) GetPosts(response http.ResponseWriter, request *http.Request) {
	// set content type header
	response.Header().Set("Content-Type", "application/json")

	// get all posts from db
	posts, err := postService.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "error getting posts"})
		return
	}

	// return response
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(posts)
}

func (*controller) AddPost(response http.ResponseWriter, request *http.Request) {
	// set content type header
	response.Header().Set("Content-Type", "application/json")

	// read new post from request body
	var post model.Post
	err := json.NewDecoder(request.Body).Decode(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "error unmarshalling request body"})
		return
	}

	// validate request body
	err = postService.Validate(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: err.Error()})
		return
	}

	// save post in db
	result, err := postService.Create(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "error saving post"})
		return
	}

	// return response
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(result)
}
