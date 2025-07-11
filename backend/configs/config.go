package configs

import (
    "log"
    "os"
    "strconv"
    "time"

    "github.com/joho/godotenv"
)

type Config struct {
    Database DatabaseConfig
    JWT      JWTConfig
    Server   ServerConfig
}

type DatabaseConfig struct {
    Host     string
    Port     int
    User     string
    Password string
    DBName   string
    SSLMode  string
}

type JWTConfig struct {
    Secret    string
    ExpiresIn time.Duration
}

type ServerConfig struct {
    Port string
    Host string
}

func LoadConfig() *Config {
    if err := godotenv.Load(); err != nil {
        log.Printf("No .env file found")
    }

    port, _ := strconv.Atoi(getEnv("DB_PORT", "5432"))
    expiresIn, _ := time.ParseDuration(getEnv("JWT_EXPIRES_IN", "24h"))

    return &Config{
        Database: DatabaseConfig{
            Host:     getEnv("DB_HOST", "localhost"),
            Port:     port,
            User:     getEnv("DB_USER", "postgres"),
            Password: getEnv("DB_PASSWORD", ""),
            DBName:   getEnv("DB_NAME", "academiaflow"),
            SSLMode:  getEnv("DB_SSLMODE", "disable"),
        },
        JWT: JWTConfig{
            Secret:    getEnv("JWT_SECRET", "your-secret-key"),
            ExpiresIn: expiresIn,
        },
        Server: ServerConfig{
            Port: getEnv("SERVER_PORT", "8080"),
            Host: getEnv("SERVER_HOST", "localhost"),
        },
    }
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}