package conf

import (
	"gogofly/model"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitDB() (*gorm.DB, error) {
	logMode := logger.Info
	if !viper.GetBool("mode.develop") {
		logMode = logger.Error
	}

	db, err := gorm.Open(mysql.Open(viper.GetString("db.dsn")), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "sys_",
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logMode),
	})

	if err != nil {
		return nil, err
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(viper.GetInt("db.maxIdleConn"))
	sqlDB.SetMaxOpenConns(viper.GetInt("db.maxOpenConn"))
	sqlDB.SetConnMaxLifetime(time.Hour)

	// 将数据model映射到表    如果表存在，则无事发生，如果表不存在，则创建
	db.AutoMigrate(&model.User{})
	return db, nil
}
