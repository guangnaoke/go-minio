package common

import (
	"log"
	"minio_server/global"
	"os"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
)

//权限结构
type CasbinModel struct {
	Ptype    string `json:"p_type" bson:"p_type"`
	RoleName string `json:"rolename" bson:"v0"`
	Path     string `json:"path" bson:"v1"`
	Method   string `json:"method" bson:"v2"`
}

//添加权限
func (c *CasbinModel) AddCasbin(cm CasbinModel) bool {
	e := Casbin()
	isTrue, _ := e.AddPolicy(cm.RoleName, cm.Path, cm.Method)
	return isTrue
}

//持久化到数据库
func Casbin() *casbin.Enforcer {
	workDor, _ := os.Getwd()
	if global.DB == nil {
		log.Fatalln("数据库连接失败")
	}

	g, _ := gormadapter.NewAdapterByDB(global.DB)
	c, _ := casbin.NewEnforcer(workDor+"/conf/auth_model.conf", g)
	c.LoadPolicy()
	return c
}

var (
	casbins = CasbinModel{}
)

func AddCasbin(c *gin.Context) {
	rolename := c.PostForm("rolename")
	path := c.PostForm("path")
	method := c.PostForm("method")
	ptype := "p"
	casbin := CasbinModel{
		Ptype:    ptype,
		RoleName: rolename,
		Path:     path,
		Method:   method,
	}
	isok := casbins.AddCasbin(casbin)
	if isok {
		log.Println("Add Cabin Success")
	} else {
		log.Println("Add Cabin Error")
	}

}
