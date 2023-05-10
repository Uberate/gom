package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

func GinLoggerMiddleWare(logger *logrus.Logger) func(ctx *gin.Context) {
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
