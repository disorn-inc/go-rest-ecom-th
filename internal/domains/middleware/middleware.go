package middleware

import (
	"net/http"

	"github.com/disorn-inc/go-rest-ecom-th/pkg/entities"
	"github.com/disorn-inc/go-rest-ecom-th/pkg/router"
)

type middlewareHandlerErrorCode string

const (
	routerCheckErr middlewareHandlerErrorCode = "middleware-001"
)

type Middleware struct {
	middlewareService *MiddlewareService
}

func NewMiddleware(middlewareService *MiddlewareService) *Middleware {
	return &Middleware{
		middlewareService: middlewareService,
	}
}

func (h *Middleware) Cors(c router.Context) {

}

func (h *Middleware) RouteCheck(c router.Context) {
	entities.NewResponse(c).Error(
		http.StatusNotFound,
		string(routerCheckErr),
		"Router not found",
	).Response()
	c.Next()
}