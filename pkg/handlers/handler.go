package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

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

    var user models.User

    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Error parsing user data", http.StatusBadRequest)
        return
    }

    emailParts := strings.Split(user.Email, "@")
    if len(emailParts) < 2 {
        http.Error(w, "Invalid email format", http.StatusBadRequest)
        return
    }
    usernamePart := emailParts[0]

    wallet, err := h.CreateWallet(usernamePart)
    if err != nil {
        http.Error(w, "Failed to create wallet: "+err.Error(), http.StatusInternalServerError)
        return
    }

    user.WalletID = wallet.WalletID
    user.PaymailID = wallet.PaymailID
    
    
    if dbResult := h.DB.Create(&user); dbResult.Error != nil {
        http.Error(w, "Error creating user in database", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

func (h *Handler) CreateWallet(walletName string) (models.Wallet, error) {
    var wallet models.Wallet
    url := "https://dev.neucron.io/v1/wallet/create?walletName=" + walletName
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(nil)) // No request body
    if err != nil {
        return wallet, err
    }

    req.Header.Set("Accept", "application/json")
    req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTYyODAwOTQsImlhdCI6MTcxMzY4ODA5NCwiaXNzIjoiaHR0cHM6Ly9uZXVjcm9uLmlvIiwianRpIjoiNzIzNzgyM2YtMjgyZS00YjY5LWFiNWMtOTA4MTgzYjllNWQwIiwibmJmIjoxNzEzNjg4MDk0LCJzdWIiOiIwZjMxNDU3ZS04Njk1LTQxYjAtODMyMC1mMDZmODQ3Mzc5OWYiLCJ1c2VyX2lkIjoiMGYzMTQ1N2UtODY5NS00MWIwLTgzMjAtZjA2Zjg0NzM3OTlmIn0.fbncpiwM_AlfLqCc4L_1K5HahIgZDUto4EJsaIHPSM8")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return wallet, err
    }
    defer resp.Body.Close()

    responseBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return wallet, err
    }

    var result struct {
        Data struct {
            PaymailID string `json:"paymailID"`
            WalletID  string `json:"walletID"`
        } `json:"data"`
        StatusCode int `json:"status_code"`
    }

    if err := json.Unmarshal(responseBody, &result); err != nil {
        return wallet, err
    }

    if result.StatusCode != 200 {
        return wallet, fmt.Errorf("API error: %s", responseBody)
    }

    wallet.PaymailID = result.Data.PaymailID
    wallet.WalletID = result.Data.WalletID

    return wallet, nil
}


func (h *Handler) GetUserByEmailAndPassword(w http.ResponseWriter, r *http.Request) {
    // Decode JSON request body into a struct
    var request struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
        http.Error(w, "Error parsing request body", http.StatusBadRequest)
        return
    }

    // Fetch user from the database
    var user models.User
    if err := h.DB.Where("email = ? AND password = ?", request.Email, request.Password).First(&user).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            http.Error(w, "User not found", http.StatusNotFound)
        } else {
            http.Error(w, "Error fetching user from database", http.StatusInternalServerError)
        }
        return
    }

    // Respond with the fetched user
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(user)
}


func (h *Handler) ListUsers(w http.ResponseWriter, r *http.Request) {
    var listOfUsers []models.User

    if err := h.DB.Find(&listOfUsers).Error; err != nil {
        http.Error(w, "Error fetching list of users from database", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(listOfUsers)
}



func (h *Handler) GetIc(w http.ResponseWriter, r *http.Request) {
    // Decode JSON request body into a struct
    var request struct {
        IcIp string `json:"ic_ip"`
    }

    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
        http.Error(w, "Error parsing request body", http.StatusBadRequest)
        return
    }

    // Fetch IC details from the database
    var ic models.Ic
    if err := h.DB.Where("ic_ip = ?", request.IcIp).First(&ic).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            http.Error(w, "IC not found", http.StatusNotFound)
        } else {
            http.Error(w, "Error fetching IC details from database", http.StatusInternalServerError)
        }
        return
    }

    // Respond with the fetched IC details
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(ic)
}
