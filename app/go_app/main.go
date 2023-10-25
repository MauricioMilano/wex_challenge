package main

import (
	"os"
	"prj/repository"
	"prj/server"
	"time"
)

func main() {
	time.Sleep(60 * time.Second)

	db := repository.RepositoryConfig{
		DBHost:         os.Getenv("APP_POSTGRES_HOST"),
		DBPort:         os.Getenv("APP_POSTGRES_PORT"),
		DBName:         os.Getenv("APP_POSTGRES_DB"),
		DBUserName:     os.Getenv("APP_POSTGRES_USER"),
		DBUserPassword: os.Getenv("APP_POSTGRES_PASS"),
	}

	repo := &repository.Repository{}
	repo.Connect(db)
	port := os.Getenv("APP_PORT")
	server.StartServer(port, repo)
}
