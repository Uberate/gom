package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GinRecoverLog(logger *logrus.Logger) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				logger.Panic(r)
			}
		}()
		ctx.Next()
	}
}
