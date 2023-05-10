package web

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/uberate/gom/cmd/web/bc"
	"github.com/uberate/gom/internal/web/middleware"
	"net/http"
	"strings"
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

	engine.Use(
		middleware.GinRecoverLog(logger),
		middleware.GinLoggerMiddleWare(logger),
	)

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
