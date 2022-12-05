package internal

import (
	"fmt"
	"gorm.io/gorm/logger"

	"github.com/MuserQuantity/gin-project-example-without-users/global"
)

type writer struct {
	logger.Writer
}

// NewWriter creates a new writer
func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}

// Printf implements gorm logger interface
func (w *writer) Printf(message string, data ...interface{}) {
	var logZap bool
	logZap = global.SysConfig.Mysql.LogZap
	if logZap {
		global.SysLog.Info(fmt.Sprintf(message+"\n", data...))
	} else {
		w.Writer.Printf(message, data...)
	}
}
