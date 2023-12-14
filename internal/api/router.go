package api

import (
	"ac_blog_api/internal/entity"

	"github.com/go-chi/chi"
)

func Router(postRepository entity.PostRepository) *chi.Mux {

	blog := entity.NewBlog(postRepository)
	controller := NewController(blog)

	router := chi.NewRouter()
	router.Post("/posts", controller.CreatePost)
	router.Get("/posts", controller.GetPostList)
	router.Get("/posts/{id}", controller.GetPost)
	router.Put("/posts/{id}", controller.UpdatePost)
	router.Delete("/posts/{id}", controller.DeletePost)
	router.Post("/posts/{id}/comments", controller.CreateComment)
	router.Get("/posts/{id}/comments", controller.GetCommentList)
	router.Delete("/posts/{postId}/comments/{commentId}", controller.DeletePostComment)
	return router
}
