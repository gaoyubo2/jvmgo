package classpath

import "io/ioutil"
import "path/filepath"

type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry {
	//转化为绝对路径
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}

func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	//拼接类文件目录 和 类文件名
	// '/gyb/xxx/' + 'helloworld.class' = '/gyb/xxx/helloworld.class'
	fileName := filepath.Join(self.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}

func (self *DirEntry) String() string {
	return self.absDir
}
