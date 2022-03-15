package auth

import (
	"net/http"

	"github.com/mars-terminal/chat-draw/internal/entities/auth"
	"github.com/mars-terminal/chat-draw/internal/util"

	"github.com/gin-gonic/gin"
)

// signIn godoc
// @Summary      Signs In the user.
// @Description  Gets body and Authenticate user if exists.
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        SignUpQuery body auth.SignInQuery false "Request payload"
// @Success      200  {object}  auth.Tokens
// @Failure      400  {object}  entities.Response
// @Failure      500  {object}  entities.Response
// @Router       /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	logger := log.WithField("handler", "signIn")

	var (
		body auth.SignInQuery
	)

	if err := c.BindJSON(&body); err != nil {
		util.NewErrorResponse(logger.WithError(err), c.Writer, http.StatusBadRequest, err.Error())
		return
	}

	if err := util.Validate.Struct(&body); err != nil {
		util.NewErrorResponse(logger.WithError(err), c.Writer, http.StatusBadRequest, err.Error())
		return
	}

	tokens, err := h.Service.SignIn(c, &body)

	if err != nil {
		util.NewErrorResponse(logger.WithError(err), c.Writer, util.ParseErrorToHTTPErrorCode(err), err.Error())
		return
	}

	c.JSON(http.StatusOK, tokens)
}
