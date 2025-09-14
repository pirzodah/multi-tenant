package main

import (
	"log"
	"net/http"

	"github.com/pirzodah/multi-tenant/pkg/config"
	"github.com/pirzodah/multi-tenant/pkg/db"
)

func main() {
	cfg := config.LoadConfig()

	// Подключение к БД
	database, err := db.NewPostgres(cfg.DBUrl)
	if err != nil {
		log.Fatal("Database connection failed: ", err)
	}
	defer database.Pool.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	log.Println("API server running on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
