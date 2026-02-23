package postsHandler

import (
	"github.com/DotNicolasPenha/Posts-CRUD/internal/http/responses"
	"github.com/gin-gonic/gin"
)

func (h *handler) deletePostHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	err := h.service.RemoveOne(id)
	if err != nil {
		responses.BadRequest(ctx, err)
		return
	}
	responses.OK_DELETE(ctx)
}
