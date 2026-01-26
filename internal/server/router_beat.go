package server

import (
	"enriqueFcoG/cypher/internal/adapters/http"

	"github.com/gin-gonic/gin"
)

func RegisterBeatRoutes(r *gin.RouterGroup, h *http.BeatHandler) {
	beats := r.Group("beats")
	{
		beats.POST("", h.Create)
		beats.GET("/:id", h.Get)
	}
}
