package services

import (
	"encoding/base64"
	"minio_server/common"
	"minio_server/models"
	"minio_server/repositories"
	"minio_server/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type IUserService interface {
	Login(c *gin.Context)
	UserInfo(c *gin.Context)
}

type UserService struct {
	UserRepository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) IUserService {
	return &UserService{UserRepository: repository}
}

// Login godoc
// @Summary 登录
// @Tags users
// @Accept json
// @Param bucket body swagger.Login true "账号密码必须填"
// @Success 200 "{"message": "登录成功", status: 1}"
// @Failure 400 "{"message": "登录失败", status: -1}"
// @Router /api/user/login [post]
func (u *UserService) Login(c *gin.Context) {
	var reqInfo models.User

	if err := c.ShouldBindBodyWith(&reqInfo, binding.JSON); err != nil {
		response.Fail(c, nil, err.Error())
	} else {

		if len(reqInfo.SecretKey) < 6 {
			response.Info(c, http.StatusUnprocessableEntity, -1, 422, nil, "密码必须大于6位数!")
			return
		}

		if len(reqInfo.AccessKey) == 0 {
			response.Info(c, http.StatusUnprocessableEntity, -1, 422, nil, "用户名称不能为空!")
			return
		}

		if token, errLogin := u.UserRepository.Login(&reqInfo); errLogin != nil {
			response.Info(c, http.StatusPreconditionFailed, -1, 412, nil, errLogin.Error())
		} else {
			response.Success(c, token, "登陆成功")
		}
	}
}

// UserInfo godoc
// @Summary 获取用户信息
// @Description 注意: header设置token时前面带 Bearer + 空格
// @Tags users
// @Accept json
// @security ApiKeyAuth
// @Success 200 "{"message": "获取成功", status: 1}"
// @Failure 400 "{"message": "获取失败", status: -1}"
// @Router /api/user/info [get]
func (*UserService) UserInfo(c *gin.Context) {
	// 如果用户信息随token一起发送给了前端, 这个方法可以去掉.
	tokenString := c.GetHeader("Authorization")
	tokenString = tokenString[7:]
	_, claims, err := common.ParseToken(tokenString)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}

	access := base64.StdEncoding.EncodeToString([]byte(claims.Access))

	if len(access) <= 0 {
		response.Fail(c, nil, "获取信息失败")
		return
	}
	response.Success(c, access, "成功获取身份信息")
}
