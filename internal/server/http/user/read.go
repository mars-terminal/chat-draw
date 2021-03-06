package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/mars-terminal/chat-draw/internal/entities/user"
	"github.com/mars-terminal/chat-draw/internal/service"
	"github.com/mars-terminal/chat-draw/internal/util"
)

// me godoc
// @Summary      Showing users profile.
// @Description  Gets user if everything OK gives back user with 4 fields.
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  auth.Tokens
// @Failure      400  {object}  entities.Response
// @Failure      500  {object}  entities.Response
// @Security 	 ApiKeyAuth
// @Router       /users/me [get]
func (h *Handler) me(c *gin.Context) {
	logger := log.WithField("user", "me")

	u, err := service.GetUserFromContext(c)

	if err != nil {
		util.NewErrorResponse(logger, c.Writer, util.ParseErrorToHTTPErrorCode(err), err.Error())
		return
	}

	c.JSON(http.StatusOK, &user.User{
		ID:         u.ID,
		FirstName:  u.FirstName,
		SecondName: u.SecondName,
		NickName:   u.NickName,
		Phone:      u.Phone,
	})
}

// search godoc
// @Summary      Searching by Nick Name of ID.
// @Description  Gets user ID or Nick Name if everything OK gives back user.
// @Tags         users
// @Accept       json
// @Produce      json
// @Param		 Query path string true "query nickname or id"
// @Success      200  {object}  auth.Tokens
// @Failure      400  {object}  entities.Response
// @Failure      500  {object}  entities.Response
// @Security 	 ApiKeyAuth
// @Router       /users/search/{query} [get]
func (h *Handler) search(c *gin.Context) {
	logger := log.WithField("user", "search")

	var (
		isID = false

		query = c.Param("query")
	)

	id, err := strconv.Atoi(query)
	if err == nil {
		isID = true
	}

	var users []*user.User

	switch {
	case isID:
		u, err := h.Service.GetByID(c, int64(id))
		if err != nil {
			util.NewErrorResponse(logger, c.Writer, util.ParseErrorToHTTPErrorCode(err), err.Error())
			return
		}
		users = append(users, u)
	default:
		users, err = h.Service.GetByNickName(c, query)
		if err != nil {
			util.NewErrorResponse(logger, c.Writer, util.ParseErrorToHTTPErrorCode(err), err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, users)
}
