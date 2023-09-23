package api

import "github.com/disorn-inc/go-rest-ecom-th/pkg/router"

type Routers interface {
	router.Router
}

type routers struct {}

func NewRouters() Routers {
	return &routers{}
}

func (r *routers) Initial(app *router.FiberRouter) {
	
}