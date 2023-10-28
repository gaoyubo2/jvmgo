package classpath

import (
	"os"
	"strings"
)

// :(linux/unix) or ;(windows)
const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	// className: fully/qualified/ClassName.class
	readClass(className string) ([]byte, Entry, error)
	String() string
}

func newEntry(path string) Entry {
	//如果路径包含分隔符 表示有多个文件
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	//如果路径 后缀名有*匹配模式
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	//如果路径 是jar zip等压缩包形式
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {

		return newZipEntry(path)
	}

	return newDirEntry(path)
}
