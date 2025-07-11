// Package storage 存储服务提供者相关功能
// 支持多种云存储和本地文件系统的统一接口
package storage

import "github.com/smart-unicom/oss"

// StorageType 存储类型定义
type StorageType string

// 支持的存储类型常量定义
const (
	LOCAL_FILE_SYSTEM    StorageType = "Local File System"    // 本地文件系统
	AWS_S3               StorageType = "AWS S3"               // 亚马逊S3存储
	MINIO                StorageType = "MinIO"                // MinIO对象存储
	ALIYUN_OSS           StorageType = "Aliyun OSS"           // 阿里云对象存储
	TENCENT_CLOUD_COS    StorageType = "Tencent Cloud COS"    // 腾讯云对象存储
	AZURE_BLOB           StorageType = "Azure Blob"           // 微软Azure Blob存储
	QINIU_CLOUD_KODO     StorageType = "Qiniu Cloud Kodo"     // 七牛云对象存储
	GOOGLE_CLOUD_STORAGE StorageType = "Google Cloud Storage" // 谷歌云存储
	HUAWEI_CLOUD_OBS     StorageType = "Huawei Cloud OBS"     // 华为云对象存储
	SYNOLOGY             StorageType = "Synology"             // 群晖NAS存储
)

// NewStorageProvider 创建存储服务提供者
// 参数:
//   - pt: 存储类型
//   - clientId: 客户端ID
//   - clientSecret: 客户端密钥
//   - region: 区域
//   - bucket: 存储桶名称
//   - endpoint: 服务端点
//   - cert: 证书
//   - content: 内容
//
// 返回:
//   - oss.StorageInterface: 存储接口实例
//   - error: 错误信息
func NewStorageProvider(pt StorageType, clientId string, clientSecret string, region string,
	bucket string, endpoint string, cert string, content string) (oss.StorageInterface, error) {
	switch pt {
	case LOCAL_FILE_SYSTEM:
		// 创建本地文件系统存储提供者
		return NewLocalFileSystemStorageProvider(), nil
	case AWS_S3:
		// 创建AWS S3存储提供者
		return NewAwsS3StorageProvider(clientId, clientSecret, region, bucket, endpoint), nil
	case MINIO:
		// 创建MinIO存储提供者
		return NewMinIOStorageProvider(clientId, clientSecret, "_", bucket, endpoint), nil
	case ALIYUN_OSS:
		// 创建阿里云OSS存储提供者
		return NewAliyunOssStorageProvider(clientId, clientSecret, region, bucket, endpoint), nil
	case TENCENT_CLOUD_COS:
		// 创建腾讯云COS存储提供者
		return NewTencentCosStorageProvider(clientId, clientSecret, region, bucket, endpoint), nil
	case AZURE_BLOB:
		// 创建Azure Blob存储提供者
		return NewAzureStorageProvider(clientId, clientSecret, region, bucket, endpoint), nil
	case QINIU_CLOUD_KODO:
		// 创建七牛云存储提供者
		return NewQiniuKodoStorageProvider(clientId, clientSecret, region, bucket, endpoint)
	case GOOGLE_CLOUD_STORAGE:
		// 创建谷歌云存储提供者
		return NewGoogleCloudStorageProvider(clientSecret, bucket, endpoint), nil
	case HUAWEI_CLOUD_OBS:
		// 创建华为云OBS存储提供者
		return NewHuaweiObsStorageProvider(clientId, clientSecret, region, bucket, endpoint), nil
	case SYNOLOGY:
		// 创建群晖NAS存储提供者
		return NewSynologyNasStorageProvider(clientId, clientSecret, endpoint), nil

	}
	// 未知存储类型，返回空值
	return nil, nil
}
