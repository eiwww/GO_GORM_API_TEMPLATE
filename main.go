package main

import (
	"log"
	"os"

	route "rest/api/v1/route"
	"rest/database"

	"rest/seeder"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {

	database.ConnectDB()
	database.MigrateDB()
	seeder.SeedDB()

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Fail to load .env file")
	}

	app := gin.New()

	router := app.Group("/api/v1")
	route.AddRoutes(router)

	var port string = "127.0.0.1:" + os.Getenv("PORT")
	app.Run(port)

}
