package logger

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/disorn-inc/go-rest-ecom-th/pkg/router"
	"github.com/disorn-inc/go-rest-ecom-th/pkg/utils"
)

type ILogger interface {
	Print() ILogger
	Save()
	SetQuery(ctx router.Context)
	SetBody(ctx router.Context)
	SetResponse(response any)
}

type logger struct {
	Time       string `json:"time"`
	Ip         string `json:"ip"`
	Method     string `json:"method"`
	StatusCode int    `json:"statusCode"`
	Path       string `json:"path"`
	Query      any    `json:"query"`
	Body       any    `json:"body"`
	Response   any    `json:"response"`
}

func InitLogger(ctx router.Context, res any) ILogger {
	log := &logger{
		Time:       time.Now().Local().Format("2006-01-02 15:04:05"),
		Ip:         ctx.GetClientIP(),
		Method:     ctx.GetMethod(),
		Path:       ctx.GetPath(),
		StatusCode: ctx.GetStatus(),
	}

	log.SetQuery(ctx)
	log.SetBody(ctx)
	log.SetResponse(res)
	return log
}

func (l *logger) Print() ILogger {
	utils.Debug(l)
	return l
}

func (l *logger) Save() {
	data := utils.Output(l)

	fileName := fmt.Sprintf("./assets/logs/ecom_logger_%v.txt", strings.ReplaceAll(time.Now().Local().Format("2006-01-02"), "-", ""))
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	defer file.Close()

	file.WriteString(string(data) + "\n")
}

func (l *logger) SetQuery(ctx router.Context) {
	var query any
	if err := ctx.BindQuery(&query); err != nil {
		slog.Error("Error logger set query", "error:", err)
	}

	l.Query = query
}

func (l *logger) SetBody(ctx router.Context) {
	var body any
	if err := ctx.Bind(&body); err != nil {
		slog.Error("Error logger set body", "error:", err)
	}

	switch l.Path {
	case "/api/v1/user/login":
		l.Body = "secret"
	default:
		l.Body = body
	}
}

func (l *logger) SetResponse(response any) {
	l.Response = response
}
