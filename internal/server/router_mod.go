package server

import (
	"enriqueFcoG/cypher/internal/adapters/http"

	"github.com/gin-gonic/gin"
)

func RegisterModRoutes(r *gin.RouterGroup, h *http.ModHandler) {
	mods := r.Group("mods")
	{
		mods.POST("", h.Create)
		mods.GET("/:id", h.Get)
	}
}
