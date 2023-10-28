package classpath

import "errors"
import "strings"

type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}

	for _, path := range strings.Split(pathList, pathListSeparator) {
		//去判断 path 属于哪其他三种哪一种情况 生成对应的 ClassDirEntry类目录对象
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}

	return compositeEntry
}

func (self CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	//遍历切片 中的 类目录对象
	for _, entry := range self {
		//如果找到了 对应的 类 直接返回
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, nil
		}

	}
	//没找到 返回错误
	return nil, nil, errors.New("class not found: " + className)
}

func (self CompositeEntry) String() string {
	strs := make([]string, len(self))

	for i, entry := range self {
		strs[i] = entry.String()
	}

	return strings.Join(strs, pathListSeparator)
}
