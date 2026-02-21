package handlers

import (
	"github.com/DotNicolasPenha/Posts-CRUD/internal/modules/user"
	"github.com/gin-gonic/gin"
)

func UsersHandler(g *gin.Engine, serviceUser *user.Service) {
	g.POST("/users", func(ctx *gin.Context) {
		var createUserDto user.CreateUserDTO
		err := ctx.ShouldBindJSON(&createUserDto)
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": err.Error(),
				"ok":    false,
			})
			return
		}
		err = serviceUser.AddUser(createUserDto)
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": err.Error(),
				"ok":    false,
			})
			return
		}
		ctx.JSON(201, gin.H{
			"msg": "user created",
			"ok":  true,
		})
	})
	g.GET("/users", func(ctx *gin.Context) {
		users, err := serviceUser.GetUsers()
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
				"ok":    false,
			})
			return
		}
		ctx.JSON(200, gin.H{
			"users": users,
			"ok":    true,
		})
	})
}
