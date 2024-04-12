package fileutil

import (
	"os"
	"strings"
)

// DirExist 检查目录是否存在，不存在则创建
// dirPath: 目录路径
//
// return: 错误信息
func DirExist(dirPath string) error {
	_, err := os.Stat(dirPath)
	if err == nil {
		return nil
	}

	if os.IsNotExist(err) {
		err = os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			return err
		}
		return nil
	}
	return err
}

// DeleteFile 删除文件
// filePath: 文件路径
//
// return: 错误信息
func DeleteFile(filePath string) error {
	return os.Remove(filePath)
}

// CheckFileFormat 检查文件格式是否在允许范围内
// fileName: 文件名
// allowedFormats: 允许的文件格式
//
// return: 是否在允许范围内(true:在允许范围内;false:不在允许范围内)
func CheckFileFormat(fileName string, allowedFormats []string) bool {
	fileExt := strings.ToLower(fileName[strings.LastIndex(fileName, ".")+1:])
	for _, f := range allowedFormats {
		if fileExt == f {
			return true
		}
	}
	return false
}
