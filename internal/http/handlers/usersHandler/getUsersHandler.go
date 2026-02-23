package usersHandler

import (
	"github.com/DotNicolasPenha/Posts-CRUD/internal/http/responses"
	"github.com/gin-gonic/gin"
)

func (h *handler) getUsersHandler(ctx *gin.Context) {
	users, err := h.service.GetUsers()
	if err != nil {
		responses.BadRequest(ctx, err)
		return
	}
	responses.OK_DATA(ctx, users)
}
