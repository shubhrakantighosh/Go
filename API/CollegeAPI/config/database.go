package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	host         = "localhost"
	port         = "3306"
	user         = "root"
	password     = "shubhrakanti"
	databaseName = "db2"
)

func DataBaseConnection() *gorm.DB {

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + databaseName
	db, _ := gorm.Open(mysql.Open(dsn))
	return db
}
