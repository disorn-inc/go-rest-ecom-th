package middleware

type IMiddlewareService interface {

}

type MiddlewareService struct {
	middlewareRepository IMiddlewareRepository
}

func NewMiddlewareService(middlewareRepository IMiddlewareRepository) *MiddlewareService {
	return &MiddlewareService{
		middlewareRepository: middlewareRepository,
	}
}