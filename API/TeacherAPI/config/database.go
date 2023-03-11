package config

import (
	"TeacherAPI/helper"
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
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		helper.ErrorHandler(err)
	}
	return db
}
