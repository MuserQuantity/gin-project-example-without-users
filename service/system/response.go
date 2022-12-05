package system

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Result interface{} `json:"result"`
}

const (
	ERROR   = 1
	SUCCESS = 0
)

func Result(c *gin.Context, code int, data interface{}, msg string) {
	c.JSON(http.StatusOK, Response{
		code,
		msg,
		data,
	})
}

func Succeed(c *gin.Context) {
	Result(c, SUCCESS, map[string]interface{}{}, "success")
}

func SucceedWithMessage(c *gin.Context, message string) {
	Result(c, SUCCESS, map[string]interface{}{}, message)
}

func SucceedWithData(c *gin.Context, data interface{}) {
	Result(c, SUCCESS, data, "success")
}

func SucceedWithDetailed(c *gin.Context, data interface{}, message string) {
	Result(c, SUCCESS, data, message)
}

func Fail(c *gin.Context) {
	Result(c, ERROR, map[string]interface{}{}, "error")
}

func FailWithErrorArgs(c *gin.Context) {
	Result(c, ERROR, map[string]interface{}{}, "args error")
}

func FailWithMessage(c *gin.Context, message string) {
	Result(c, ERROR, map[string]interface{}{}, message)
}

func FailWithDetailed(c *gin.Context, data interface{}, message string) {
	Result(c, ERROR, data, message)
}
