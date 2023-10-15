package monitor

import (
	"log/slog"
	"net/http"

	"github.com/disorn-inc/go-rest-ecom-th/config"
	"github.com/disorn-inc/go-rest-ecom-th/pkg/entities"
	"github.com/disorn-inc/go-rest-ecom-th/pkg/router"
)

type MonitorController struct {
	cfg config.IConfig
}

func NewMonitorController(cfg config.IConfig) *MonitorController {
	return &MonitorController{
		cfg: cfg,
	}
}

func (h *MonitorController) HealthCheck(c router.Context) {
	slog.Info("Health check")
	entities.NewResponse(c).Succeed(http.StatusOK, &Monitor{
		Name: h.cfg.App().Name(),
		Version: h.cfg.App().Version(),
	}).Response()
}