// Package storage Azure Blob存储提供者
package storage

import (
	"github.com/smart-unicom/oss"
	"github.com/smart-unicom/oss/azureblob"
)

// NewAzureStorageProvider 创建Azure Blob存储提供者
// 参数:
//   - clientId: 客户端ID
//   - clientSecret: 客户端密钥
//   - region: 区域
//   - bucket: 存储容器名称
//   - endpoint: 服务端点
//
// 返回:
//   - oss.StorageInterface: 存储接口实例
func NewAzureStorageProvider(clientId string, clientSecret string, region string, bucket string, endpoint string) oss.StorageInterface {
	// 创建Azure Blob存储配置
	sp := azureblob.New(&azureblob.Config{
		AccessId:  clientId,
		AccessKey: clientSecret,
		Region:    region,
		Bucket:    bucket,
		Endpoint:  endpoint,
	})
	return sp
}
