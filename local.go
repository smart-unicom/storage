// Package storage 本地文件系统存储提供者
package storage

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/smart-unicom/oss"
)

// LocalFileSystemProvider 本地文件系统存储提供者
// 提供基于本地文件系统的存储功能
type LocalFileSystemProvider struct {
	BaseDir string // 基础目录路径
}

// NewLocalFileSystemStorageProvider 初始化本地文件系统存储提供者
// 返回:
//   - *LocalFileSystemProvider: 本地文件系统存储提供者实例
func NewLocalFileSystemStorageProvider() *LocalFileSystemProvider {
	baseFolder := "files"
	// 获取基础文件夹的绝对路径
	absBase, err := filepath.Abs(baseFolder)
	if err != nil {
		panic(err)
	}

	return &LocalFileSystemProvider{BaseDir: absBase}
}

// GetFullPath 从绝对路径或相对路径获取完整路径
// 参数:
//   - path: 文件路径
// 返回:
//   - string: 完整的文件路径
func (sp LocalFileSystemProvider) GetFullPath(path string) string {
	fullPath := path
	// 如果路径不是以基础目录开头，则拼接基础目录
	if !strings.HasPrefix(path, sp.BaseDir) {
		fullPath, _ = filepath.Abs(filepath.Join(sp.BaseDir, path))
	}
	return fullPath
}

// Get 根据给定路径获取文件
// 参数:
//   - path: 文件路径
// 返回:
//   - *os.File: 文件对象
//   - error: 错误信息
func (sp LocalFileSystemProvider) Get(path string) (*os.File, error) {
	return os.Open(sp.GetFullPath(path))
}

// GetStream 以流的形式获取文件
// 参数:
//   - path: 文件路径
// 返回:
//   - io.ReadCloser: 可读可关闭的流
//   - error: 错误信息
func (sp LocalFileSystemProvider) GetStream(path string) (io.ReadCloser, error) {
	return os.Open(sp.GetFullPath(path))
}

// Put 将读取器的内容存储到指定路径
// 参数:
//   - path: 存储路径
//   - reader: 数据读取器
// 返回:
//   - *oss.Object: 存储对象
//   - error: 错误信息
func (sp LocalFileSystemProvider) Put(path string, reader io.Reader) (*oss.Object, error) {
	fullPath := sp.GetFullPath(path)

	// 创建目录结构
	err := os.MkdirAll(filepath.Dir(fullPath), os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("创建文件夹失败: \"%s\" 本地文件系统存储提供者错误: %s. 请确保进程有正确的权限创建/访问该目录，或者手动提前创建", filepath.Dir(fullPath), err.Error())
	}

	// 创建目标文件
	dst, err := os.Create(filepath.Clean(fullPath))
	if err == nil {
		defer dst.Close()
		// 如果读取器支持定位，则重置到开始位置
		if seeker, ok := reader.(io.ReadSeeker); ok {
			seeker.Seek(0, 0)
		}
		// 复制数据到目标文件
		_, err = io.Copy(dst, reader)
	}
	return &oss.Object{Path: path, Name: filepath.Base(path), StorageInterface: sp}, err
}

// Delete 删除文件
// 参数:
//   - path: 文件路径
// 返回:
//   - error: 错误信息
func (sp LocalFileSystemProvider) Delete(path string) error {
	return os.Remove(sp.GetFullPath(path))
}

// List 列出当前路径下的所有对象
// 参数:
//   - path: 目录路径
// 返回:
//   - []*oss.Object: 对象列表
//   - error: 错误信息
func (sp LocalFileSystemProvider) List(path string) ([]*oss.Object, error) {
	objects := []*oss.Object{}
	fullPath := sp.GetFullPath(path)

	// 遍历目录下的所有文件
	filepath.Walk(fullPath, func(path string, info os.FileInfo, err error) error {
		// 跳过根目录本身
		if path == fullPath {
			return nil
		}

		// 只处理文件，不处理目录
		if err == nil && !info.IsDir() {
			modTime := info.ModTime()
			// 创建对象信息
			objects = append(objects, &oss.Object{
				Path:             strings.TrimPrefix(path, sp.BaseDir),
				Name:             info.Name(),
				LastModified:     &modTime,
				StorageInterface: sp,
			})
		}
		return nil
	})

	return objects, nil
}

// GetEndpoint 获取端点，本地文件系统提供者的端点是 /
// 返回:
//   - string: 端点路径
func (sp LocalFileSystemProvider) GetEndpoint() string {
	return "/"
}

// GetURL 获取公共可访问的URL
// 参数:
//   - path: 文件路径
// 返回:
//   - url: 文件URL
//   - err: 错误信息
func (sp LocalFileSystemProvider) GetURL(path string) (url string, err error) {
	return path, nil
}
