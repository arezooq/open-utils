package main

import (
	"log"

	"open-utils/internal/db/connection"
	"open-utils/internal/errors"
	"open-utils/internal/logger"
)

func main() {
	// 1. Connection to DB
	conn := connection.ConnectDB()
	if conn == nil {
		log.Fatal("Could not connect to database")
	}

	log.Println("DB connection initialized")

	// 2. Logger
	log := logger.New("open-utils")
	reqID := logger.GenerateRequestID()
	log.Info(reqID, "Server started successfully")

	// 3. Example usage of errors
	err := errors.ErrUnauthorized
	log.Error(reqID, err.Error())

	// Block main to avoid exit
	select {}
}