package postsHandler

import (
	"github.com/DotNicolasPenha/Posts-CRUD/internal/http/responses"
	"github.com/DotNicolasPenha/Posts-CRUD/internal/modules/post"
	"github.com/gin-gonic/gin"
)

func (h *handler) createPostHandler(ctx *gin.Context) {
	var postToCreate post.CreatePostDTO
	if err := ctx.ShouldBindJSON(&postToCreate); err != nil {
		responses.BadRequest(ctx, err)
		return
	}
	if err := h.service.AddPost(postToCreate); err != nil {
		responses.BadRequest(ctx, err)
		return
	}
	responses.OK_CREATED(ctx, "post created")
}
