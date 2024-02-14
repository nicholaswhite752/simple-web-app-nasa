package main

import (
	"log"
	"os"
	"strconv"
	"web-api/handlers"
)

func initializeApp() (*application, error) {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		infoLog.Println("failed to parse port, using 5050")
		port = 5050
	}

	cfg := config{
		port: port,
	}
	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	handlers := handlers.NewHandlers(infoLog, errorLog, env)

	a := &application{
		config:      cfg,
		infoLog:     infoLog,
		errorLog:    errorLog,
		handlers:    handlers,
		environment: env,
	}

	return a, nil
}
