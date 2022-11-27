package log

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	DefaultDirPerm = 0755
)

func getPwd() string {
	ex, err := os.Executable()
	if err != nil {
		return os.TempDir()
	}
	return filepath.Dir(ex)
}

// 创建目录，并返回日志模式
func GetLogFile(logPath string) (f string, err error) {
	dir := filepath.Dir(filepath.Join(getPwd(), logPath))
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		if err = os.MkdirAll(dir, DefaultDirPerm); err != nil {
			err = fmt.Errorf("create log dir err:%v", err)
			return
		}
	}
	f = filepath.Join(dir, filepath.Base(logPath))
	return
}
