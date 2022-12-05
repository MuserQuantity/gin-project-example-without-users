package initialize

import (
	"github.com/MuserQuantity/gin-project-example-without-users/global"
)

func Timer() {
	// start timer
	if global.SysConfig.Timer.Start {
		// add your timer tasks(from initialize/task.go) here
	}
}
