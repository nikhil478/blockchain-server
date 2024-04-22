package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/FalconX80/blockchain-server/pkg/models"
)

func (h *Handler) CreateIc(w http.ResponseWriter, r *http.Request) {
	
	var ic models.Ic

	if err := json.NewDecoder(r.Body).Decode(&ic); err != nil {
		http.Error(w, "Error parsing request data", http.StatusBadRequest)
		return
	}

	wallet, err := h.CreateWallet(ic.IcIp)
    if err != nil {
        http.Error(w, "Failed to create wallet: "+err.Error(), http.StatusInternalServerError)
        return
    }

    ic.WalletID = wallet.WalletID
    ic.PaymailID = wallet.PaymailID

	if result := h.DB.Create(&ic); result.Error != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(ic)
}
