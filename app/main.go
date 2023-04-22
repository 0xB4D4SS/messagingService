package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	httptransport "github.com/go-kit/kit/transport/http"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

const tokenDefaultLength = 11
const dbConfig = "test:test@tcp(gomysql:3306)/messaging" // todo: move to env file

// is this required?
// type Endpoint func(ctx context.Context, request interface{}) (response interface{}, err error)

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

	return string(hash[:])
}

// Transports expose the service to the network. In this first example we utilize JSON over HTTP.
func main() {
	authSvc := authService{}

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

	http.Handle("/register", registerHandler)
	http.Handle("/login", loginHandler)
	http.Handle("/logout", logoutHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
