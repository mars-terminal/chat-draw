package auth

import (
	"net/http"

	"github.com/mars-terminal/chat-draw/internal/entities/auth"
	"github.com/mars-terminal/chat-draw/internal/util"

	"github.com/gin-gonic/gin"
)

// signUp godoc
// @Summary      Signs Up new user.
// @Description  Gets body and create user if not exists.
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        SignUpQuery body auth.SignUpQuery true "Request payload"
// @Success      200  {object}  auth.Tokens
// @Failure      400  {object}  entities.Response
// @Failure      500  {object}  entities.Response
// @Router       /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	logger := log.WithField("handler", "signUp")

	var (
		body auth.SignUpQuery
	)

	if err := c.BindJSON(&body); err != nil {
		util.NewErrorResponse(logger.WithError(err), c.Writer, http.StatusBadRequest, err.Error())
		return
	}

	if err := util.Validate.Struct(&body); err != nil {
		util.NewErrorResponse(logger.WithError(err), c.Writer, http.StatusBadRequest, err.Error())
		return
	}

	tokens, err := h.Service.SignUp(c, &body)

	if err != nil {
		util.NewErrorResponse(logger.WithError(err), c.Writer, util.ParseErrorToHTTPErrorCode(err), err.Error())
		return
	}

	c.JSON(http.StatusOK, tokens)
}
