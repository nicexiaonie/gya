package gya

import "os"

func Init() {

}

func Run() {
    basePath, _ := os.Getwd()
    println(basePath)
}
