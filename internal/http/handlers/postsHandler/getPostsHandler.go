package postsHandler

import (
	"github.com/DotNicolasPenha/Posts-CRUD/internal/http/responses"
	"github.com/gin-gonic/gin"
)

func (h *handler) getPostsHandler(ctx *gin.Context) {
	posts, err := h.service.GetPosts()
	if err != nil {
		responses.BadRequest(ctx, err)
		return
	}
	responses.OK_DATA(ctx, posts)
}
