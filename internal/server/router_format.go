package server

import (
	"enriqueFcoG/cypher/internal/adapters/http"

	"github.com/gin-gonic/gin"
)

func RegisterFormat(r *gin.RouterGroup, h *http.FormatHandler) {
	formats := r.Group("format")
	{
		formats.POST("", h.Create)
		formats.GET("/:id", h.Get)
	}
}
