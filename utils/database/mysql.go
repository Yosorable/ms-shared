package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// todo: 入参定义统一的结构体
func InitMysql(
	dbUser string,
	dbPassword string,
	dbAddr string,
	dbName string,
) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser,
		dbPassword,
		dbAddr,
		dbName,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return
}
