package http

import (
	"enriqueFcoG/cypher/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HistoryHandler struct {
	Service *usecase.HistoryService
}

func NewHistoryHandler(service *usecase.HistoryService) *HistoryHandler {
	return &HistoryHandler{Service: service}
}

func (h *HistoryHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var req struct {
		Title string `json:"string"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := h.Service.CreateHistory(ctx, req.Title)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func (h *HistoryHandler) Get(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	user, err := h.Service.GetHistory(ctx, id)
	if err != nil {
		c.JSON(404, gin.H{"error": "not found"})
		return
	}
	c.JSON(200, user)
}
