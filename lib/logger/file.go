package logger

import (
	"fmt"
	"os"
)

// checkNotExist 检查文件夹路径是否存在
func checkNotExist(src string) bool {
	_, err := os.Stat(src)
	return os.IsNotExist(err)
}

// checkPermission 检查是否有文件夹权限
func checkPermission(src string) bool {
	_, err := os.Stat(src)
	return os.IsPermission(err)
}

// isNotExistMkDir 检查文件夹路径是否存在, 不存在则创建
func isNotExistMkDir(src string) error {
	if notExist := checkNotExist(src); notExist == true {
		if err := mkDir(src); err != nil {
			return err
		}
	}
	return nil
}

// mkDir 创建文件夹
func mkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func mustOpen(fileName, dir string) (*os.File, error) {
	perm := checkPermission(dir)
	if perm {
		return nil, fmt.Errorf("permission denied dir: %s", dir)
	}
	err := isNotExistMkDir(dir)
	if err != nil {
		return nil, fmt.Errorf("error during make dir %s, err: %s", dir, err)
	}
	// 644权限: rw-r--r--
	f, err := os.OpenFile(dir+string(os.PathSeparator)+fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("fail to open file, err: %s", err)
	}
	return f, nil
}
