package initialize

import (
	"context"
	"github.com/MuserQuantity/gin-project-example-without-users/global"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func Redis() *redis.Client {
	redisCfg := global.SysConfig.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.GetRedisAddr(),
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.SysLog.Error("redis connect ping failed, err:", zap.Error(err))
		panic(err)
	} else {
		global.SysLog.Info("redis connect ping response:", zap.String("pong", pong))
		return client
	}
}
