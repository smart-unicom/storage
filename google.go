// Package storage 谷歌云存储提供者
package storage

import (
	"github.com/smart-unicom/oss"
	"github.com/smart-unicom/oss/googlecloud"
)

// NewGoogleCloudStorageProvider 创建谷歌云存储提供者
// 参数:
//   - clientSecret: 服务账户JSON密钥
//   - bucket: 存储桶名称
//   - endpoint: 服务端点
// 返回:
//   - oss.StorageInterface: 存储接口实例
func NewGoogleCloudStorageProvider(clientSecret string, bucket string, endpoint string) oss.StorageInterface {
	// 创建谷歌云存储配置
	sp, err := googlecloud.New(&googlecloud.Config{
		ServiceAccountJson: clientSecret,
		Bucket:             bucket,
		Endpoint:           endpoint,
	})
	if err != nil {
		panic(err)
	}

	return sp
}
