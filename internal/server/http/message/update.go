package message

import (
	"net/http"
	"repositorie/internal/entities/message"
	"repositorie/internal/util"

	"github.com/gin-gonic/gin"
)

// update godoc
// @Summary      Updating message.
// @Description  Gets body if everything OK gives back updated message.
// @Tags         messages
// @Accept       json
// @Produce      json
// @Success      200  {array} 	message.Message
// @Failure      400  {object}  entities.Response
// @Failure      500  {object}  entities.Response
// @Security 	ApiKeyAuth
// @Router       /message/update [put]
func (h *Handler) update(c *gin.Context) {
	logger := log.WithField("message", "update")

	var body message.UpdateMessageQuery

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
