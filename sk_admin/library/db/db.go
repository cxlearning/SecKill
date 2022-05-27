package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"myProject/SecKill/sk_admin/conf"

	"fmt"
	"time"
)

var db *gorm.DB

func GetDBInstance() *gorm.DB {
	return db
}

func Init() {
	var err error
	// 默认设置了隔离级别为RC，以避免间隙锁导致的死锁。
	// 间隙锁可以避免幻读，如果担心幻读，可以使用默认的隔离级别RR，把&tx_isolation=%%27READ-COMMITTED%%27去掉即可
	dsn := fmt.Sprintf("%s@%s/%s?charset=utf8&parseTime=True&loc=Local&timeout=3s&tx_isolation=%%27READ-COMMITTED%%27",
		conf.Config.Database.UserPassword, conf.Config.Database.HostPort, conf.Config.Database.DB)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println(dsn)
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetConnMaxLifetime(time.Duration(conf.Config.Database.Conn.MaxLifeTime) * time.Second)
	sqlDB.SetMaxIdleConns(conf.Config.Database.Conn.MaxIdle)
	sqlDB.SetMaxOpenConns(conf.Config.Database.Conn.MaxOpen)
}
