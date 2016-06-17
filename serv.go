package main

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	SLEEP = 5 * time.Second
)

func hashRequest(w http.ResponseWriter, r *http.Request) {

	// Request Info
	log.Printf("Method: %s\n", r.Method)
	log.Printf("Header: %s\n", r.Header)

	// Extract Password
	password := r.FormValue("password")
	log.Printf("Password: %s\n", password)
	password_bytes := []byte(password)

	// Hash Password
	hash_bytes := sha512.Sum512(password_bytes)
	hash := base64.StdEncoding.EncodeToString(hash_bytes[:])
	log.Printf("Hash: %s\n", hash)

	// Sleep
	log.Printf("Sleeping for %d seconds", SLEEP/time.Second)
	time.Sleep(SLEEP)

	// Reply
	log.Printf("Responding")
	fmt.Fprintf(w, "%s\n", hash)

}

func main() {

	http.HandleFunc("/", hashRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
