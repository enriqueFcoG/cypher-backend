package server

import (
	"enriqueFcoG/cypher/internal/adapters/http"

	"github.com/gin-gonic/gin"
)

func RegisterFormatDetailsRoutes(r *gin.RouterGroup, h *http.FormatDetailsHandler) {
	format_details := r.Group("format-details")
	{
		format_details.POST("", h.Create)
		format_details.GET("/:id", h.Get)
	}
}
