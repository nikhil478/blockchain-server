package main

import (
	"log"
	"net/http"

	"github.com/FalconX80/blockchain-server/pkg/config"
	"github.com/FalconX80/blockchain-server/pkg/database"
	"github.com/FalconX80/blockchain-server/pkg/handlers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }
	cfg := config.New()
	db := database.SetupDatabase(cfg)
	handler := handlers.NewHandler(db)

	http.HandleFunc("/create_user", handler.CreateUser)
	http.HandleFunc("/create_ic", handler.CreateIc)
	http.HandleFunc("/get_ic", handler.GetIc)
	http.HandleFunc("/get_user", handler.GetUserByEmailAndPassword)
	http.HandleFunc("/user_list", handler.ListUsers)

	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
