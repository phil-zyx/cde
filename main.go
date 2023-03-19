package main

import (
	"github.com/cde/config"
	"github.com/cde/database"
	"github.com/cde/router"
	"github.com/cde/util/logger"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initGin() {
	r := router.InitRouter()
	config.RunProf(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	port := viper.GetString("port")
	if port != "" {
		if err := r.Run(":" + port); err != nil {
			logger.Fatalf("Service startup failed！err %v", err)
		}
	}
	err := r.Run()
	if err != nil {
		logger.Fatalf("Service startup failed！err %v", err)
	}
}

func main() {
	// 配置初始化
	config.Init()
	// 初始化数据库
	database.Init()
	// 初始化 gin
	initGin()
}
