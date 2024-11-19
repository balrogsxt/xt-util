package xfile

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// IsDir 判断文件夹是否存在
func IsDir(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// IsFile 判断文件是否存在
func IsFile(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
func Exists(path string) bool {
	if stat, err := os.Stat(path); stat != nil && !os.IsNotExist(err) {
		return true
	}
	return false
}

// Join 文件路径组合
func Join(elem ...string) string {
	return filepath.Join(elem...)
}

// Mkdir 创建文件夹
func Mkdir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}
func GetContent(file string) (string, error) {
	data, err := GetBytes(file)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
func GetBytes(file string) ([]byte, error) {
	return os.ReadFile(file)
}

func PutContent(file, content string) error {
	return PutBytes(file, []byte(content))
}

func Dir(path string) string {
	return filepath.Dir(path)
}
func Abs(path string) string {
	p, _ := filepath.Abs(path)
	return p
}
func FixPath(path string) string {
	switch runtime.GOOS {
	case "windows":
		path = strings.ReplaceAll(path, "/", "\\")
	}
	path = strings.ReplaceAll(path, "//", "/")
	path = strings.ReplaceAll(path, "\\\\", "\\")
	return path
}
func RealPath(path string) string {
	p, err := filepath.Abs(path)
	if err != nil {
		return ""
	}
	if !Exists(p) {
		return ""
	}
	return p
}

func PutBytes(file string, data []byte) error {
	dir := Dir(file)
	if !IsDir(dir) {
		if err := Mkdir(dir); err != nil {
			return err
		}
	}
	return os.WriteFile(file, data, os.ModePerm)
}

func Delete(path string) error {
	if len(path) == 0 {
		return nil
	}
	return os.RemoveAll(path)
}
