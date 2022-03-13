package message

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"repositorie/internal/entities/message"
	"repositorie/internal/util"
)

// create godoc
// @Summary      Create message.
// @Description  Gets body and  if everything is OK creating message.
// @Tags         messages
// @Accept       json
// @Produce      json
// @Param		 CreateMessageQuery body message.CreateMessageQuery true "Request payload"
// @Success      200  {array} 	message.Message
// @Failure      400  {object}  entities.Response
// @Failure      500  {object}  entities.Response
// @Security 	ApiKeyAuth
// @Router       /message/create [post]
func (h *Handler) create(c *gin.Context) {
	logger := log.WithField("message", "create")

	var body message.CreateMessageQuery

	if err := c.BindJSON(&body); err != nil {
		util.NewErrorResponse(logger, c.Writer, util.ParseErrorToHTTPErrorCode(err), err.Error())
		return
	}

	if err := util.Validate.Struct(body); err != nil {
		util.NewErrorResponse(logger, c.Writer, util.ParseErrorToHTTPErrorCode(err), err.Error())
		return
	}

	m, err := h.Service.Create(c, &body)
	if err != nil {
		util.NewErrorResponse(logger, c.Writer, util.ParseErrorToHTTPErrorCode(err), err.Error())
		return
	}

	c.JSON(http.StatusOK, m)
}
