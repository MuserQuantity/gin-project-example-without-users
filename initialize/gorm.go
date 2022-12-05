package initialize

import (
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"

	"github.com/MuserQuantity/gin-project-example-without-users/global"
	"github.com/MuserQuantity/gin-project-example-without-users/initialize/internal"
)

func Mysql() *gorm.DB {
	m := global.SysConfig.Mysql
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string default size
		SkipInitializeWithVersion: false,   // autoconfigure based on currently MySQL version
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config()); err != nil {
		panic(err)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
	// add your tables here

	)
	if err != nil {
		global.SysLog.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.SysLog.Info("register table success")
}
