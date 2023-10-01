package config

import (
	"github.com/satriohtmo/go-gin-gorm.git/config/app_config"
	"github.com/satriohtmo/go-gin-gorm.git/config/database_config"
)

func InitConfig() {
	app_config.InitAppConfig()
	database_config.InitDatabaseConfig()
}