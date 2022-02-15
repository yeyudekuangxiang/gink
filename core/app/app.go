package app

import (
	"github.com/yeyudekuangxiang/gink/internal/zap"
	"gopkg.in/ini.v1"
	"gorm.io/gorm"
	"net/http"
)

var (
	DB     = new(gorm.DB)
	Ini    *ini.File
	Logger = zap.DefaultLogger().Sugar()
	Server *http.Server
)
