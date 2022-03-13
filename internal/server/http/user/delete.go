package user

import (
	"net/http"
	"repositorie/internal/entities"
	"repositorie/internal/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
