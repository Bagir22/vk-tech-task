package handler

import (
	"Quest/internal/service"
	"Quest/internal/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	service service.Repository
}

func InitHandler(service service.Repository) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()
	router.POST("/user", h.AddUser)
	//router.GET("/signal/:id/history", h.GetUserHistory)
	//router.POST("/quest", h.AddQuest)
	//router.POST("/signal", h.AddSignal)

	return router
}

func (h *Handler) AddUser(c *gin.Context) {
	var user types.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.Response{"Can't parse user", err.Error()})
		return
	}

	err = h.service.AddUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.Response{"Can't save user", err.Error()})
		return
	}

	c.JSON(http.StatusOK, types.Response{"User saved", user})
}
