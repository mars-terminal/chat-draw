package auth

import (
	"net/http"
	"strings"

	"github.com/mars-terminal/chat-draw/internal/entities"
	"github.com/mars-terminal/chat-draw/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/go-http-utils/headers"
)

type Handler struct {
	Service service.AuthService
}

func NewHandler(service service.AuthService) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) InitRoutes(r gin.IRouter) {
	group := r.Group("/auth")

	group.POST("/sign-in", h.signIn)
	group.POST("/sign-up", h.signUp)
}

func (h *Handler) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := strings.Split(c.GetHeader(headers.Authorization), " ")

		if len(authHeader) != 2 && authHeader[0] != TokenPrefix {
			c.AbortWithStatusJSON(http.StatusUnauthorized, &entities.Response{
				Message: "invalid auth header",
				Status:  http.StatusUnauthorized,
			})
			return
		}

		token := authHeader[1]

		user, err := h.Service.GetUserByToken(c, token)
		if err != nil {
			return
		}

		c.Set(gin.AuthUserKey, user)
	}
}
