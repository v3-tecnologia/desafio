package config

import (
    "log"
    "os"
)

type ConfigStruct struct {
    ServerPort   string
    DatabaseURL  string
    AWSRegion    string
    AWSAccessKey string
    AWSSecretKey string
}

var Config ConfigStruct

func LoadConfig() {
    Config = ConfigStruct{
        ServerPort:   getEnv("SERVER_PORT", "8080"),
        DatabaseURL:  getEnv("DATABASE_URL", "mongodb://localhost:27017/telemetry"),
        AWSRegion:    getEnv("AWS_REGION", "us-east-1"),
        AWSAccessKey: getEnv("AWS_ACCESS_KEY", ""),
        AWSSecretKey: getEnv("AWS_SECRET_KEY", ""),
    }

    if Config.AWSAccessKey == "" || Config.AWSSecretKey == "" {
        log.Fatal("AWS credentials are not set in environment variables.")
    }
}

func getEnv(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return fallback
}
