package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*") // *代理允许访问所有域 正式环境慎用
		c.Header("Access-Control-Allow-Headers", "Cache-Control, Content-Language, Content-Type, Expires, Last-Modified, Pragma, Authorization, x-role")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")                                                   //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type") // 首部可以作为响应的一部分暴露给外部
		c.Header("Access-Control-Allow-Credentials", "true")                                                                                 // 允许客户端携带验证信息

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}
