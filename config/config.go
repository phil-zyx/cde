package config

import (
	"fmt"
	"os"

	"github.com/cde/util/logger"
	"github.com/spf13/viper"
)

// Init .
func Init() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("function ID %v", err))
	}
	logFile := viper.GetString("logFile")
	err = logger.Init(logFile)
	if err != nil {
		panic(fmt.Sprintf("读取YML配置 %v", err))
	}
}
