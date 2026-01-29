package http

import (
	"enriqueFcoG/cypher/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CrewHandler struct {
	Service *usecase.CrewService
}

func NewCrewHandler(service *usecase.CrewService) *CrewHandler {
	return &CrewHandler{Service: service}
}

func (h *CrewHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var req struct {
		Title string `json:"string"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := h.Service.CreateCrew(ctx, req.Title)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func (h *CrewHandler) Get(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	user, err := h.Service.GetCrew(ctx, id)
	if err != nil {
		c.JSON(404, gin.H{"error": "not found"})
		return
	}
	c.JSON(200, user)
}
