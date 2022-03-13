package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"repositorie/internal/entities/user"
	"repositorie/internal/util"
)

// update godoc
// @Summary      Updating user fields.
// @Description  Gets user if everything OK gives back updated user.
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  auth.Tokens
// @Failure      400  {object}  entities.ErrorResponse
// @Failure      500  {object}  entities.ErrorResponse
// @Router       /users/update [put]
func (h *Handler) update(c *gin.Context) {
	logger := log.WithField("user", "update")

	var body user.UpdateUserQuery

	if err := c.BindJSON(&body); err != nil {
		util.NewErrorResponse(logger, c.Writer, util.ParseErrorToHTTPErrorCode(err), err.Error())
		return
	}

	if err := util.Validate.Struct(body); err != nil {
		util.NewErrorResponse(logger, c.Writer, util.ParseErrorToHTTPErrorCode(err), err.Error())
		return
	}

	u, err := h.Service.Update(c, &body)
	if err != nil {
		util.NewErrorResponse(logger, c.Writer, util.ParseErrorToHTTPErrorCode(err), err.Error())
		return
	}

	c.JSON(http.StatusOK, u)
}
