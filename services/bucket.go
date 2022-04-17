package services

import (
	"minio_server/global"
	"minio_server/models"
	"minio_server/response"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

type IBucketsService interface {
	List(c *gin.Context)
	Exists(c *gin.Context)
	Remove(c *gin.Context)
	ListObjects(c *gin.Context)
}

type BucketsService struct{}

func NewBucketsService() IBucketsService {
	return &BucketsService{}
}

// List godoc
// @Summary 获取存储桶列表
// @Tags buckets
// @Accept json
// @security ApiKeyAuth
// @Success 200 "{"message": "获取成功", status: 1}"
// @Failure 400 "{"message": "获取失败", status: -1}"
// @Router /api/buckets/list [get]
func (*BucketsService) List(c *gin.Context) {

	buckets, err := global.MinioClient.ListBuckets(c)
	if err != nil {
		response.Fail(c, nil, err.Error())
	}

	response.Success(c, buckets, "获取列表成功")
}

// Exists godoc
// @Summary 获取存储桶详细信息
// @Tags buckets
// @Accept json
// @security ApiKeyAuth
// @security ApiKeyXRole
// @Success 200 "{"message": "获取成功", status: 1}"
// @Failure 400 "{"message": "获取失败", status: -1}"
// @Router /api/buckets/exists [get]
func (*BucketsService) Exists(c *gin.Context) {
	bucketName := c.Query("bucket")

	if len(bucketName) <= 0 {
		response.Fail(c, nil, "桶名为空")
		return
	}

	ok, _ := global.MinioClient.BucketExists(c, bucketName)
	if !ok {
		response.Fail(c, nil, "查询不到该桶")
		return
	}

	response.Success(c, nil, "查询成功")
}

// Remove godoc
// @Summary 删除存储桶
// @Tags buckets
// @Accept json
// @security ApiKeyAuth
// @security ApiKeyXRole
// @Param bucket query string true "存储桶名必传"
// @Success 200 "{"message": "删除成功", status: 1}"
// @Failure 400 "{"message": "删除失败", status: -1}"
// @Router /api/buckets/remove [post]
func (*BucketsService) Remove(c *gin.Context) {
	bucketName := c.Query("bucket")

	if len(bucketName) <= 0 {
		response.Fail(c, nil, "桶名为空, 无法删除")
		return
	}

	err := global.MinioClient.RemoveBucket(c, bucketName)
	if err != nil {
		response.Fail(c, nil, "删除失败")
		return
	}

	response.Success(c, nil, "删除成功")

}

// ListObjects godoc
// @Summary 获取存储桶内所有文件列表
// @Tags buckets
// @Accept json
// @security ApiKeyAuth
// @Param bucket query string true "存储桶名必传"
// @Success 200 "{"message": "获取成功", status: 1}"
// @Failure 400 "{"message": "获取失败", status: -1}"
// @Router /api/buckets/listobjects [get]
func (*BucketsService) ListObjects(c *gin.Context) {
	bucketName := c.Query("bucket")

	if len(bucketName) <= 0 {
		response.Fail(c, nil, "桶名为空, 无法获取列表")
		return
	}

	list := make(chan []models.ListObjects)

	go func() {
		defer close(list)

		var arr []models.ListObjects

		opts := minio.ListObjectsOptions{
			UseV1:     true,
			Recursive: true,
		}

		objects := global.MinioClient.ListObjects(c, bucketName, opts)

		for object := range objects {
			if object.Err != nil {
				return
			}

			arr = append(arr, models.ListObjects{
				Name:         object.Key,
				Size:         int(object.Size),
				LastModified: object.LastModified.Format("2006-01-02 15:04:05"),
			})
		}

		list <- arr

	}()

	data, ok := <-list

	if !ok {
		response.Fail(c, nil, "指定的存储桶不存在")
		return
	}

	response.Success(c, data, "查询成功")
}
