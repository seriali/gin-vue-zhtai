package response

import (
	"gin-vue-zhtai-server/utils/message"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Status     int         `json:"status"`
	Data       interface{} `json:"data"`
	StatusText string      `json:"statusText"`
}

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(message.SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(code int, c *gin.Context) {
	Result(code, map[string]interface{}{}, message.GetMsg(code), c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(message.SUCCESS, data, "操作成功", c)
}

func OkWithDetailed(code int, data interface{}, c *gin.Context) {
	Result(code, data, message.GetMsg(code), c)
}

func Fail(c *gin.Context) {
	Result(message.ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(code int, c *gin.Context) {
	Result(code, map[string]interface{}{}, message.GetMsg(code), c)
}

func FailWithDetailed(data interface{}, msg string, c *gin.Context) {
	Result(message.ERROR, data, msg, c)
}
