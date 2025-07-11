// Package storage 阿里云OSS存储提供者
package storage

import (
	"github.com/smart-unicom/oss"
	"github.com/smart-unicom/oss/aliyun"
)

// NewAliyunOssStorageProvider 创建阿里云OSS存储提供者
// 参数:
//   - clientId: 访问密钥ID
//   - clientSecret: 访问密钥
//   - region: 区域
//   - bucket: 存储桶名称
//   - endpoint: 服务端点
// 返回:
//   - oss.StorageInterface: 存储接口实例
func NewAliyunOssStorageProvider(clientId string, clientSecret string, region string, bucket string, endpoint string) oss.StorageInterface {
	// 创建阿里云OSS配置
	sp := aliyun.New(&aliyun.Config{
		AccessId:  clientId,
		AccessKey: clientSecret,
		Bucket:    bucket,
		Endpoint:  endpoint,
	})

	return sp
}
