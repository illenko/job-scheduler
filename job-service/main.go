package main

import (
	"context"
	"database/sql"
	"job-service/internal/handler"
	"job-service/internal/repository"
	"job-service/internal/service"
	"log"
	"net/http"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := goose.Up(db, "migrations"); err != nil {
		log.Fatal(err)
	}

	dbpool, err := pgxpool.Connect(context.Background(), "postgres://postgres:postgres@localhost:5432/postgres")
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.Close()

	jobRepo := repository.NewJobRepository(dbpool)
	jobService := service.NewJobService(jobRepo)
	jobHandler := handler.NewJobHandler(jobService)

	http.HandleFunc("POST /jobs", jobHandler.CreateJob)
	http.HandleFunc("GET /jobs", jobHandler.GetJobs)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
