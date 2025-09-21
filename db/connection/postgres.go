package connection

import (
	"fmt"
	"log"
	"time"

	"github.com/arezooq/open-utils/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

type DBConfig struct {
		Host string
		Port string
		User string
		Password string
		DBName string
		SSLMode string
}

func ConnectDB(cfg DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=UTC",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, errors.New("Failed to get db instance: %v", err.Error(), 500)
	}

	// Connection pool config
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	log.Println("Connected to PostgreSQL with GORM successfully")

	return db, nil
}
