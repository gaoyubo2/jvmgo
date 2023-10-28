package classpath

import "os"
import "path/filepath"
import "strings"

func newWildcardEntry(path string) CompositeEntry {
	//截取通用匹配符 /gyb/* 截取掉 *
	baseDir := path[:len(path)-1] // remove *
	//多个 类目录对象
	compositeEntry := []Entry{}

	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		//如果为空
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		//如果是 .jar  或者 .JAR 结尾的文件
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	}
	//遍历 目录下所有 .jar .JAR 文件 生成ZipEntry目录对象 放在切片中返回
	//walFn为函数
	filepath.Walk(baseDir, walkFn)

	return compositeEntry
}
