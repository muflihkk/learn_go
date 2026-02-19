package main

import (
	"fmt"
	"go_test/src/config"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file");
	}

	config.ConnectDB()
	
	port:= os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": true,
			"message": "server is running ok done.",
		})
	})

	fmt.Println("server is running http://localhost:",port)
	r.Run(":" + port)
}