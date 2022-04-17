package repositories

import (
	"errors"
	"minio_server/common"
	"minio_server/global"
	"minio_server/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	Login(*models.User) (string, error)
}

type UserManageRepository struct {
	table string
}

func NewUserManagerRepository(table string, sql *gorm.DB) UserRepository {
	return &UserManageRepository{
		table: table,
	}
}

func (*UserManageRepository) Login(user *models.User) (string, error) {

	if global.DB == nil {
		return "", errors.New("数据库连接失败")
	}

	// 初始化user表, 不是必须
	global.DB.AutoMigrate(&models.User{})

	var m models.User

	// 上面table字段是表名, 有的喜欢用DB.table查询, 我就保留下来了, 不需要的可以删除
	if err := global.DB.Where("access_key = ?", &user.AccessKey).First(&m).Error; err != nil {
		if m.UserID == 0 {
			return "", errors.New("用户不存在")
		}
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(m.SecretKey), []byte(user.SecretKey)); err != nil {
		return "", errors.New("密码错误")
	}

	token, err := common.ReleaseToken(m)
	if err != nil {
		return "", err
	}

	return token, nil
}
