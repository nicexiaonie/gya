package gya

import (
	"fmt"
	"os"
	"path/filepath"
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

}

func Run() {
	basePath, _ := os.Getwd()
	println(basePath)
}
