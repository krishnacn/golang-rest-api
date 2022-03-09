package service

import (
	"golang-rest-api/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) FindAll() ([]model.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]model.Post), args.Error(1)

}
func (mock *MockRepository) Save(post *model.Post) (*model.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*model.Post), args.Error(1)
}

func TestService_ValidateEmptyPost(t *testing.T) {
	postService := NewPostService(nil)
	err := postService.Validate(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "the post is empty", err.Error())
}

func TestService_ValidatePostWithEmptyTitle(t *testing.T) {
	post := model.Post{
		Id:    "",
		Title: "",
		Text:  "",
	}
	postService := NewPostService(nil)
	err := postService.Validate(&post)

	assert.NotNil(t, err)
	assert.Equal(t, "the title of the post is empty", err.Error())
}

func TestService_FindAll(t *testing.T) {
	mockPost := model.Post{
		Id:    "mock_id",
		Title: "mock_title",
		Text:  "mock_text",
	}

	// mock repo method FindAll
	mockRepo := new(MockRepository)
	mockRepo.On("FindAll").Return([]model.Post{mockPost}, nil)

	testService := NewPostService(mockRepo)
	result, err := testService.FindAll()

	mockRepo.AssertExpectations(t)
	assert.Equal(t, 1, len(result))
	assert.Equal(t, mockPost.Id, result[0].Id)
	assert.Equal(t, mockPost.Title, result[0].Title)
	assert.Equal(t, mockPost.Text, result[0].Text)
	assert.Nil(t, err)
}

func TestService_Create(t *testing.T) {
	mockPost := model.Post{
		Id:    "mock_id",
		Title: "mock_title",
		Text:  "mock_text",
	}

	// mock repo method Save
	mockRepo := new(MockRepository)
	mockRepo.On("Save").Return(&mockPost, nil)

	testService := NewPostService(mockRepo)
	result, err := testService.Create(&mockPost)

	mockRepo.AssertExpectations(t)
	assert.NotNil(t, result.Id)
	assert.Equal(t, mockPost.Title, result.Title)
	assert.Equal(t, mockPost.Text, result.Text)
	assert.Nil(t, err)
}
