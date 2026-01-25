package http

import (
	"enriqueFcoG/cypher/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HistoryHandler struct {
	Service *usecase.HistoryService
}

func (h *HistoryHandler) Create(c *gin.Context) {
	var req struct {
		Title string `json:"string"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := h.Service.CreateHistory(req.Title)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func (h *HistoryHandler) Get(c *gin.Context) {
	id := c.Param("id")
	user, err := h.Service.GetHistory(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "not found"})
		return
	}
	c.JSON(200, user)
}
