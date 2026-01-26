package server

import (
	"enriqueFcoG/cypher/internal/adapters/http"

	"github.com/gin-gonic/gin"
)

func RegisterCrewRoutes(r *gin.RouterGroup, h *http.CrewHandler) {
	crews := r.Group("crews")
	{
		crews.POST("", h.Create)
		crews.GET("/:id", h.Get)
	}
}
