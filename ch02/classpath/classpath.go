package classpath

import "os"
import "path/filepath"

type Classpath struct {
	BootClasspath Entry
	ExtClasspath  Entry
	UserClasspath Entry
}

func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)

	// 拼接成jre 的路径
	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	//加载 所有底下的 jar包
	self.BootClasspath = newWildcardEntry(jreLibPath)

	// 拼接 扩展类 的路径
	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	//加载 所有底下的jar包
	self.ExtClasspath = newWildcardEntry(jreExtPath)
}

func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("Can not find jre folder!")
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (self *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	self.UserClasspath = newEntry(cpOption)
}

// className: fully/qualified/ClassName
func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := self.BootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := self.ExtClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	return self.UserClasspath.readClass(className)
}

func (self *Classpath) String() string {
	return self.UserClasspath.String()
}
