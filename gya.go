package gya

import (
	"os"
	"path/filepath"
	"runtime"
)

type GyaOption struct {
	StartUpDirMod string // 启动目录模式
}

var env string
var basePath string
var configPath string = "config"

func init() {
	println("-------------")
}
func Init(option GyaOption) {

	// 确定目录
	if option.StartUpDirMod == "EXEC" {
		// 当前可执行文件所在的目录
		basePath = getExePath()
	} else if option.StartUpDirMod == "Caller" {
		// 源文件位置
		basePath = getCallerPath()
	} else if option.StartUpDirMod == "Current" {
		// 当前目录
		basePath = getCurrentPath()
	} else {
		basePath = getCurrentPath()
	}

	// 加载配置
	runConfig()
}

func getCurrentPath() string {
	exePath, _ := os.Getwd()
	return exePath
}

func getExePath() string {
	exePath, err := os.Executable()
	if err != nil {
		return ""
	}
	// 获取当前可执行文件所在的目录
	dir := filepath.Dir(exePath)
	return dir
}

func getCallerPath() string {
	// 获取当前源文件的路径
	_, filename, _, ok := runtime.Caller(2)
	if !ok {
		return ""
	}
	// 获取当前源文件所在的目录
	dir := filepath.Dir(filename)
	return dir
}

func Run() {
	basePath, _ := os.Getwd()
	println(basePath)
}
