package gya

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/nicexiaonie/gconf"
	"github.com/spf13/viper"
	"io/fs"
	"path"
	"path/filepath"
)

var config map[string]*viper.Viper

func getConfigPath() string {
	return basePath + "/" + configPath + "/"
}

func runConfig() {
	config = make(map[string]*viper.Viper)
	cpate := getConfigPath()

	fmt.Println("project config dir: " + cpate)
	err := filepath.Walk(cpate, func(filePath string, info fs.FileInfo, err error) error {
		ext := path.Ext(filePath)
		if ext == ".yaml" {

			fileName := path.Base(filePath)
			cName := fileName[:len(fileName)-len(ext)]
			//fmt.Println(cName[0:2])
			if cName[0:1] == "." {
				return nil
			}
			Config, err := gconf.New(gconf.Config{
				ConfigPath:  cpate,
				ConfigName:  cName,
				WatchConfig: true,
				CallOnConfigChange: func(in fsnotify.Event) {
					fmt.Println(fmt.Sprintf("权限配置监控变化: Name:%s, Op:%s, String:%s", in.Name, in.Op, in.String()))
				},
			})
			if err != nil {
				fmt.Println(fmt.Sprintf("读取配置: cName:%s, filePath:%s", cName, filePath))
				panic(err)
			}
			fmt.Println(fmt.Sprintf("读取配置: cName:%s, filePath:%s", cName, filePath))
			//fmt.Println(Config)
			config[cName] = Config
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func GetConfig(sign string) *viper.Viper {
	return config[sign]
}
