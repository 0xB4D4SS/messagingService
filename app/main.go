package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
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

func GenerateSecureToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

func GenerateSHA256Hash(input string) string {
	h := sha256.New()
	h.Write([]byte(input))
	hash := h.Sum(nil)
	output := fmt.Sprintf("%x", string(hash[:]))

	return output
}

// Transports expose the service to the network. In this first example we utilize JSON over HTTP.
func main() {
	initEnv()
	authSvc := authService{}
	messageSvc := messageService{}

	log.Fatal(http.ListenAndServe(":8080", MakeHTTPHandler(authSvc, messageSvc)))
}
