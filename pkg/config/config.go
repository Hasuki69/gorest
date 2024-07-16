package config

import (
	"fmt"
	"gorest/internal/model/domain"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type (
	Server struct {
		Port int `mapstructure:"port"`
	}

	DB struct {
		Dsn string `mapstructure:"dsn"`
	}

	Gorm struct {
		PrepareStmt bool `mapstructure:"prepare_stmt"`
	}

	Config struct {
		Server Server `mapstructure:"server"`
		DB     DB     `mapstructure:"database"`
		Gorm   Gorm   `mapstructure:"gorm"`
	}
)

func LoadConfig() (*Config, error) {
	// Set Viper Configuration
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("pkg/config")

	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	// Unmarshal the config into the struct
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error unmarshalling config: %w", err)
	}

	// Return the config
	return &config, nil
}

func InitDB(cfg *Config) (*gorm.DB, error) {
	// Create GORM Configuration
	gormConfig := &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info),
		PrepareStmt: cfg.Gorm.PrepareStmt,
	}

	// Open Database Connection
	db, err := gorm.Open(postgres.Open(cfg.DB.Dsn), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// AutoMigrate the Database
	err = db.AutoMigrate(
		&domain.User{}, &domain.Session{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}
