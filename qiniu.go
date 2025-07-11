// Package storage 七牛云对象存储提供者package storage
package storage

import (
	"github.com/smart-unicom/oss"
	"github.com/smart-unicom/oss/qiniu"
)

// NewQiniuKodoStorageProvider 创建七牛云对象存储提供者
// 参数:
//   - clientId: 访问密钥ID
//   - clientSecret: 访问密钥
//   - region: 区域
//   - bucket: 存储空间名称
//   - endpoint: 服务端点
//
// 返回:
//   - oss.StorageInterface: 存储接口实例
//   - error: 错误信息
func NewQiniuKodoStorageProvider(clientId string, clientSecret string, region string, bucket string, endpoint string) (oss.StorageInterface, error) {
	// 创建七牛云存储配置
	sp, err := qiniu.New(&qiniu.Config{
		AccessId:  clientId,
		AccessKey: clientSecret,
		Region:    region,
		Bucket:    bucket,
		Endpoint:  endpoint,
	})

	return sp, err
}
