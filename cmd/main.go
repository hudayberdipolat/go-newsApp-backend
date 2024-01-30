package main

import (
	"fmt"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/app"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/setup/constructor"
	"log"
)

func main() {
	getDependencies, err := app.GetDependencies()
	if err != nil {
		log.Fatal("error -->>", err.Error())
	}
	constructor.Build(getDependencies)
	appRouter := app.NewApp(getDependencies)
	runServer := fmt.Sprintf("%s:%s", getDependencies.Config.HttpConfig.ServerHost,
		getDependencies.Config.HttpConfig.ServerPort)
	if errRunServer := appRouter.Listen(runServer); errRunServer != nil {
		log.Fatal("run server error", errRunServer.Error())
	}
}
