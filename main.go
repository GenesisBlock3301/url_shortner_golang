package main

import (
	"github.com/GenesisBlock3301/url_shortner_golang/config/db"
	"github.com/GenesisBlock3301/url_shortner_golang/logger"
	"github.com/GenesisBlock3301/url_shortner_golang/route"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func init() {
	initEnv()
	db.ConnectDB()
}

func initEnv() {
	logger.Log{Message: "Loading environment settings"}.Info()
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("There is not exist env file")
		logger.Log{Message: "No local env file. using global OS environment variables."}.Info()
	}
	db.SetEnvironment()
}

func main() {
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Hello",
		})
	})
	route.RootRouter(router)
	HOST := os.Getenv("HOST")
	PORT := os.Getenv("PORT")
	err := router.Run(HOST + `:` + PORT)
	if err != nil {
		return
	}
}
