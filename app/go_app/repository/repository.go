package repository

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type RepositoryConfig struct {
	DBHost         string
	DBUserName     string
	DBUserPassword string
	DBName         string
	DBPort         string
}

type Repository struct {
	Connection *sql.DB
}

type RepositoryInterface interface {
	Connect(config RepositoryConfig)
	Close()
	Exec(query string, args ...any) (sql.Result, error)
	QueryRow(query string, args ...any) *sql.Row
}

func (repo *Repository) Connect(config RepositoryConfig) {

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUserName, config.DBUserPassword, config.DBName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	repo.Connection = db

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database Connected")
}

func (repo *Repository) Close() {
	repo.Connection.Close()
}
func (r *Repository) Exec(query string, args ...any) (sql.Result, error) {
	return r.Connection.Exec(query, args...)
}
func (r *Repository) QueryRow(query string, args ...any) *sql.Row {
	return r.Connection.QueryRow(query, args...)
}
