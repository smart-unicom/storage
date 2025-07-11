// Package storage MinIO对象存储提供者
package storage

import (
	awss3 "github.com/aws/aws-sdk-go/service/s3"
	"github.com/smart-unicom/oss"
	"github.com/smart-unicom/oss/s3"
)

// NewMinIOStorageProvider 创建MinIO存储提供者
// 参数:
//   - clientId: MinIO的访问ID
//   - clientSecret: MinIO的访问密钥
//   - region: MinIO的地域
//   - bucket: MinIO的存储桶名称
//   - endpoint: MinIO的端点地址
// 返回:
//   - oss.StorageInterface: 存储接口实例
func NewMinIOStorageProvider(clientId string, clientSecret string, region string, bucket string, endpoint string) oss.StorageInterface {
	// 创建MinIO存储配置
	sp := s3.New(&s3.Config{
		AccessId:         clientId,
		AccessKey:        clientSecret,
		Region:           region,
		Bucket:           bucket,
		Endpoint:         endpoint,
		S3Endpoint:       endpoint,
		ACL:              awss3.BucketCannedACLPublicReadWrite, // 设置为公共读写权限
		S3ForcePathStyle: true,                                 // 强制使用路径样式
	})

	return sp
}
