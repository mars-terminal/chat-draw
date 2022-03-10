package user

import (
	"github.com/gin-gonic/gin"
	"repositorie/internal/service"
)

type Handler struct {
	Service service.UserService
}

func NewHandler(service service.UserService) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) InitRoutes(r gin.IRouter) {
	group := r.Group("/users")

	group.GET("/me", h.me)
}
