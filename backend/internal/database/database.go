package database

import (
    "fmt"
    "log"

    "github.com/anayy09/academiaflow-backend/configs"
    "github.com/anayy09/academiaflow-backend/internal/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect(config *configs.Config) {
    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=UTC",
        config.Database.Host,
        config.Database.User,
        config.Database.Password,
        config.Database.DBName,
        config.Database.Port,
        config.Database.SSLMode,
    )

    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    })

    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    log.Println("Database connected successfully!")
}

func Migrate() {
    err := DB.AutoMigrate(
        &models.User{},
        &models.Course{},
        &models.Assignment{},
    )

    if err != nil {
        log.Fatal("Failed to migrate database:", err)
    }

    log.Println("Database migration completed!")
}

func GetDB() *gorm.DB {
    return DB
}