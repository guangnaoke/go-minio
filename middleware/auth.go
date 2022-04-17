package middleware

import (
	"encoding/base64"
	"minio_server/common"
	"minio_server/global"
	"minio_server/models"
	"minio_server/response"
	"strings"

	"github.com/gin-gonic/gin"
)

// token验证
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			response.Unauthorized(c, "权限不足")
			c.Abort()
			return
		}

		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			response.Unauthorized(c, "权限不足")
			c.Abort()
			return
		}

		userId := claims.UserID
		var user models.User

		if errSearch := global.DB.Table("user").First(&user, userId).Error; errSearch != nil {
			response.Fail(c, nil, errSearch.Error())
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

// 身份验证 casbin权限验证
func AuthCheckRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取头部x-role 身份
		roleString := c.GetHeader("x-role")
		if roleString == "" {
			response.Unauthorized(c, "无效身份")
			c.Abort()
			return
		}

		// base64 解密
		role, err := base64.StdEncoding.DecodeString(roleString)
		if err != nil {
			response.Unauthorized(c, "无效身份")
			c.Abort()
			return
		}

		e := common.Casbin()

		//检查权限
		res, errRes := e.Enforce(string(role), c.Request.URL.Path, c.Request.Method)
		if errRes != nil {
			response.Fail(c, nil, errRes.Error())
			c.Abort()
			return
		}
		if res {
			c.Next()
		} else {
			response.Unauthorized(c, "权限不足")
			c.Abort()
			return
		}
	}
}
