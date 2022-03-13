package message

import (
	"repositorie/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service service.MessageService
}

func NewHandler(service service.MessageService) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) InitRoutes(r gin.IRouter) {
	group := r.Group("/message")

	group.GET("/something")
}
