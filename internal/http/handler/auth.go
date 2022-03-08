package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"repositorie/internal/entities/auth"
	"repositorie/internal/util"
)

func (h *Handler) signUp(c *gin.Context) {
	var (
		body auth.SignUpQuery
	)

	if err := c.BindJSON(&body); err != nil {
		return // TODO: implement error fallback
	}

	if err := util.Validate.Struct(&body); err != nil {
		log.WithError(err)
		util.NewErrorResponse(log, c, http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) signIn(c *gin.Context) {

}
