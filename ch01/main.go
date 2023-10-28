package main

import "fmt"

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		//模拟输出版本
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

// 模拟启动JVM
func startJVM(cmd *Cmd) {
	fmt.Printf("classpath:%s class:%s args:%v\n",
		cmd.cpOption, cmd.class, cmd.args)
}
