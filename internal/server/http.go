package server

import (
	"enriqueFcoG/cypher/internal/adapters/http"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Beat          *http.BeatHandler
	BeatTypes     *http.BeatTypeHandler
	Crew          *http.CrewHandler
	FormatDetails *http.FormatDetailsHandler
	Format        *http.FormatHandler
	History       *http.HistoryHandler
	Mod           *http.ModHandler
	User          *http.UserHandler
}

func NewHTTPServer(handler *Handlers) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery(), gin.Logger())

	v1 := r.Group("/api/v2")

	RegisterBeatRoutes(v1, handler.Beat)
	RegisterBeatTypeRoutes(v1, handler.BeatTypes)
	RegisterCrewRoutes(v1, handler.Crew)
	RegisterFormatDetailsRoutes(v1, handler.FormatDetails)
	RegisterFormat(v1, handler.Format)
	RegisterHistoryRoutes(v1, handler.History)
	RegisterModRoutes(v1, handler.Mod)
	RegisterUserRoutes(v1, handler.User)

	return r
}
