package config

import (
	"fmt"
	"gorest/internal/model/domain"
	"gorm.io/driver/mysql"

	"github.com/spf13/viper"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type (
	Server struct {
		Port int `mapstructure:"port"`
	}

	DB struct {
		Dsn               string `mapstructure:"dsn"`
		DefaultStringSize uint   `mapstructure:"default_string_size"`
	}

	Gorm struct {
		PrepareStmt bool `mapstructure:"prepare_stmt"`
	}

	App struct {
		Path string `mapstructure:"path"`
		File string `mapstructure:"file"`
	}

	Config struct {
		Server Server `mapstructure:"server"`
		DB     DB     `mapstructure:"database"`
		Gorm   Gorm   `mapstructure:"gorm"`
		App    App    `mapstructure:"app"`
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
	db, err := gorm.Open(
		mysql.New(mysql.Config{
			DSN:               cfg.DB.Dsn,
			DefaultStringSize: cfg.DB.DefaultStringSize,
		}),
		gormConfig,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// AutoMigrate the Database
	if err = db.AutoMigrate(
		&domain.User{}, &domain.Session{},
	); err != nil {
		return nil, err
	}

	return db, nil
}
