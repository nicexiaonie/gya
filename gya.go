package gya

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func Init() {

	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("获取可执行文件路径失败:", err)
		return
	}

	// 获取当前可执行文件所在的目录
	dir := filepath.Dir(exePath)
	fmt.Println("当前文件所在目录:", dir)

	// 获取当前源文件的路径
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("获取源文件路径失败")
		return
	}

	// 获取当前源文件所在的目录
	dir = filepath.Dir(filename)
	fmt.Println("当前源文件所在目录:", dir)

}

func Run() {
	basePath, _ := os.Getwd()
	println(basePath)
}
