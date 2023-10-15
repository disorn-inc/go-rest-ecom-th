package monitor

import (
	"github.com/disorn-inc/go-rest-ecom-th/config"
	"github.com/disorn-inc/go-rest-ecom-th/internal/domains/middleware"
	"github.com/disorn-inc/go-rest-ecom-th/pkg/router"
)

type Router interface {
	router.Router
}

type monitorRouter struct {
	Cfg config.IConfig
	Middleware *middleware.Middleware
}

func (r *monitorRouter) Initial(app *router.FiberRouter) {
	controller := NewMonitorController(r.Cfg)
	app.GET("/health-check", controller.HealthCheck)
}

func NewMonitorRouter(
	cfg config.IConfig,
	middleware *middleware.Middleware,
) Router {
	return &monitorRouter{
		Cfg: cfg,
		Middleware: middleware,
	}
}