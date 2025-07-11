// Package storage 腾讯云对象存储提供者
package storage

import (
	"github.com/smart-unicom/oss"
	"github.com/smart-unicom/oss/tencent"
)

// NewTencentCosStorageProvider 创建腾讯云对象存储提供者
// 参数:
//   - clientId: 密钥ID
//   - clientSecret: 密钥
//   - region: 区域
//   - bucket: 存储桶名称
//   - endpoint: 服务端点
//
// 返回:
//   - oss.StorageInterface: 存储接口实例
func NewTencentCosStorageProvider(clientId string, clientSecret string, region string, bucket string, endpoint string) oss.StorageInterface {
	// 创建腾讯云COS存储配置
	sp := tencent.New(&tencent.Config{
		SecretID:  clientId,
		SecretKey: clientSecret,
		Region:    region,
		Bucket:    bucket,
		Endpoint:  endpoint,
	})

	return sp
}
