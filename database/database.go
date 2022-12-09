package database

import (
	"fmt"
	"goblog/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	db  *gorm.DB
	err error
)

func InitDatabase() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.GlobalConfig.User,
		config.GlobalConfig.Passwd,
		config.GlobalConfig.Host,
		config.GlobalConfig.Port,
		config.GlobalConfig.Name)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("打开数据库失败")
	}

	//db.AutoMigrate(model.)
	rdb, err := db.DB()
	if err != nil {
		return
	}
	rdb.SetMaxIdleConns(10)  // 最大闲置连接数
	rdb.SetMaxOpenConns(100) // 最大连接数量
	rdb.SetConnMaxLifetime(10 * time.Second)
}
