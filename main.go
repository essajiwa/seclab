package main

import (
	"database/sql"
	"log"
	"seclab/api"
	"seclab/repository"
	"seclab/service"

	_ "github.com/lib/pq"
)

func main() {
	// connect to postgres
	conn, err := sql.Open("postgres", "host=127.0.0.1 port=55432 user=postgres password=password dbname=lokal sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	if err := conn.Ping(); err != nil {
		log.Fatal(err)
	}

	repo := repository.NewRepository(conn)
	service := service.NewService(repo)
	handler := api.NewHandler(service)

	api.RunHttpServer(handler)
}
