package handlers

import (
	"errors"

	"github.com/DotNicolasPenha/Posts-CRUD/internal/http/responses"
	"github.com/DotNicolasPenha/Posts-CRUD/internal/modules/post"
	"github.com/gin-gonic/gin"
)

func PostsHandler(g *gin.Engine, service *post.Service) {
	postsHandler := g.Group("/posts")
	postsHandler.Use()
	{
		postsHandler.POST("/posts", func(ctx *gin.Context) {
			var postToCreate post.CreatePostDTO
			if err := ctx.ShouldBindJSON(&postToCreate); err != nil {
				responses.BadRequest(ctx, err)
				return
			}
			if err := service.AddPost(postToCreate); err != nil {
				responses.BadRequest(ctx, err)
				return
			}
			responses.OK_CREATED(ctx, "post created")
		})
		postsHandler.GET("/posts", func(ctx *gin.Context) {
			posts, err := service.GetPosts()
			if err != nil {
				responses.BadRequest(ctx, err)
				return
			}
			responses.OK_DATA(ctx, posts)
		})
		postsHandler.GET("/posts/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")
			post, err := service.GetOnePost(id)
			if err != nil {
				responses.BadRequest(ctx, err)
				return
			}
			if post == nil {
				responses.NotFound(ctx, errors.New("post not found"))
				return
			}
			responses.OK_DATA(ctx, post)
		})
		postsHandler.DELETE("/posts/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")
			err := service.RemoveOne(id)
			if err != nil {
				responses.BadRequest(ctx, err)
				return
			}
			responses.OK_DELETE(ctx)
		})
	}
}
