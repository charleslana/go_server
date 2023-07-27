package handlers

import (
	"net/http"

	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{db}
}

func (h *Handler) GetRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API Go com go-chi e GORM!"))
}
