package server

import (
	"enriqueFcoG/cypher/internal/adapters/http"

	"github.com/gin-gonic/gin"
)

func RegisterHistoryRoutes(r *gin.RouterGroup, h *http.HistoryHandler) {
	histories := r.Group("histories")
	{
		histories.POST("", h.Create)
		histories.GET("/:id", h.Get)
	}
}
