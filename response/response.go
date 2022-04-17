package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Info(c *gin.Context, httpStatus int, status int, code int, data interface{}, message string) {
	c.JSON(httpStatus, gin.H{
		"status":  status,
		"code":    code,
		"data":    data,
		"message": message,
	})
}

func Success(c *gin.Context, data interface{}, message string) {
	Info(c, http.StatusOK, 1, 200, data, message)
}

func Unauthorized(c *gin.Context, message string) {
	Info(c, http.StatusUnauthorized, -1, 401, nil, message)
}

func NotFound(c *gin.Context) {
	Info(c, http.StatusNotFound, -1, 404, nil, "请求资源不存在")
}

func Fail(c *gin.Context, data interface{}, message string) {
	Info(c, http.StatusBadRequest, -1, 400, nil, message)
}
