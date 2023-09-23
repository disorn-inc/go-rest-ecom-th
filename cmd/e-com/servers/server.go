package servers

import (
	// "context"
	// "log"
	// "os/signal"
	// "syscall"

	"github.com/disorn-inc/go-rest-ecom-th/config"
	"github.com/disorn-inc/go-rest-ecom-th/internal/api"
	"github.com/disorn-inc/go-rest-ecom-th/pkg/databases"
	"github.com/disorn-inc/go-rest-ecom-th/pkg/router"

	// "github.com/disorn-inc/go-rest-ecom-th/pkg/router"
	"github.com/jmoiron/sqlx"
)

type IServer interface {
	Start()
}

type server struct {
	app *router.FiberRouter
	db  *sqlx.DB
	cfg config.IConfig
}

func NewServer(cfg config.IConfig, db *sqlx.DB) IServer {
	return &server{
		app: router.NewFiberRouter(cfg),
		db:  db,
		cfg: cfg,
	}
}

func (s *server) Start() {
	// c := make(chan os.Signal, 1)
	// ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	// defer stop()

	// go func() {
	// 	<-ctx.Done()
	// 	log.Println("Shutting down server...")
	// 	stop()

	// 	if err := s.app.Shutdown(); err != nil {
	// 		log.Fatalf("Error shutting down server: %v", err)
	// 	}

	// 	log.Println("Server gracefully stopped")
	// }()

	// s.app.Listen(s.cfg.App().Url())


	dbDriver := databases.NewDrivers(s.db)
	apis := api.CreateApi(dbDriver)
	apis.InitialRouter(s.app)
	s.app.ListenAndServe()()
	
}
