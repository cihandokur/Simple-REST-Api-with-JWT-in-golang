package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cihandokur/devlab/config"
	"github.com/cihandokur/devlab/db"
	"github.com/cihandokur/devlab/router"
)

func main() {

	err := config.LoadConfiguration()

	if err != nil {
		log.Panic(fmt.Errorf("failed to load Configs. %v", err))
	}

	db.New()

	registeredRouters := router.RegisterRoutes()

	log.Printf("Server is started and listening on : %s:%v", config.Config.Server.Host, config.Config.Server.Port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%v", config.Config.Server.Host, config.Config.Server.Port), registeredRouters))
}
