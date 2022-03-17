package message

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/mars-terminal/chat-draw/internal/util"
)

// getByChatID godoc
// @Summary      Get chat by chat id.
// @Description  Gets id and if everything is OK gives chat by chat id.
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
func (h *Handler) getByChatID(c *gin.Context) {
	logger := log.WithField("message", "get by chat id")

	id, err := strconv.Atoi(c.Param("chat_id"))
	if err != nil {
		util.NewErrorResponse(logger, c.Writer, util.ParseErrorToHTTPErrorCode(err), err.Error())
		return
	}

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

	messages, err := h.Service.GetByChatID(c, int64(id), int64(limit), int64(offset))
	if err != nil {
		util.NewErrorResponse(logger, c.Writer, util.ParseErrorToHTTPErrorCode(err), err.Error())
		return
	}

	c.JSON(http.StatusOK, messages)
}

// getByPeerID godoc
// @Summary      Get chat by peer id.
// @Description  Gets id and if everything is OK gives chat by peer id.
// @Tags         messages
// @Accept       json
// @Produce      json
// @Param		 peer_id path int true "peer id"
// @Param		 limit query int false "limit query"
// @Param		 offset query int false "offset query"
// @Success      200  {array} 	message.Message
// @Failure      400  {object}  entities.Response
// @Failure      500  {object}  entities.Response
// @Security 	 ApiKeyAuth
// @Router       /message/peer_id/{peer_id} [get]
func (h *Handler) getByPeerID(c *gin.Context) {
	logger := log.WithField("message", "get by peer id")

	id, err := strconv.Atoi(c.Param("peer_id"))
	if err != nil {
		util.NewErrorResponse(logger, c.Writer, util.ParseErrorToHTTPErrorCode(err), err.Error())
		return
	}

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

	m, err := h.Service.GetByPeerID(c, int64(id), int64(limit), int64(offset))
	if err != nil {
		util.NewErrorResponse(logger, c.Writer, util.ParseErrorToHTTPErrorCode(err), err.Error())
		return
	}

	c.JSON(http.StatusOK, m)
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
