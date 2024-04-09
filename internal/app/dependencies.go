package app

import (
	"net/http"

	"github.com/hudayberdipolat/go-newsApp-backend/pkg/config"
	"github.com/hudayberdipolat/go-newsApp-backend/pkg/database/dbConfig"
	"github.com/hudayberdipolat/go-newsApp-backend/pkg/database/seeder"
	CustomHttp "github.com/hudayberdipolat/go-newsApp-backend/pkg/http"
	"gorm.io/gorm"
)

type Dependencies struct {
	Config     *config.Config
	DB         *gorm.DB
	HttpServer *http.Client
}

func GetDependencies() (*Dependencies, error) {

	getConfig, err := config.GetConfig()
	if err != nil {
		return nil, err
	}
	// database config
	newDBConfig := dbConfig.NewDBConnection(getConfig)
	db, err := newDBConfig.GetDBConnection()

	seed := seeder.NewSeeder(db)
	if errSeed := seed.Seeder(); errSeed != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	//http
	httpClient := CustomHttp.NewHttpClient()
	return &Dependencies{
		Config:     getConfig,
		DB:         db,
		HttpServer: httpClient,
	}, nil
}
