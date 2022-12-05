package core

import (
	"fmt"
	"github.com/MuserQuantity/gin-project-example-without-users/service/system"
	"time"

	"github.com/MuserQuantity/gin-project-example-without-users/global"
	"github.com/MuserQuantity/gin-project-example-without-users/initialize"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	Router := initialize.Routers()
	// use cors
	//Router.Use(middleware.Cors())
	address := fmt.Sprintf("%s", global.SysConfig.System.GetAddr())
	s := initServer(address, Router)
	time.Sleep(10 * time.Microsecond)
	system.LogInfo("server run success on " + global.SysConfig.System.GetAddr())
	global.SysLog.Error(s.ListenAndServe().Error())
}
