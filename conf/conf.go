package conf

import (
	"path/filepath"
)

func BasePath(path string) string {
	return filepath.Join("C:\\Users\\Administrator\\Myads", path)
}
