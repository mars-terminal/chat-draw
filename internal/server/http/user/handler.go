package user

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"repositorie/internal/service"
)

var log = logrus.WithFields(logrus.Fields{
	"package": "auth",
	"layer":   "server",
})

type Handler struct {
	Service service.UserService
}

func NewHandler(service service.UserService) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) InitRoutes(r gin.IRouter) {
	group := r.Group("/users")

	group.GET("/me", h.me)
	group.GET("/search/:query", h.search)
	group.GET("/settings", h.settings)
	group.PUT("/update", h.update)
}
