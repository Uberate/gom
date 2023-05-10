package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RespStruct struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func Ok(ctx *gin.Context, res interface{}) {
	if res == nil {
		ctx.JSON(http.StatusNoContent, nil)
		return
	}

	resObj := RespStruct{
		Data: res,
	}

	ctx.JSON(http.StatusOK, resObj)
}

func Fail(ctx *gin.Context, err error) {
	resObj := RespStruct{
		Error: err.Error(),
	}

	ctx.JSON(http.StatusInternalServerError, resObj)
}

func Resp(ctx *gin.Context, code int, res interface{}) {
	ctx.JSON(code, res)
}
