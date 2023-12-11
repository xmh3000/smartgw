package io

import (
	"go.uber.org/zap"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func GetCurrentPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)

	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return ""
	}

	return string(path[0 : i+1])
}

func SureExists(dir string) {
	cur := GetCurrentPath()
	_, err := os.Stat(cur + dir)
	if err != nil {
		err := os.Mkdir(cur+dir, 0777)
		if err != nil {
			zap.S().Errorf("创建%s文件夹失败 %v", cur+dir, err)
		}
		_ = os.Chmod(cur+dir, 0777)
	}
}

func PathExists(dir string) (bool, error) {
	cur := GetCurrentPath()

	_, err := os.Stat(cur + dir)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func ReadFile(filename string) ([]byte, error) {
	return os.ReadFile(GetCurrentPath() + filename)
}

func WriteFile(filename string, data []byte) error {
	return os.WriteFile(GetCurrentPath()+filename, data, 0777)
}
