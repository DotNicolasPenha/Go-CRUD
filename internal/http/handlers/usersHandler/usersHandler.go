package usersHandler

import (
	"github.com/DotNicolasPenha/Posts-CRUD/internal/modules/user"
	"github.com/gin-gonic/gin"
)

type handler struct {
	g       *gin.Engine
	service *user.Service
}

func UsersHandler(serviceUser *user.Service, g *gin.Engine) {
	userHandler := g.Group("/users")
	handler := handler{g: g, service: serviceUser}
	{
		userHandler.POST("/", handler.createUserHandler)
		userHandler.GET("/", handler.getUsersHandler)
		userHandler.POST("/login", handler.loginHandler)
	}
}
