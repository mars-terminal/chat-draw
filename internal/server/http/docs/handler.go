package docs

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mars-terminal/chat-draw/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	port string
}

func NewHandler(port string) *Handler {
	return &Handler{
		port: port,
	}
}

func (h *Handler) InitRoutes(r gin.IRouter) {
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", h.port)
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
