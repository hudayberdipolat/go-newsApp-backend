package dbConfig

import (
	"fmt"

	"github.com/hudayberdipolat/go-newsApp-backend/pkg/config"
	_ "github.com/jackc/pgx/v4/stdlib" // For github.com/jackc/pgx/v4 driver
	_ "github.com/lib/pq"              // For lib/pq driver
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbConfig struct {
	config *config.Config
}

func NewDBConnection(config *config.Config) DbConfig {
	return DbConfig{
		config: config,
	}
}

func (dbConfig DbConfig) GetDBConnection() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbConfig.config.DBConfig.DbHost,
		dbConfig.config.DBConfig.DbUser,
		dbConfig.config.DBConfig.DbPassword,
		dbConfig.config.DBConfig.DbName,
		dbConfig.config.DBConfig.DbPort,
		dbConfig.config.DBConfig.DbSllMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})

	if err != nil {
		return nil, err
	}
	return db, err
}
