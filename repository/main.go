package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/mrspec7er/mailserv/repository/internal"
	"github.com/mrspec7er/mailserv/repository/internal/database"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	DB := database.StartConnection()

	config := &internal.Server{
		DB: DB,
	}

	server := internal.NewServer(*config)

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}

}
