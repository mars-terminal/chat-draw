package user

import (
	"net/http"
	"repositorie/internal/entities"
	"repositorie/internal/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

// delete godoc
// @Summary      Deleting user.
// @Description  Gets user id if everything OK deleting user.
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  auth.Tokens
// @Failure      400  {object}  entities.Response
// @Failure      500  {object}  entities.Response
// @Security 	ApiKeyAuth
// @Router       /users/delete [delete]
func (h *Handler) delete(c *gin.Context) {
	logger := log.WithField("user", "delete")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		util.NewErrorResponse(logger, c.Writer, util.ParseErrorToHTTPErrorCode(err), err.Error())
		return
	}

	err = h.Service.DeleteByID(c, int64(id))
	if err != nil {
		util.NewErrorResponse(logger, c.Writer, util.ParseErrorToHTTPErrorCode(err), err.Error())
		return
	}

	c.JSON(http.StatusOK, entities.Response{
		Status:  http.StatusOK,
		Message: "ok",
	})
}
