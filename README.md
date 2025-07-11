# storage
Golang 对象存储组件，支持多种云存储服务，包括本地文件系统、AWS S3、MinIO、阿里云 OSS、腾讯云 COS、Azure Blob、七牛云 Kodo、谷歌云存储、华为云 OBS、群晖 NAS 等。

## 功能特性

本组件提供统一的对象存储接口，支持多种云存储服务和本地文件系统：

- **本地文件系统** - 基于本地磁盘的文件存储
- **AWS S3** - 亚马逊简单存储服务
- **MinIO** - 高性能分布式对象存储
- **阿里云 OSS** - 阿里云对象存储服务
- **腾讯云 COS** - 腾讯云对象存储
- **Azure Blob** - 微软Azure Blob存储
- **七牛云 Kodo** - 七牛云对象存储
- **谷歌云存储** - Google Cloud Storage
- **华为云 OBS** - 华为云对象存储服务
- **群晖 NAS** - Synology网络附加存储

## 使用方法

```go
import "github.com/smart-unicom/storage"

// 创建存储提供者
provider, err := storage.NewStorageProvider(
    storage.HUAWEI_CLOUD_OBS, // 存储类型
    "your-access-key",        // 访问密钥ID
    "your-secret-key",        // 访问密钥
    "cn-north-1",             // 区域
    "your-bucket",            // 存储桶名称
    "https://obs.cn-north-1.myhuaweicloud.com", // 端点
    "",                       // 证书（可选）
    "",                       // 内容（可选）
)
if err != nil {
    log.Fatal(err)
}

// 使用存储接口进行文件操作
// ...
```

## 支持的存储类型

| 存储类型 | 常量 | 说明 |
|---------|------|------|
| 本地文件系统 | `LOCAL_FILE_SYSTEM` | 基于本地磁盘存储 |
| AWS S3 | `AWS_S3` | 亚马逊S3存储服务 |
| MinIO | `MINIO` | MinIO对象存储 |
| 阿里云OSS | `ALIYUN_OSS` | 阿里云对象存储 |
| 腾讯云COS | `TENCENT_CLOUD_COS` | 腾讯云对象存储 |
| Azure Blob | `AZURE_BLOB` | 微软Azure Blob存储 |
| 七牛云Kodo | `QINIU_CLOUD_KODO` | 七牛云对象存储 |
| 谷歌云存储 | `GOOGLE_CLOUD_STORAGE` | 谷歌云存储服务 |
| 华为云OBS | `HUAWEI_CLOUD_OBS` | 华为云对象存储服务 |
| 群晖NAS | `SYNOLOGY` | Synology网络附加存储 |
