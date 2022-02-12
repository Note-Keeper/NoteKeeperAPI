package config

import (
	log "log"
	os "os"

	gin "github.com/gin-gonic/gin"
	dotenv "github.com/joho/godotenv"
)

var (
	Port               string
	Mode               string
	DB_ClusterEndpoint string
	DB_Username        string
	DB_Password        string
	DB_Collection      string
)

func Init() {
	err := dotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	Port = os.Getenv("PORT")
	Mode = os.Getenv("MODE")
	DB_ClusterEndpoint = os.Getenv("DB_ENDPOINT")
	DB_Username = os.Getenv("DB_USERNAME")
	DB_Password = os.Getenv("DB_PASSWORD")
	DB_Collection = os.Getenv("DB_COLLECTION")

	if Mode == "PROD" {
		gin.SetMode(gin.ReleaseMode)
	}
}
