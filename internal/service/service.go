package service

import (
	"repositorie/internal/entities"
	"repositorie/internal/entities/user"

	"github.com/gin-gonic/gin"
)

func GetUserFromContext(ctx *gin.Context) (*user.User, error) {
	u, exists := ctx.Get(gin.AuthUserKey)
	if !exists {
		return nil, entities.ErrUnauthorized
	}

	if u, ok := u.(*user.User); ok {
		return u, nil
	}

	return nil, entities.ErrUnauthorized
}
