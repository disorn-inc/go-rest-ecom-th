package middleware

type Middleware struct {
	middlewareService *MiddlewareService
}

func NewMiddleware(middlewareService *MiddlewareService) *Middleware {
	return &Middleware{
		middlewareService: middlewareService,
	}
}