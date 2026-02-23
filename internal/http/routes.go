package http

import (
	"github.com/DotNicolasPenha/Posts-CRUD/internal/http/handlers/postsHandler"
	"github.com/DotNicolasPenha/Posts-CRUD/internal/http/handlers/usersHandler"
	"github.com/DotNicolasPenha/Posts-CRUD/internal/modules/post"
	"github.com/DotNicolasPenha/Posts-CRUD/internal/modules/user"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(g *gin.Engine, servicePost *post.Service, serviceUser *user.Service) {
	postsHandler.NewPostHandler(servicePost, g)
	usersHandler.UsersHandler(serviceUser, g)
}
