package server

import (
	"enriqueFcoG/cypher/internal/adapters/http"

	"github.com/gin-gonic/gin"
)

func RegisterBeatTypeRoutes(r *gin.RouterGroup, h *http.BeatTypeHandler) {
	beatTypes := r.Group("beat-types")
	{
		beatTypes.POST("", h.Create)
		beatTypes.GET("/:id", h.Get)
	}
}
