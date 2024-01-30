package app

import (
	"github.com/hudayberdipolat/go-newsApp-backend/pkg/config"
	"github.com/hudayberdipolat/go-newsApp-backend/pkg/database/dbConfig"
	CustomHttp "github.com/hudayberdipolat/go-newsApp-backend/pkg/http"
	"gorm.io/gorm"
	"net/http"
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
