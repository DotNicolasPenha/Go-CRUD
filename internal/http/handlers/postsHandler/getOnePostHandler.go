package postsHandler

import (
	"errors"

	"github.com/DotNicolasPenha/Posts-CRUD/internal/http/responses"
	"github.com/gin-gonic/gin"
)

func (h *handler) getOnePostHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	post, err := h.service.GetOnePost(id)
	if err != nil {
		responses.BadRequest(ctx, err)
		return
	}
	if post == nil {
		responses.NotFound(ctx, errors.New("post not found"))
		return
	}
	responses.OK_DATA(ctx, post)
}
