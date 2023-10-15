package api

import (
	"github.com/disorn-inc/go-rest-ecom-th/internal/domains/monitor"
	"github.com/disorn-inc/go-rest-ecom-th/pkg/router"
)

type Routers interface {
	router.Router
}

type routers struct {
	MonitorRouter monitor.Router
}

func NewRouters(
	monitorRouter monitor.Router,
) Routers {
	return &routers{
		MonitorRouter: monitorRouter,
	}
}

func (r *routers) Initial(app *router.FiberRouter) {
	r.MonitorRouter.Initial(app)
}