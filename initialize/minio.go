package initialize

import (
	"log"
	"minio_server/global"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func InitMinIO() error {
	minioInfo := global.Settings.MinioInfo

	minioClient, err := minio.New(minioInfo.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(minioInfo.AccessKey, minioInfo.SecretKey, ""),
		Secure: false,
	})
	if err != nil {
		global.MinioClient = nil
		log.Fatalln(err)

		return err
	}

	global.MinioClient = minioClient
	log.Println("Minio Init Success")

	return nil
}
