package main

import (
	"log"
	"web-api/handlers"
)

//	@title			Web API
//	@version		1.0
//	@description	API backend for the Web project

//	@host		localhost:5050
//	@BasePath	/api/v1
//	@accept		json

type config struct {
	port int
}

type application struct {
	config      config
	infoLog     *log.Logger
	errorLog    *log.Logger
	handlers    *handlers.Handlers
	environment string
}

func main() {
	app, err := initializeApp()
	if err != nil {
		log.Fatal(err)
	}

	err = app.serve()
	if err != nil {
		log.Fatal(err)
	}
}
