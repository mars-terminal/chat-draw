package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"repositorie/internal/service"
)

var log = logrus.WithField("package", "handler")

type Handler struct {
	authService service.AuthService
}

func NewHandler(authService service.AuthService) *Handler {
	return &Handler{
		authService: authService,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	return router
}
