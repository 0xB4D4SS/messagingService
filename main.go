package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	dotenv "github.com/joho/godotenv"
	"log"
	"messagingService/app/services"
	"messagingService/app/transports"
	"net/http"
	"os"
)

func initEnv() {
	err := dotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading env file. Err: %s", err)
	}
}

func main() {
	initEnv()
	db, conErr := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_CONFIG"))
	authSvc := services.AuthService{}
	messageSvc := services.MessageService{}

	if conErr != nil {
		log.Fatal(conErr)
	}

	log.Fatal(http.ListenAndServe(":8080", transports.MakeHTTPHandler(authSvc, messageSvc, db)))
}
