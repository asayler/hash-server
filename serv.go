package main

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"github.com/braintree/manners"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	SLEEP       = 5  //seconds
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello\n"))
}

type HashHandler struct{}

func (h HashHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Request Info
	log.Printf("Method: %s\n", r.Method)
	log.Printf("Header: %s\n", r.Header)

	// Check for shutdown
	r.FormValue("shutdown")
	_, shutdown := r.Form["shutdown"]
	if shutdown {
		log.Printf("Received shutdown request")
		manners.Close()
		fmt.Fprintf(w, "Shutting down\n")
		return
	}

	// Extract Password
	password := r.FormValue("password")
	length := len(password)
	redacted := strings.Repeat("*", length)
	if length == 0 {
		log.Printf("Missing Password")
		http.Error(w, "Password required", http.StatusBadRequest)
		return
	}
	if length > 2 {
		redacted = string(password[0]) + strings.Repeat("*", length-2) + string(password[length-1])
	}
	log.Printf("Password: %s\n", redacted)
	password_bytes := []byte(password)

	// Sleep
	log.Printf("Sleeping for %d seconds", SLEEP)
	time.Sleep(SLEEP * time.Second)

	// Hash Password
	hash_bytes := sha512.Sum512(password_bytes)
	hash := base64.StdEncoding.EncodeToString(hash_bytes[:])
	log.Printf("Hash: %s\n", hash)

	// Reply
	log.Printf("Responding")
	fmt.Fprintf(w, "%s\n", hash)

}

func main() {

	h := HashHandler{}
	manners.ListenAndServe(":8080", h)

}
