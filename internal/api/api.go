package api

import (
	"github.com/disorn-inc/go-rest-ecom-th/config"
	"github.com/disorn-inc/go-rest-ecom-th/internal/domains/middleware"
	"github.com/disorn-inc/go-rest-ecom-th/internal/domains/monitor"
	"github.com/disorn-inc/go-rest-ecom-th/pkg/databases"
	"github.com/disorn-inc/go-rest-ecom-th/pkg/router"
)

type API interface {
	Register(cfg config.IConfig)
	InitialRouter(r *router.FiberRouter)
}

type api struct {
	Router Routers
	Middleware *middleware.Middleware
}

func NewAPI(
	router Routers,
	middleWare *middleware.Middleware,
	) API {
	return &api{
		Router: router,
		Middleware: middleWare,
	}
}

func (a *api) Register(cfg config.IConfig) {
	r := router.NewFiberRouter(cfg)
	a.Router.Initial(r)
	r.ListenAndServe()()
}

func (a *api) InitialRouter(r *router.FiberRouter) {
	a.Router.Initial(r)
}


func CreateApi(dbDriver databases.Driver, cfg config.IConfig) API {
	middlewareService := middleware.NewMiddlewareService(middleware.NewMiddlewareRepository(dbDriver.GetPostgres()))
	middleware := middleware.NewMiddleware(middlewareService)

	monitorRouter := monitor.NewMonitorRouter(cfg, middleware)

	apiRouter := NewRouters(
		monitorRouter,
	)
	return NewAPI(
		apiRouter,
		middleware,
	)
}