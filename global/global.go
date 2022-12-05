package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
	"sync"

	"github.com/MuserQuantity/gin-project-example-without-users/config"
	"github.com/MuserQuantity/gin-project-example-without-users/utils/timer"
)

var (
	SysDb                 *gorm.DB
	SysRedis              *redis.Client
	SysVp                 *viper.Viper
	SysConfig             config.Server
	SysLog                *zap.Logger
	SysTimer              = timer.NewTimerTask()
	SysConcurrencyControl = &singleflight.Group{}
	lock                  sync.RWMutex
)
