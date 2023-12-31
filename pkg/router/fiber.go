package router

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"encoding/json"

	"github.com/disorn-inc/go-rest-ecom-th/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type FiberContext struct {
	*fiber.Ctx
}

func NewFiberContext(c *fiber.Ctx) *FiberContext {
	return &FiberContext{Ctx: c}
}

func (c *FiberContext) Bind(v interface{}) error {
	return c.Ctx.BodyParser(v)
}

func (c *FiberContext) BindQuery(v interface{}) error {
	return c.Ctx.QueryParser(v)
}

func (c *FiberContext) JSON(statusCode int, v interface{}) {
	c.Ctx.Status(statusCode).JSON(v)
}

func (c *FiberContext) Query(key string) string {
	return c.Ctx.Query(key)
}

func (c *FiberContext) Param(key string) string {
	return c.Ctx.Params(key)
}

func (c *FiberContext) GetHeader(key string) string {
	return c.Ctx.Get(key)
}

func (c *FiberContext) SetHeader(key, value string) {
	c.Ctx.Set(key, value)
}

func (c *FiberContext) Next() {
	c.Ctx.Next()
}

func (c *FiberContext) GetClientIP() string {
	return c.Ctx.IP()
}

func (c *FiberContext) GetRemoteIP() string {
	return c.Ctx.IP()
}

func (c *FiberContext) GetMethod() string {
	return c.Ctx.Method()
}

func (c *FiberContext) GetPath() string {
	return c.Ctx.Path()
}

func (c *FiberContext) GetHost() string {
	return c.Ctx.Hostname()
}

func (c *FiberContext) GetProtocol() string {
	return c.Ctx.Protocol()
}

func (c *FiberContext) GetStatus() int {
	return c.Ctx.Response().StatusCode()
}

func (c *FiberContext) GetBody() []byte {
	return c.Ctx.Body()
}

func (c *FiberContext) GetBodyString() string {
	return string(c.Ctx.Body())
}

// func (c *FiberContext) GetQuery() string {
// 	return c.Ctx.Q
// }

type FiberRouter struct {
	*fiber.App
}

func NewFiberRouter(cfg config.IConfig) *FiberRouter {
	r := fiber.New(fiber.Config{
		ReadTimeout:  cfg.App().ReadTimeout(),
		WriteTimeout: cfg.App().WriteTimeout(),
		BodyLimit:    cfg.App().BodyLimit(),
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
	})

	r.Use(logger.New())

	r.Use(cors.New(cors.Config{
		Next:             cors.ConfigDefault.Next,
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,PATCH,OPTIONS",
		AllowCredentials: false,
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, X-Requested-With, X-Access-Token, X-Refresh-Token, X-Trace-Id",
		ExposeHeaders:    "",
		MaxAge:           0,
	}))

	return &FiberRouter{r}
}

func NewFiberHandler(handler func(Context)) fiber.Handler {
	return func(c *fiber.Ctx) error {
		handler(NewFiberContext(c))
		return nil
	}
}

func NewFiberHandlerMiddleware(handler ...func(Context)) []fiber.Handler {
	var handlers []fiber.Handler

	for _, h := range handler {
		handlers = append(handlers, NewFiberHandler(h))
	}

	return handlers
}

func NewFiberMiddleware(handler func(Context)) fiber.Handler {
	return func(c *fiber.Ctx) error {
		handler(NewFiberContext(c))
		return c.Next()
	}
}

func (r *FiberRouter) POST(path string, handler func(Context)) {
	r.App.Post(path, NewFiberHandler(handler))
}

func (r *FiberRouter) GET(path string, handler func(Context)) {
	r.App.Get(path, NewFiberHandler(handler))
}

func (r *FiberRouter) PUT(path string, handler func(Context)) {
	r.App.Put(path, NewFiberHandler(handler))
}

func (r *FiberRouter) DELETE(path string, handler func(Context)) {
	r.App.Delete(path, NewFiberHandler(handler))
}

func (r *FiberRouter) PATCH(path string, handler func(Context)) {
	r.App.Patch(path, NewFiberHandler(handler))
}

func (r *FiberRouter) OPTIONS(path string, handler func(Context)) {
	r.App.Options(path, NewFiberHandler(handler))
}

func (r *FiberRouter) Use(middleware func(Context)) {
	r.App.Use(NewFiberHandler(middleware))
}

func (r *FiberRouter) Static(prefix, root string) {
	r.App.Static(prefix, root)
}

func (r *FiberRouter) Group(path string) *FiberGroup {
	return &FiberGroup{r.App.Group(path)}
}

func (r *FiberRouter) ListenAndServe() func() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	go func() {
		if err := r.App.Listen(":" + port); err != nil {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	return func() {
		tx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		defer stop()

		<-tx.Done()
		stop()
		fmt.Println("shutting down gracefully, press Ctrl+C again to force")

		tCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := r.App.ShutdownWithContext(tCtx); err != nil {
			slog.Error("server shutdown error", err)
		}
	}
}

type FiberGroup struct {
	fiber.Router
}

func (r *FiberGroup) POST(path string, handler func(Context)) {
	r.Router.Post(path, NewFiberHandler(handler))
}

func (r *FiberGroup) GET(path string, handler func(Context)) {
	r.Router.Get(path, NewFiberHandler(handler))
}

func (r *FiberGroup) PUT(path string, handler func(Context)) {
	r.Router.Put(path, NewFiberHandler(handler))
}

func (r *FiberGroup) DELETE(path string, handler func(Context)) {
	r.Router.Delete(path, NewFiberHandler(handler))
}

func (r *FiberGroup) PATCH(path string, handler func(Context)) {
	r.Router.Patch(path, NewFiberHandler(handler))
}

func (r *FiberGroup) OPTIONS(path string, handler func(Context)) {
	r.Router.Options(path, NewFiberHandler(handler))
}

func (r *FiberGroup) Use(middleware func(Context)) {
	r.Router.Use(NewFiberHandler(middleware))
}

func (r *FiberGroup) Static(prefix, root string) {
	r.Router.Static(prefix, root)
}

func (r *FiberGroup) Group(path string) *FiberGroup {
	return &FiberGroup{r.Router.Group(path)}
}
