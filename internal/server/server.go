package server

import (
	"github.com/gin-gonic/gin"
	"github.com/wujunyi792/flatten-winner/config"
	_ "github.com/wujunyi792/flatten-winner/internal/controller/config"
	_ "github.com/wujunyi792/flatten-winner/internal/controller/file"
	_ "github.com/wujunyi792/flatten-winner/internal/corn"
	"github.com/wujunyi792/flatten-winner/internal/logger"
	"github.com/wujunyi792/flatten-winner/internal/middleware"
	"github.com/wujunyi792/flatten-winner/internal/redis"
	v1 "github.com/wujunyi792/flatten-winner/internal/router/v1"
)

var E *gin.Engine

func init() {
	logger.Info.Println("start init gin")
	if config.GetConfig().PORT == "" {
		panic("Port not set, please check config")
	}
	gin.SetMode(config.GetConfig().MODE)
	E = gin.New()
	E.Use(middleware.GinRequestLog, gin.Recovery(), middleware.Cors(E))
}

func Run() {
	redis.GetRedis()
	v1.MainRouter(E)
	if err := E.Run("0.0.0.0:" + config.GetConfig().PORT); err != nil {
		logger.Error.Fatalln(err)
	}
}
