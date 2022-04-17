package models

import "time"

type User struct {
	UserID     int16     `sql:"user_id" json:"user_id"`         // 用户ID
	Access     string    `sql:"access" json:"access"`           // 用户权限
	AccessKey  string    `sql:"access_key" json:"access_key"`   // 用户名称
	SecretKey  string    `sql:"secret_key" json:"secret_key"`   // 用户密码
	Level      int       `sql:"level" json:"level"`             // 用户等级
	CreateTime time.Time `sql:"create_time" json:"create_time"` // 创建时间
	UpdateTime time.Time `sql:"update_time" json:"update_time"` // 更新时间
}
