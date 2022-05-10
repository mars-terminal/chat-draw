package message

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/mars-terminal/chat-draw/internal/util"
)

// getByChatIDAndPeerID godoc
// @Summary      Get chat by chat id.
// @Description  Gets chat and peer id and if everything is OK gives chat by chat id.
// @Tags         messages
// @Accept       json
// @Produce      json
// @Param		 chat_id path int true "chat id"
// @Param		 limit query int false "limit query"
// @Param		 offset query int false "offset query"
// @Success      200  {array} 	message.Message
// @Failure      400  {object}  entities.Response
// @Failure      500  {object}  entities.Response
// @Security 	 ApiKeyAuth
// @Router       /message/chat_id/{chat_id} [get]
func (h *Handler) getByChatIDAndPeerID(c *gin.Context) {
	logger := log.WithField("message", "get by chat id")

	chatId, err := strconv.Atoi(c.Param("chat_id"))
	if err != nil {
		util.NewErrorResponse(logger, c.Writer, util.ParseErrorToHTTPErrorCode(err), err.Error())
		return
	}

	peerId, err := strconv.Atoi(c.Param("peer_id"))
	if err != nil {
		util.NewErrorResponse(logger, c.Writer, util.ParseErrorToHTTPErrorCode(err), err.Error())
		return
	}

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 50
	}

	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		offset = 0
	}

	messages, err := h.Service.GetByChatIDAndPeerID(c, int64(chatId), int64(peerId), int64(limit), int64(offset))
	if err != nil {
		util.NewErrorResponse(logger, c.Writer, util.ParseErrorToHTTPErrorCode(err), err.Error())
		return
	}

	c.JSON(http.StatusOK, messages)
}

// search godoc
// @Summary      Search message.
// @Description  Gets query and if everything is OK searching message.
// @Tags         messages
// @Accept       json
// @Produce      json
// @Param		 query path string true "Search query"
// @Param		 limit query int false "limit query"
// @Param		 offset query int false "offset query"
// @Success      200  {array} 	message.Message
// @Failure      400  {object}  entities.Response
// @Failure      500  {object}  entities.Response
// @Security 	 ApiKeyAuth
// @Router       /message/search/{query} [get]
func (h *Handler) search(c *gin.Context) {
	logger := log.WithField("message", "search")

	query := c.Param("query")

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		util.NewErrorResponse(logger, c.Writer, util.ParseErrorToHTTPErrorCode(err), err.Error())
		return
	}

	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		util.NewErrorResponse(logger, c.Writer, util.ParseErrorToHTTPErrorCode(err), err.Error())
		return
	}

	search, err := h.Service.Search(c, query, int64(limit), int64(offset))
	if err != nil {
		util.NewErrorResponse(logger, c.Writer, util.ParseErrorToHTTPErrorCode(err), err.Error())
		return
	}

	c.JSON(http.StatusOK, search)
}
