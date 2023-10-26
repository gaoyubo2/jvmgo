package ch01

type Cmd struct {
	// 标注是否为 --help
	helpFlag bool
	//标注是否为 --version
	versionFlag bool
	//选项
	cpOption string
	//主类名，或者是jar文件
	class string
	//参数
	args []string
}
