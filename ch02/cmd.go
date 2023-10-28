package main

import (
	"flag"
	"fmt"
	"os"
)

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
	// jre路径
	XjreOption string
}

func parseCmd() *Cmd {
	cmd := &Cmd{}

	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		//第一个参数是主类名
		cmd.class = args[0]
		cmd.args = args[1:]
	}

	return cmd
}
func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
	//flag.PrintDefaults()
}
