package main

import (
	"os"
	"runtime"

	"github.com/astaxie/beego"
)

const (
	APP_VER = "0.0.0.1"
)

func init() {
	runtime.GOMACPROCS(runtime.NumCPU())
}

func main() {
	beego.AppName = "Soul wifi portal Server"
}
