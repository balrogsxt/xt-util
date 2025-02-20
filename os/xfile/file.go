package xfile

import (
	"github.com/balrogsxt/xt-util/encoding/xjson"
	"github.com/balrogsxt/xt-util/valid"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// IsDir 判断文件夹是否存在
func IsDir(path string) bool {
	info, err := Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}
func Stat(path string) (os.FileInfo, error) {
	return os.Stat(path)
}

// IsFile 判断文件是否存在
func IsFile(path string) bool {
	info, err := Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}
func Exists(path string) bool {
	if stat, err := Stat(path); stat != nil && !os.IsNotExist(err) {
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
func Base(path string) string {
	return filepath.Base(path)
}
func Ext(path string) string {
	return filepath.Ext(path)
}

// ScanFiles 获取目录下所有文件
// Deprecated: 此函数已弃用,换xfile.ScanFile替代
func ScanFiles(path string, pattern string, callback func(path string, d fs.DirEntry) bool, isRecursives ...bool) ([]string, error) {
	isRecursive := false
	if len(isRecursives) > 0 {
		isRecursive = isRecursives[0]
	}
	files := make([]string, 0)
	p := filepath.Dir(path)
	if err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !isRecursive { //跳过子目录
			if d.IsDir() && filepath.Dir(path) != filepath.Clean(p) {
				return filepath.SkipDir
			}
		}
		if !d.IsDir() {
			isAppend := true
			if !valid.IsEmpty(pattern) {
				if matched, _ := filepath.Match(pattern, Base(path)); matched {
					if callback != nil {
						isAppend = callback(path, d)
					} else {
						isAppend = false
					}
				}
			}
			if isAppend {
				files = append(files, path)
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return files, nil
}

func Move(oldPath, newPath string) error {
	return os.Rename(oldPath, newPath)
}

func RemoveFile(path string) error {
	return os.Remove(path)
}
func RemoveAll(path string) error {
	return os.RemoveAll(path)
}

// CopyDir 复制文件夹下的所有文件
func CopyDir(src string, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		dstPath := filepath.Join(dst, relPath)
		if info.IsDir() {
			return os.MkdirAll(dstPath, info.Mode())
		}
		return CopyFile(path, dstPath)
	})
}

// CopyFile 复制文件
func CopyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}
	info, err := srcFile.Stat()
	if err != nil {
		return err
	}
	return dstFile.Chmod(info.Mode())
}

// GetFileToJson 读取文件转为JSON结构体
func GetFileToJson(file string, v any) error {
	res, err := GetBytes(file)
	if err != nil {
		return err
	}
	return xjson.Unmarshal(res, v)
}

// PutContentJson 格式化写入JSON格式
func PutContentJson(file string, v any) error {
	res, err := xjson.MarshalByte(v)
	if err != nil {
		return err
	}
	return PutBytes(file, res)
}
