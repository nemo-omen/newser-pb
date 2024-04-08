package handler

import (
	"fmt"
	"net/http"

	"server/service"

	"github.com/labstack/echo/v5"
)

type SearchRequest struct {
	SearchUrl string `json:"searchurl"`
}

type SearchHandler struct {
	service *service.SearchService
}

func NewSearchHandler(service *service.SearchService) *SearchHandler {
	return &SearchHandler{
		service: service,
	}
}

func (h *SearchHandler) Search(c echo.Context) error {
	var req SearchRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	fmt.Println("searchUrl: ", req.SearchUrl)

	feedLinks, err := h.service.FindFeedLinks(req.SearchUrl)
	if err != nil {
		// fmt.Println("error: ", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
		// return c.JSON(500, err.Error())
	}

	feeds, err := h.service.GetFeeds(feedLinks)
	if err != nil {
		return c.JSON(500, err.Error())
	}
	return c.JSON(200, feeds)
}
