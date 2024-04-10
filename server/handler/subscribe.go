package handler

import (
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
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return h.service.Subscribe(&req)
}
