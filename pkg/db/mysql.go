package db

import (
	"fmt"
	"github.com/exiaohao/golang-template/pkg/common"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

func init() {
	dialect := common.GetEnv("dialect", "mysql")

	dbHost := common.GetEnv("db_host", "")
	dbUsername := common.GetEnv("db_user", "root")
	dbPassword := common.GetEnv("db_password", "")
	dbDatabase := common.GetEnv("db_database", "example")

	// parseTime=True helps parse type=`timestamp` fields to time.Time, more details to be filled.
	dsn := fmt.Sprintf("%s:%s@%s/%s?charset=utf8mb4&parseTime=True", dbUsername, dbPassword, dbHost, dbDatabase)
	db, err = gorm.Open(dialect, dsn)
	if err != nil {
		glog.Fatalf("Connect to %s via %s failed, because: %s", dialect, dsn, err)
	}

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