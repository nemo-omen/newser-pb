package handler

import (
	"fmt"
	"net/http"
	"server/service"
	"server/types"

	"github.com/labstack/echo/v5"
)

type SubscribeHandler struct {
	service *service.SubscribeService
}

func NewSubscribeHandler(service *service.SubscribeService) *SubscribeHandler {
	return &SubscribeHandler{service: service}
}

func (h *SubscribeHandler) Subscribe(c echo.Context) error {
	var req types.SubscribeRequest
	err := c.Bind(&req)
	fmt.Println("req: ", req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	sub, err := h.service.Subscribe(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, sub)
}
