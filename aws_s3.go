// Package storage AWS S3存储提供者
package storage

import (
	awss3 "github.com/aws/aws-sdk-go/service/s3"
	"github.com/smart-unicom/oss"
	"github.com/smart-unicom/oss/s3"
)

// NewAwsS3StorageProvider 创建AWS S3存储提供者
// 参数:
//   - clientId: 访问密钥ID
//   - clientSecret: 访问密钥
//   - region: 区域
//   - bucket: 存储桶名称
//   - endpoint: 服务端点
// 返回:
//   - oss.StorageInterface: 存储接口实例
func NewAwsS3StorageProvider(clientId string, clientSecret string, region string, bucket string, endpoint string) oss.StorageInterface {
	// 创建AWS S3配置
	sp := s3.New(&s3.Config{
		AccessId:   clientId,
		AccessKey:  clientSecret,
		Region:     region,
		Bucket:     bucket,
		Endpoint:   endpoint,
		S3Endpoint: endpoint,
		ACL:        awss3.BucketCannedACLPublicRead, // 设置为公共读权限
	})

	return sp
}
