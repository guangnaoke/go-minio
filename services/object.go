package services

import (
	"fmt"
	"minio_server/global"
	"minio_server/response"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

type IObjectService interface {
	Get(c *gin.Context)
	GetObjectUrl(c *gin.Context)
	Stat(c *gin.Context)
	Remove(c *gin.Context)
	Put(c *gin.Context)
}

type ObjectService struct{}

func NewObjectService() IObjectService {
	return &ObjectService{}
}

func (*ObjectService) Get(c *gin.Context) {
	bucketName := c.Query("bucket")
	objectName := c.Query("object")

	if bucketName == "" || objectName == "" {
		response.Fail(c, nil, "请检查传参是否正确")
		return
	}

	reader, err := global.MinioClient.GetObject(c, bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}

	response.Success(c, nil, "获取文件成功")
	defer reader.Close()
}

// GetObjectUrl godoc
// @Summary 获取文件的url
// @Tags objects
// @Accept json
// @security ApiKeyAuth
// @security ApiKeyXRole
// @Param bucket query string true "存储桶名必传"
// @Param object query string true "文件名必传"
// @Success 200 "{"message": "获取成功", status: 1}"
// @Failure 400 "{"message": "获取失败", status: -1}"
// @Router /api/object/url [get]
func (*ObjectService) GetObjectUrl(c *gin.Context) {
	bucketName := c.Query("bucket")
	objectName := c.Query("object")

	if bucketName == "" || objectName == "" {
		response.Fail(c, nil, "请检查传参是否正确")
		return
	}

	reqParams := make(url.Values)
	fileName := fmt.Sprintf("attachment; filename=\"%s\"", objectName)
	reqParams.Set("response-content-disposition", fileName)

	presignedURL, err := global.MinioClient.PresignedGetObject(c, bucketName, objectName, time.Duration(1000)*time.Second, reqParams)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}

	// presignedURL 一定要string
	response.Success(c, presignedURL.String(), "获取文件路径成功")
}

func (*ObjectService) Stat(c *gin.Context) {
	bucketName := c.Query("bucket")
	objectName := c.Query("object")

	if bucketName == "" || objectName == "" {
		response.Fail(c, nil, "请检查传参是否正确")
		return
	}

	stat, err := global.MinioClient.StatObject(c, bucketName, objectName, minio.StatObjectOptions{})
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}

	response.Success(c, stat, "获取文件信息成功")
}

// Remove godoc
// @Summary 删除文件
// @Tags objects
// @Accept json
// @security ApiKeyAuth
// @security ApiKeyXRole
// @Param bucket query string true "存储桶名必传"
// @Param object query string true "文件名必传"
// @Success 200 "{"message": "删除成功", status: 1}"
// @Failure 400 "{"message": "删除失败", status: -1}"
// @Router /api/object/remove [post]
func (*ObjectService) Remove(c *gin.Context) {
	bucketName := c.Query("bucket")
	objectName := c.Query("object")

	if bucketName == "" || objectName == "" {
		response.Fail(c, nil, "请检查传参是否正确")
		return
	}

	opts := minio.RemoveObjectOptions{
		GovernanceBypass: true,
	}

	err := global.MinioClient.RemoveObject(c, bucketName, objectName, opts)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}

	response.Success(c, nil, "删除成功")
}

// Upload godoc
// @Summary 上传文件
// @Tags objects
// @Accept json
// @security ApiKeyAuth
// @security ApiKeyXRole
// @Accept multipart/form-data
// @Param file formData file true "文件名必传"
// @Param bucket formData string true "存储桶名必传"
// @Success 200 "{"message": "删除成功", status: 1}"
// @Failure 400 "{"message": "删除失败", status: -1}"
// @Router /api/object/upload [post]
func (*ObjectService) Put(c *gin.Context) {
	file, _ := c.FormFile("file")
	bucket := c.PostForm("bucket")

	reader, errFile := file.Open()
	if errFile != nil {
		response.Fail(c, nil, errFile.Error())
		return
	}
	uoloadInfo, err := global.MinioClient.PutObject(c, bucket, file.Filename, reader, file.Size, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}

	response.Success(c, uoloadInfo, "文件上传成功")
}
