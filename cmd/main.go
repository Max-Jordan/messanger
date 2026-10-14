package main

import (
	"database/sql"
	"log/slog"
	"net/http"
	"os"

	"github.com/Max-Jordan/messanger/internal"
	_ "github.com/lib/pq"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	db, err := sql.Open("postgres", "dbname=messanger")
	if err != nil {
		logger.Error("Openning data base failed", "error", err)
	}
	defer db.Close()
	mux := http.NewServeMux()
	server := internal.NewServer(":8080", logger, mux)
	if err := server.Start(); err != nil {
		logger.Error("Server stopped", "error", err)
	}
}
