package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	dotenv "github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

const tokenDefaultLength = 11

// is this required?
// type Endpoint func(ctx context.Context, request interface{}) (response interface{}, err error)

func initEnv() {
	err := dotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading env file. Err: %s", err)
	}
}

func main() {
	initEnv()
	db, conErr := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_CONFIG"))
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)
	authSvc := authService{}
	messageSvc := messageService{}

	if conErr != nil {
		log.Fatal(conErr)
	}

	log.Fatal(http.ListenAndServe(":8080", MakeHTTPHandler(authSvc, messageSvc, db)))
}
