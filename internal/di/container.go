package di

import (
	"context"
	httpadapter "enriqueFcoG/cypher/internal/adapters/http"
	"enriqueFcoG/cypher/internal/adapters/persistence"
	"enriqueFcoG/cypher/internal/config"
	"enriqueFcoG/cypher/internal/server"
	"enriqueFcoG/cypher/internal/usecase"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/gin-gonic/gin"
)

type Container struct {
	BeatHandler   *httpadapter.BeatHandler
	Beat          *httpadapter.BeatTypeHandler
	Crew          *httpadapter.CrewHandler
	FormatDetails *httpadapter.FormatDetailsHandler
	Format        *httpadapter.FormatHandler
	History       *httpadapter.HistoryHandler
	Mod           *httpadapter.ModHandler
	User          *httpadapter.UserHandler
}

func BuildServer(cfg config.Config) *gin.Engine {
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, cfg.Database)
	fmt.Println(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to create the DB pool")
	}

	//db := connectPostgres(cfg.Database)

	// repos
	beatRepo := persistence.NewBeatRepoPG(pool)
	beatTypeRepo := persistence.NewBeatTypeRepoPG(pool)
	crewRepo := persistence.NewCrewRepoPG(pool)
	formatDetailsRepo := persistence.NewFormatDetailsRepoPG(pool)
	formatRepo := persistence.NewFormatRepoPG(pool)
	historyRepo := persistence.NewHistoryRepoPG(pool)
	modRepo := persistence.NewModRepoPG(pool)
	userRepo := persistence.NewUserRepoPG(pool)

	// services / usecases
	beatService := &usecase.BeatService{
		Repo: beatRepo,
	}
	beatTypeService := &usecase.BeatTypeService{
		Repo: beatTypeRepo,
	}
	crewService := &usecase.CrewService{
		Repo: crewRepo,
	}
	formatDetailsService := &usecase.FormatDetailsService{
		Repo: formatDetailsRepo,
	}
	formatService := &usecase.FormatService{
		Repo: formatRepo,
	}
	historyService := &usecase.HistoryService{
		Repo: historyRepo,
	}
	modService := &usecase.ModService{
		Repo: modRepo,
	}
	userService := &usecase.UserService{
		Repo: userRepo,
	}

	// //handlers
	beatHandler := &httpadapter.BeatHandler{
		Service: beatService,
	}
	beatTypeHandler := &httpadapter.BeatTypeHandler{
		Service: beatTypeService,
	}
	crewHandler := &httpadapter.CrewHandler{
		Service: crewService,
	}
	formatDetailsHandler := &httpadapter.FormatDetailsHandler{
		Service: formatDetailsService,
	}
	formatHandler := &httpadapter.FormatHandler{
		Service: formatService,
	}
	historyHandler := &httpadapter.HistoryHandler{
		Service: historyService,
	}
	modHandler := &httpadapter.ModHandler{
		Service: modService,
	}
	userHandler := &httpadapter.UserHandler{
		Service: userService,
	}

	//router
	handlers := &server.Handlers{
		Beat:          beatHandler,
		BeatTypes:     beatTypeHandler,
		Crew:          crewHandler,
		FormatDetails: formatDetailsHandler,
		Format:        formatHandler,
		History:       historyHandler,
		Mod:           modHandler,
		User:          userHandler,
	}

	return server.NewHTTPServer(cfg, handlers)

}
