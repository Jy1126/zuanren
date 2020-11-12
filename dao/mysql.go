package dao

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
	"xtpwidget/utils"
)

var (
	DB *gorm.DB
	db *sql.DB
)

func InitMySQL() (err error) {
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/joomla?charset=utf8&parseTime=True&loc=Local", utils.Config.User, utils.Config.Password, utils.Config.Host, utils.Config.Port)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	db, _ = DB.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	db.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	db.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	db.SetConnMaxLifetime(60 * time.Second)

	return err
}

func Close() {
	if err := db.Close(); err != nil {
		log.Fatal("db Close failed:", err)
	} else {
		log.Println("db Close succeed")
	}
}
