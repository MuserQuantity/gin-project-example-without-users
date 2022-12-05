package main

import (
	"github.com/MuserQuantity/gin-project-example-without-users/core"
	"github.com/MuserQuantity/gin-project-example-without-users/global"
	"github.com/MuserQuantity/gin-project-example-without-users/initialize"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	global.SysVp = core.Viper("config.yaml") // config file
	global.SysLog = core.Zap()               // log
	global.SysDb = initialize.Mysql()        // mysql database
	global.SysRedis = initialize.Redis()     // redis database
	if global.SysDb != nil {
		initialize.RegisterTables(global.SysDb) // initialize tables
		// close database connection before program ends
		db, _ := global.SysDb.DB()
		defer db.Close()
	}
	initialize.Timer() // timer
	core.RunServer()
}
