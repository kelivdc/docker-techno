package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"techno.com/models"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func SetupDatabase() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Singapore",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  false,         // Disable color
		},
	)
	data, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
		Logger:               newLogger,
		DisableAutomaticPing: true,
	})
	if err != nil {
		log.Fatal("Failed to connect to database.\n", err)
		os.Exit(2)
	}
	// db.Exec("CREATE EXTENSION IF NOT EXISTS postgis")
	data.Exec("ALTER DATABASE techno SET timezone TO 'Asia/Jakarta';")
	data.AutoMigrate(
		&models.Admin{},
	)
	fmt.Println("> Success database migration")
	DB = Dbinstance{
		Db: data,
	}
}
