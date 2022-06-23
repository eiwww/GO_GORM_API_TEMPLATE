package main

import (
	"example/hello/database"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {

	database.ConnectDB()
	database.MigrateDB()
	database.SeedDB() //after run for first comment this seed

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Fail to load .env file")
	}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {

		fmt.Print("server is run on " + os.Getenv("PORT"))
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	var port string = "127.0.0.1:" + os.Getenv("PORT")
	r.Run(port)

}
