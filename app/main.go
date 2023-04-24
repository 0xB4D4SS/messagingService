package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	httptransport "github.com/go-kit/kit/transport/http"
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

	registerHandler := httptransport.NewServer(
		makeRegisterEndpoint(authSvc),
		decodeRegisterRequest,
		encodeResponse,
	)

	loginHandler := httptransport.NewServer(
		makeLoginEndpoint(authSvc),
		decodeLoginRequest,
		encodeResponse,
	)

	logoutHandler := httptransport.NewServer(
		makeLogoutEndpoint(authSvc),
		decodeLogoutRequest,
		encodeResponse,
	)

	sendHandler := httptransport.NewServer(
		makeSendEndpoint(messageSvc),
		decodeSendRequest,
		encodeResponse,
	)

	getHandler := httptransport.NewServer(
		makeGetEndpoint(messageSvc),
		decodeGetRequest,
		encodeResponse,
	)

	http.Handle("/register", registerHandler)
	http.Handle("/login", loginHandler)
	http.Handle("/logout", logoutHandler)
	http.Handle("/send", sendHandler)
	http.Handle("/get", getHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
