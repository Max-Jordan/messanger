// Package internal
package internal

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/Max-Jordan/messanger/internal/models"
	"github.com/Max-Jordan/messanger/internal/network"
)

type server struct {
	address string
	logger *slog.Logger
	mux *http.ServeMux
}

func NewServer(addr string, logger *slog.Logger, mux *http.ServeMux) *server {
	return &server{address: addr, logger: logger, mux: mux}
}

func (s *server) Start() error {
	s.mux.HandleFunc("POST /register", registerHandler)
	err := http.ListenAndServe(s.address, s.mux)
	if err != nil {
		return fmt.Errorf("starting server failed: %w", err)
	}
	return nil
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	var request network.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Incorrect JSON", http.StatusBadRequest)
		return
	}
	user := models.User{UserName: request.UserName, Password: request.Password}

	if err := json.NewEncoder(w).Encode(models.NewUserDTO(user)); err != nil {
		return
	}
}
