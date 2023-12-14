package api

import (
	"ac_blog_api/internal/entity"
	"ac_blog_api/internal/exception"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
)

type Controller struct {
	blog entity.Blog
}

func EncodeError(response *http.ResponseWriter, msg string, status int) {
	httpError := exception.NewHttpError(msg, status)
	json.NewEncoder(*response).Encode(httpError)
}

func NewController(blog entity.Blog) *Controller {
	return &Controller{
		blog: blog,
	}
}

func (c Controller) CreatePost(response http.ResponseWriter, request *http.Request) {
	post := entity.Post{}

	err := json.NewDecoder(request.Body).Decode(&post)
	if err != nil {
		EncodeError(&response, "Could not decode payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	validate := validator.New()
	err = validate.Struct(post)
	if err != nil {
		EncodeError(&response, "Validation: "+err.Error(), http.StatusBadRequest)
		return
	}

	post.Comments = []entity.Comment{}
	err = c.blog.CreatePost(post)
	if err != nil {
		EncodeError(&response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusCreated)
}

func (a *Controller) GetPostList(response http.ResponseWriter, request *http.Request) {
	posts, err := a.blog.GetPostList()
	if err != nil {
		EncodeError(&response, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(posts.Posts) == 0 {
		EncodeError(&response, "No posts found", http.StatusNotFound)
		return
	}

	json.NewEncoder(response).Encode(posts)
}

func (a *Controller) GetPost(response http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")
	if id == "" {
		EncodeError(&response, "Invalid post id", http.StatusBadRequest)
		return
	}

	post, err := a.blog.GetPost(id)
	if err != nil {
		EncodeError(&response, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(response).Encode(post)
}

func (a *Controller) UpdatePost(response http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")
	if id == "" {
		EncodeError(&response, "Invalid post id", http.StatusBadRequest)
		return
	}

	post := entity.Post{}
	err := json.NewDecoder(request.Body).Decode(&post)
	if err != nil {
		EncodeError(&response, "Invalid payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	validate := validator.New()
	err = validate.Struct(post)
	if err != nil {
		EncodeError(&response, "Validation: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = a.blog.UpdatePost(id, post)
	if err != nil {
		EncodeError(&response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
}

func (a *Controller) DeletePost(response http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")
	if id == "" {
		EncodeError(&response, "Invalid post id", http.StatusBadRequest)
		return
	}

	err := a.blog.DeletePost(id)
	if err != nil {
		EncodeError(&response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusNoContent)
}

func (a *Controller) CreateComment(response http.ResponseWriter, request *http.Request) {
	postId := chi.URLParam(request, "id")
	if postId == "" {
		EncodeError(&response, "Invalid post id", http.StatusBadRequest)
		return
	}

	comment := entity.Comment{}
	err := json.NewDecoder(request.Body).Decode(&comment)
	if err != nil {
		EncodeError(&response, "Invalid payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	validate := validator.New()
	err = validate.Struct(comment)
	if err != nil {
		EncodeError(&response, "Validation: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = a.blog.CreateComment(postId, comment)
	if err != nil {
		EncodeError(&response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusCreated)
}

func (a *Controller) GetCommentList(response http.ResponseWriter, request *http.Request) {
	postID := chi.URLParam(request, "id")
	if postID == "" {
		EncodeError(&response, "Invalid post id", http.StatusBadRequest)
		return
	}

	comments, err := a.blog.GetCommentList(postID)
	if err != nil {
		EncodeError(&response, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(response).Encode(comments)
}

func (a *Controller) DeletePostComment(response http.ResponseWriter, request *http.Request) {
	postID := chi.URLParam(request, "postId")
	if postID == "" {
		EncodeError(&response, "Invalid post id", http.StatusBadRequest)
		return
	}

	commentID := chi.URLParam(request, "commentId")
	if commentID == "" {
		EncodeError(&response, "Invalid comment id", http.StatusBadRequest)
		return
	}

	err := a.blog.DeletePostComment(postID, commentID)
	if err != nil {
		EncodeError(&response, err.Error(), http.StatusInternalServerError)
		return
	}
	response.WriteHeader(http.StatusNoContent)
}
