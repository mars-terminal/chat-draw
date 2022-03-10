package user

import (
	"github.com/gin-gonic/gin"
	"repositorie/internal/entities/user"
)

func (h *Handler) me(c *gin.Context) {
	userInterface, exists := c.Get(gin.AuthUserKey)
	if exists {

	}

	u, ok := userInterface.(*user.User)
	if !ok {

	}

	c.BindJSON(u)
	return
}
