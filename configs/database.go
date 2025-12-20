package configs

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var PoolDB *gorm.DB = nil
var SqlDB *sql.DB = nil

func OpenConnection() {
	var err error

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

	PoolDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})

	if err != nil {
		panic("Error")
	}

	SqlDB, err = PoolDB.DB()
	if err != nil {
		panic("Failed to connect")
	}

	if SqlDB.Ping() != nil {
		panic("Connection not alive")
	}

	SqlDB.SetMaxIdleConns(2)
	SqlDB.SetMaxOpenConns(5)
	SqlDB.SetConnMaxIdleTime(5 * time.Minute)
	SqlDB.SetConnMaxLifetime(1 * time.Hour)
	fmt.Println("Success make connection")
}

func GetDatabaseInstance() *gorm.DB {
	return PoolDB
}
