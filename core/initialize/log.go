package initialize

import (
	"github.com/yeyudekuangxiang/gink/core/app"
	"github.com/yeyudekuangxiang/gink/internal/zap"
	"log"
)

func InitLog() {
	var loggerConfig zap.LoggerConfig
	var err error
	err = app.Ini.Section("log").MapTo(&loggerConfig)
	if err != nil {
		log.Fatal(err)
	}
	loggerConfig.Path = "runtime"
	loggerConfig.FileName = "log.log"
	app.Logger = zap.NewZapLogger(loggerConfig).Sugar()
}
