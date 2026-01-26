package server

import (
	"enriqueFcoG/cypher/internal/adapters/http"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.RouterGroup, h *http.UserHandler) {
	users := r.Group("users")
	{
		users.POST("", h.Create)
		users.GET("/:id", h.Get)
	}
}
