package handler

import (
	"encoding/json"
	"net/http"

	"spd-lookup/api/data"
)

// Handler is the struct for route handler functions
type Handler struct {
	db data.DatabaseInterface
}

// NewHandler is the constructor for the handler
func NewHandler() *Handler {
	return &Handler{
		db: data.NewClient(),
	}
}

// Ping pong :^)
func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("🏓 P O N G 🏓"))
}

// DescribeDepartments returns a list of departments and the fields supported for that department
func (h *Handler) DescribeDepartments(w http.ResponseWriter, r *http.Request) {
	departments := []*data.DepartmentMetadata{
		h.db.SeattleOfficerMetadata(),
		h.db.TacomaOfficerMetadata(),
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(departments)
}
