package global

import (
	"github.com/go-ecommerce/pkg/logger"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ConfigStruct struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Mysql struct {
		Host            string `mapstructure:"host"`
		Port            int    `mapstructure:"port"`
		Username        string `mapstructure:"username"`
		Password        string `mapstructure:"password"`
		Dbname          string `mapstructure:"dbname"`
		MaxIdleConns    int    `mapstructure:"maxIdleConns"`
		ConnMaxLifetime int    `mapstructure:"connMaxLifetime"`
		MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	} `mapstructure:"mysql"`
}

var (
	Config ConfigStruct
	Logger *logger.LoggerZap
	Mdb *gorm.DB
	Rdb *redis.Client
)
