package message

import (
	"net/http"
	"strconv"

	"github.com/mars-terminal/chat-draw/internal/entities"
	"github.com/mars-terminal/chat-draw/internal/util"

	"github.com/gin-gonic/gin"
)

// delete godoc
// @Summary      Deleting message.
// @Description  Gets message id if everything is OK deleting message.
// @Tags         messages
// @Accept       json
// @Produce      json
// @Param		 messageId path int true "message id"
// @Success      200  {array} 	message.Message
// @Failure      400  {object}  entities.Response
// @Failure      500  {object}  entities.Response
// @Security 	 ApiKeyAuth
// @Router       /message/delete/{id} [delete]
func (h *Handler) delete(c *gin.Context) {
	logger := log.WithField("message", "delete")

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
