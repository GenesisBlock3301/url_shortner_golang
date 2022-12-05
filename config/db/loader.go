package db

import (
	"fmt"
	"github.com/GenesisBlock3301/url_shortner_golang/logger"
	"os"
)

var (
	MongoUrl     = ""
	Port         = ""
	DatabaseName = ""
)

func GetEnvDefault(key, defVal string) string {
	val, exist := os.LookupEnv(key)
	if !exist {
		return defVal
	}
	return val
}

func SetEnvironment() {
	logger.Log{Message: "Environment loading"}.Info()
	fmt.Println("Databse connection loading...")
	MongoUrl = GetEnvDefault("MONGO_URL", "mongodb://localhost:27017/")
	DatabaseName = GetEnvDefault("DATABASE_NAME", "shortener-link")
	Port = GetEnvDefault("PORT", "8000")
}
