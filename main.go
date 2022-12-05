package main

import (
	"github.com/GenesisBlock3301/url_shortner_golang/config"
	"github.com/GenesisBlock3301/url_shortner_golang/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func init() {
	initEnv()
	config.Connect()
}

func initEnv() {
	logger.Log{Message: "Loading environment settings"}.Info()
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("There is not exist env file")
		logger.Log{Message: "No local env file. using global OS environment variables."}.Info()
	}
	config.SetEnvironment()
}

func main() {
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Hello",
		})
	})
	HOST := os.Getenv("HOST")
	PORT := os.Getenv("PORT")
	err := router.Run(HOST + `:` + PORT)
	if err != nil {
		return
	}
}
