// Package storage 群晖NAS存储提供者
package storage

import (
	"github.com/smart-unicom/oss"
	"github.com/smart-unicom/oss/synology"
)

// NewSynologyNasStorageProvider 创建群晖NAS存储提供者
// 参数:
//   - clientId: 群晖NAS的访问ID
//   - clientSecret: 群晖NAS的访问密钥
//   - endpoint: 群晖NAS的API端点
// 返回:
//   - oss.StorageInterface: 存储接口实例
func NewSynologyNasStorageProvider(clientId string, clientSecret string, endpoint string) oss.StorageInterface {
	// 创建群晖NAS存储配置
	sp := synology.New(&synology.Config{
		AccessId:     clientId,
		AccessKey:    clientSecret,
		Endpoint:     endpoint,
		SharedFolder: "/home", // 默认共享文件夹
	})

	return sp
}
