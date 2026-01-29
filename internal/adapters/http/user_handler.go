package http

import (
	"enriqueFcoG/cypher/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service *usecase.UserService
}

func NewUserHandler(service *usecase.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (h *UserHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var req struct {
		ID    string `json:"id"`
		Email string `json:"string"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := h.Service.CreateUser(ctx, req.ID, req.Email)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func (h *UserHandler) Get(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	user, err := h.Service.GetUser(ctx, id)
	if err != nil {
		c.JSON(404, gin.H{"error": "not found"})
		return
	}
	c.JSON(200, user)
}
