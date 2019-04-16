package db

import (
	"fmt"
	"github.com/exiaohao/golang-template/pkg/common"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var db *gorm.DB
var err error

func init() {
	dialect := common.GetEnv("dialect", "mysql")

	// dbHost need to filled as `tcp(127.0.0.1:3306)`
	dbHost := common.GetEnv("db_host", "")
	dbUsername := common.GetEnv("db_user", "root")
	dbPassword := common.GetEnv("db_password", "")
	dbDatabase := common.GetEnv("db_database", "example")

	// parseTime=True helps parse type=`timestamp` fields to time.Time, more details to be filled.
	// loc=Local use your (machine) local timezone, specific a timezone: loc=Asia%2FShanghai
	dsn := fmt.Sprintf("%s:%s@%s/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUsername, dbPassword, dbHost, dbDatabase)
	db, err = gorm.Open(dialect, dsn)
	if err != nil {
		glog.Fatalf("Connect to %s via %s failed, because: %s", dialect, dsn, err)
	}

	// Make sure max lifetime is less than your ALB/MySQL instace's interactive_timeout
	// to avoid 'Error: mysql is gone away'
	// use `show variables like '%timeout';` get current interactive_timeout settings (default is 8h)
	// make sure interactive_timeout is longer than ConnMaxLifetime
	db.DB().SetConnMaxLifetime(3600 * time.Second)
	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(100)

	// Auto Migration: Automatically migrate schema, keep schema update to date.
	// AutoMigrate will ONLY create tables, missing columns and missing indexes,
	// and WON'T change existing columns' type or delete unused columns to protect your data.
	// see more http://doc.gorm.io/database.html#migration
	// db.AutoMigrate(&model.Book{})
}

// GetDB return a initialized gorm.DB instance
func GetDB() *gorm.DB {
	return db
}

// CloseDB close db connection
func CloseDB() {
	db.Close()
}