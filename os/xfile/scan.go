package xfile

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type ScanOption struct {
	Patterns  []string                 //文件匹配条件
	Recursive bool                     //是否递归查询
	Handler   func(path string) string //闭包处理
}

func ScanDir(path string, pattern string, recursives ...bool) ([]string, error) {
	recursive := false
	if len(recursives) > 0 {
		recursive = recursives[0]
	}
	files, err := doScanDir(path, ScanOption{
		Recursive: recursive,
		Handler:   nil,
		Patterns:  strings.Split(pattern, ","),
	})
	if err != nil {
		return nil, err
	}
	if len(files) > 0 {
		sort.Strings(files)
	}
	return files, nil
}

// ScanFile 获取目录下所有文件(不包含目录)
func ScanFile(path string, pattern string, recursives ...bool) ([]string, error) {
	recursive := false
	if len(recursives) > 0 {
		recursive = recursives[0]
	}
	files, err := doScanDir(path, ScanOption{
		Recursive: recursive,
		Handler: func(path string) string {
			if IsDir(path) {
				return ""
			}
			return path
		},
		Patterns: strings.Split(pattern, ","),
	})
	if err != nil {
		return nil, err
	}
	if len(files) > 0 {
		sort.Strings(files)
	}
	return files, nil
}

// ScanDirConfig 自定义扫描目录
func ScanDirConfig(path string, config ScanOption) ([]string, error) {
	files, err := doScanDir(path, config)
	if err != nil {
		return nil, err
	}
	if len(files) > 0 {
		sort.Strings(files)
	}
	return files, nil
}

// doScanDir 读取文件夹
func doScanDir(path string, config ScanOption) ([]string, error) {
	var (
		list      []string
		file, err = os.Open(path)
	)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	names, err := file.Readdirnames(-1)
	if err != nil {
		return nil, err
	}
	var filePath string
	for _, name := range names {
		filePath = path + string(filepath.Separator) + name
		if IsDir(filePath) && config.Recursive {
			array, _ := doScanDir(filePath, config)
			if len(array) > 0 {
				list = append(list, array...)
			}
		}
		if config.Handler != nil {
			filePath = config.Handler(filePath)
			if filePath == "" {
				continue
			}
		}
		if config.Patterns != nil {
			for _, p := range config.Patterns {
				if m, _ := filepath.Match(p, name); m {
					if filePath = Abs(filePath); filePath != "" {
						list = append(list, filePath)
					}
				}
			}
		}
	}
	return list, nil
}
