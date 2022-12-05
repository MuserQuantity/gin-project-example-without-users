package system

import (
	"github.com/MuserQuantity/gin-project-example-without-users/global"
	"go.uber.org/zap"
)

func LogError(err error) {
	if err != nil {
		global.SysLog.Error(err.Error())
	}
}

func LogInfo(msg string) {
	global.SysLog.Info(msg)
}

func LogDebug(msg string) {
	global.SysLog.Debug(msg)
}

func LogSystemError(sysErr *global.SystemError) {
	if sysErr != nil {
		global.SysLog.Error(sysErr.ErrMsg(), zap.Int("errCode", sysErr.ErrCode()))
	}
}
