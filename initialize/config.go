package initialize

import (
	"io/ioutil"
	"log"
	"minio_server/conf"
	"minio_server/global"
	"os"

	"gopkg.in/yaml.v3"
)

func InitConfig() error {
	workDor, _ := os.Getwd()
	yamlFile, err := ioutil.ReadFile(workDor + "/conf/conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err %v", err)
		return err
	}

	serverConfig := conf.ServerConf{}

	err = yaml.Unmarshal(yamlFile, &serverConfig)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		return err
	}

	global.Settings = serverConfig

	return nil
}
