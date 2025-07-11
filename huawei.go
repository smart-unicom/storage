// Package storage 华为云OBS存储提供者
package storage

import (
	"github.com/smart-unicom/oss"
	"github.com/smart-unicom/oss/huawei"
)

// NewHuaweiObsStorageProvider 创建华为云OBS存储提供者
// 参数:
//   - clientId: 访问密钥ID
//   - clientSecret: 访问密钥
//   - region: 区域
//   - bucket: 存储桶名称
//   - endpoint: 服务端点
//
// 返回:
//   - oss.StorageInterface: 存储接口实例
//   - error: 错误信息
func NewHuaweiObsStorageProvider(clientId string, clientSecret string, region string, bucket string, endpoint string) oss.StorageInterface {
	// 创建华为云OBS存储配置
	sp := huawei.New(&huawei.Config{
		SecretID:  clientId,
		SecretKey: clientSecret,
		Region:    region,
		Bucket:    bucket,
		Endpoint:  endpoint,
	})

	return sp
}
