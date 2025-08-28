package config

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Database struct {
	DB *sql.DB
}

func NewDatabase(config *viper.Viper) (*Database, error) {
	// dsn := "bisma:bisma@tcp(127.0.0.1:4000)/main_database?parseTime=true"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.GetString("database.username"),
		config.GetString("datbase.password"),
		config.GetString("database.host"),
		config.GetString("database.port"),
		config.GetString("database.db_name"),
	)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, fmt.Errorf("failed to open database, %w", err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return &Database{DB: db}, nil
}
