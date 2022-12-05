package config

import "os"

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
	MongoUrl = GetEnvDefault("MONGO_URL", "mongodb://localhost:27017/")
	DatabaseName = GetEnvDefault("DATABASE_NAME", "shortener-link")
	Port = GetEnvDefault("PORT", "8000")
}
