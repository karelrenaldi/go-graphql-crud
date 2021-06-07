package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	// DB is reusable gorm sql connection.
	DB *gorm.DB
)

// ConnectDB connects this application to database instance.
func ConnectDB() error {
	// TODO: change with environtment variable
	// h := "ktp-db"
	h := "localhost"
	u := "root"
	pwd := "root"
	// p := "3306"
	p := "5000"
	d := "ktp_db"

	dsn := u + ":" + pwd + "@tcp(" + h + ":" + p + ")/" + d + "?charset=utf8mb4&parseTime=True&loc=Local"

	dbConnection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = dbConnection
	return nil
}
