package config

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewDatabase(config *viper.Viper) (*sql.DB, error) {
	// dsn := "bisma:bisma@tcp(127.0.0.1:4000)/main_database?parseTime=true"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?parseTime=true",
		config.GetString("database.username"),
		config.GetString("database.password"),
		config.GetString("database.host"),
		config.GetInt("database.port"),
		config.GetString("database.db_name"),
	)
	logrus.Infof("dsn: %s", dsn)
	db, err := sql.Open("mysql", dsn)
	logrus.Info("masuk sini")
	if err != nil {
		logrus.Errorf("failed to open database: %v", err)
		return nil, fmt.Errorf("failed to open database, %w", err)
	}
	logrus.Info("ga masuk sini")
	if err := db.Ping(); err != nil {
		logrus.Errorf("failed to connect database: %v", err)
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db, nil
}
