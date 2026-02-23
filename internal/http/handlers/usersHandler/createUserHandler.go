package usersHandler

import (
	"github.com/DotNicolasPenha/Posts-CRUD/internal/http/responses"
	"github.com/DotNicolasPenha/Posts-CRUD/internal/modules/user"
	"github.com/gin-gonic/gin"
)

func (h *handler) createUserHandler(ctx *gin.Context) {
	var createUserDto user.CreateUserDTO
	err := ctx.ShouldBindJSON(&createUserDto)
	if err != nil {
		responses.BadRequest(ctx, err)
		return
	}
	err = h.service.AddUser(createUserDto)
	if err != nil {
		responses.BadRequest(ctx, err)
		return
	}
	responses.OK_CREATED(ctx, "user created")
}
