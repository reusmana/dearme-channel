package config

import (
	"fmt"
	"os"

	// _ "github.com/cengsin/oracle"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

var err error

// var dbOrcl *gorm.DB

// var errOrcl error

func InitializeDB() {

	dbdriver := os.Getenv("DB_DRIVER_DEARME")
	host := os.Getenv("DB_HOST_DEARME")
	port := os.Getenv("DB_PORT_DEARME")
	user := os.Getenv("DB_USERNAME_DEARME")
	password := os.Getenv("DB_PASSWORD_DEARME")
	dbname := os.Getenv("DB_NAME_DEARME")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = gorm.Open(dbdriver, connectionString)
	if err != nil {
		panic("connectionString error")
	}

	err = db.DB().Ping()
	if err != nil {
		panic("dsn invalid")
	}

}

func ConnectionPg() *gorm.DB {
	return db
}

// func ConnectionOrcl() *gorm.DB {

// 	dbdriver := os.Getenv("DB_DRIVER_MARIS")
// 	host := os.Getenv("DB_HOST_MARIS")
// 	port := os.Getenv("DB_PORT_MARIS")
// 	user := os.Getenv("DB_USERNAME_MARIS")
// 	password := os.Getenv("DB_PASSWORD_MARIS")
// 	dbname := os.Getenv("DB_SERVICE_NAME_MARIS")

// 	connectionString := fmt.Sprintf(`user="%s" password="%s" connectString="%s:%s/%s"`, user, password, host, port, dbname)

// 	dbOrcl, errOrcl = gorm.Open(dbdriver, connectionString)

// 	if errOrcl != nil {
// 		panic("connectionString error Oracle")
// 	}

// 	errOrcl = dbOrcl.DB().Ping()
// 	if errOrcl != nil {
// 		panic("dsn invalid")
// 	}

// 	return dbOrcl
// }
