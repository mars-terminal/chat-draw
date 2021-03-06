package message

import (
	"github.com/mars-terminal/chat-draw/internal/service"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service service.MessageService
}

func NewHandler(service service.MessageService) *Handler {
	return &Handler{Service: service}
}

var log = logrus.WithFields(logrus.Fields{
	"package": "message",
	"layer":   "server",
})

func (h *Handler) InitRoutes(r gin.IRouter) {
	group := r.Group("/message")

	group.POST("/create", h.create)
	group.GET("/:chat_id/:peer_id", h.getByChatIDAndPeerID)
	group.GET("/search/:query", h.search)
	group.PUT("/update", h.update)
	group.DELETE("/delete/:id", h.delete)
}
