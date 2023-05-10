package main

import (
	_ "github.com/go-sql-driver/mysql"
	dotenv "github.com/joho/godotenv"
	"log"
	"net/http"
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
	//db, conErr := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_CONFIG"))
	authSvc := authService{}
	messageSvc := messageService{}

	log.Fatal(http.ListenAndServe(":8080", MakeHTTPHandler(authSvc, messageSvc)))
}
