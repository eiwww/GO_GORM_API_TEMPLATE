package database

import (
	"context"
	"fmt"
	"log"
	"os"
	model "rest/model"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

type SqlLogger struct {
	logger.Interface
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	fmt.Printf("%v\n=================================\n", sql)
}

func ConnectDB() error {
	if db != nil {
		CloseDBConnection(db)
	}

	err := godotenv.Load()
	if err != nil {
		panic("Fail to load env file")
	}

	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASS")
	dbHost := os.Getenv("DBHOST")
	dbName := os.Getenv("DBNAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&loc=Local&parseTime=True", dbUser, dbPass, dbHost, dbName)
	database, er := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: &SqlLogger{},
		// DryRun: true,
		//if is true is mean only show sql but not excute, use to check if we make right sql
	})
	if er != nil {
		log.Fatal(er)
		panic("Fail to connect database")
	}

	db = database

	return er
}

func GetDatabaseConnection() (*gorm.DB, error) {
	sqlDB, err := db.DB()
	if err != nil {
		return db, err
	}
	if err := sqlDB.Ping(); err != nil {
		return db, err
	}
	return db, nil
}

func CloseDBConnection(conn *gorm.DB) {
	sqlDB, err := conn.DB()
	if err != nil {
		log.Fatal("Error occurred while closing a DB connection")
	}
	defer sqlDB.Close()
}

func MigrateDB() {
	db.AutoMigrate(model.Province{}, model.District{})
}
