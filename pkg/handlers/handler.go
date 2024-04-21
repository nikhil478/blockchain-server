package handlers

import (
	"fmt"
	"net/http"

	"github.com/FalconX80/blockchain-server/pkg/models"
	"gorm.io/gorm"
)

// Handler struct holds required services for handler to function
type Handler struct {
    DB *gorm.DB
}

// NewHandler creates a new Handler
func NewHandler(db *gorm.DB) *Handler {
    return &Handler{DB: db}
}

// CreateUser handles the user creation
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
    // Logic to parse user from request and create the user
    var user models.User
    // Assuming user is parsed correctly and validated
    if result := h.DB.Create(&user); result.Error != nil {
        http.Error(w, "Error creating user", http.StatusInternalServerError)
        return
    }
    fmt.Fprintf(w, "User created")
}

// GetUser retrieves a user by ID
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
    // Logic to retrieve user
}
