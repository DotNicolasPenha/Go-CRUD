package handlers

import (
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
				ctx.JSON(400, gin.H{
					"error": err.Error(),
					"ok":    false,
				})
				return
			}
			if err := service.AddPost(postToCreate); err != nil {
				ctx.JSON(400, gin.H{
					"error": err.Error(),
					"ok":    false,
				})
				return
			}
			ctx.JSON(201, gin.H{
				"msg": "post created",
				"ok":  true,
			})
		})
		postsHandler.GET("/posts", func(ctx *gin.Context) {
			posts, err := service.GetPosts()
			if err != nil {
				ctx.JSON(500, gin.H{
					"error": err.Error(),
					"ok":    false,
				})
				return
			}
			ctx.JSON(200, gin.H{
				"posts": posts,
				"ok":    true,
			})
		})
		postsHandler.GET("/posts/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")
			post, err := service.GetOnePost(id)
			if err != nil {
				ctx.JSON(500, gin.H{
					"error": err.Error(),
					"ok":    false,
				})
				return
			}
			if post == nil {
				ctx.JSON(404, gin.H{
					"error": "post not found",
					"ok":    false,
				})
				return
			}

			ctx.JSON(200, gin.H{
				"post": post,
				"ok":   true,
			})
		})
		postsHandler.DELETE("/posts/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")
			err := service.RemoveOne(id)
			if err != nil {
				ctx.JSON(400, gin.H{
					"error": err.Error(),
					"ok":    false,
				})
				return
			}
			ctx.JSON(200, gin.H{
				"post_deleted": id,
				"ok":           true,
			})
		})
	}
}
