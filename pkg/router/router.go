package router

type Router interface {
	Initial(app *FiberRouter)
}