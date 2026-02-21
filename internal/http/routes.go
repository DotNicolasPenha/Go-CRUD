package http

import (
	"github.com/DotNicolasPenha/Posts-CRUD/internal/http/handlers"
	"github.com/DotNicolasPenha/Posts-CRUD/internal/modules/post"
	"github.com/DotNicolasPenha/Posts-CRUD/internal/modules/user"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(g *gin.Engine, servicePost *post.Service, serviceUser *user.Service) {
	handlers.PostsHandler(g, servicePost)
	handlers.UsersHandler(g, serviceUser)
}
