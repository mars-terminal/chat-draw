package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"repositorie/internal/storage"
)

var log = logrus.WithField("package", "handler")

type Handler struct {
	Storage *storage.Storage
}

func NewHandler(storage *storage.Storage) *Handler {
	return &Handler{Storage: storage}
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
