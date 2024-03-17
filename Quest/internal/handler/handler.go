package handler

import (
	"Quest/internal/service"
	"Quest/internal/types"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	router.GET("/user/:id/history", h.GetUserHistory)
	router.POST("/quest", h.AddQuest)
	router.POST("/signal", h.ProcessSignal)

	return router
}

func (h *Handler) AddUser(c *gin.Context) {
	var user types.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.Response{"Can't parse user", err.Error()})
		return
	}

	err = h.service.AddUser(context.TODO(), user)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.Response{"Can't save user", err.Error()})
		return
	}

	c.JSON(http.StatusOK, types.Response{"User saved", user})
}

func (h *Handler) AddQuest(c *gin.Context) {
	var quest types.Quest
	err := c.BindJSON(&quest)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.Response{"Can't parse quest", err.Error()})
		return
	}

	err = h.service.AddQuest(context.TODO(), quest)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.Response{"Can't save quest", err.Error()})
		return
	}

	c.JSON(http.StatusOK, types.Response{"Quest saved", quest})
}

func (h *Handler) ProcessSignal(c *gin.Context) {
	var signal types.Signal
	err := c.BindJSON(&signal)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.Response{"Can't parse signal", err.Error()})
		return
	}

	user, err := h.service.ProcessSignal(context.TODO(), signal)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.Response{"Can't process signal", err.Error()})
		return
	}

	c.JSON(http.StatusOK, types.Response{"Signal processed", user})
}

func (h *Handler) GetUserHistory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, types.Response{"Can't parse user id", err.Error()})
		return
	}
	userHistory, err := h.service.GetUserHistory(context.TODO(), id)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.Response{"Can't get user History", err.Error()})
		return
	}

	c.JSON(http.StatusOK, types.Response{"Get User history", userHistory})
}
