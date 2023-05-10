package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/uberate/gom/cmd/web/bc"
	"net/http"
	"strings"
	"time"
)

func RunE(config bc.ApplicationConfig, logger *logrus.Logger, v map[string]string) error {
	enableWebLog := config.Web.EnableWebLog && (logger != nil)
	if enableWebLog {
		logger.Info("init web engine")
	}
	gin.SetMode(config.Web.Mode)
	ginEngine := gin.New()

	if err := routes(ginEngine, config, logger, v); err != nil {
		return err
	}

	return ginEngine.Run(strings.Join([]string{config.Web.Host, config.Web.Port}, ":"))
}

func routes(engine *gin.Engine, config bc.ApplicationConfig, logger *logrus.Logger, v map[string]string) error {

	engine.Use(ginLoggerMiddleWare(logger))

	engine.Any("/hello", versionHandler(v))
	engine.Any("/version", versionHandler(v))
	return nil
}

// versionHandler will output the version of application use at /hello and /version route(any method of http).
func versionHandler(v map[string]string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, v)
	}
}

func ginLoggerMiddleWare(logger *logrus.Logger) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()

		Latency := time.Now().Sub(start)
		ClientIP := ctx.ClientIP()
		Method := ctx.Request.Method
		StatusCode := ctx.Writer.Status()

		value := fmt.Sprintf("Request %-4s| Code: %3d | Latency: %5dms | From: %15s | Path: %s",
			Method, StatusCode, Latency.Milliseconds(), ClientIP, ctx.Request.URL.Path)

		if StatusCode <= 399 {
			logger.Infof(value)
		} else if StatusCode >= 400 && StatusCode <= 499 {
			logger.Warn(value)
		} else {
			logger.Error(value)
		}
	}
}
