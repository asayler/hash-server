package main

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"github.com/braintree/manners"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	DEFAULT_PORT = 8080
	SLEEP        = 5  //seconds
	SALT_LENGTH  = 32 //bytes
	SALT_SEP     = "|"
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

	// Extract Salt
	salt := r.FormValue("salt")
	_, salt_ok := r.Form["salt"]
	salt_bytes := []byte(salt)
	if salt_ok {
		if len(salt_bytes) == 0 {
			salt_bytes = make([]byte, SALT_LENGTH)
			_, salt_err := rand.Read(salt_bytes)
			if salt_err != nil {
				log.Printf("Error generating salt")
				http.Error(w, "Error generating salt", http.StatusInternalServerError)
				return
			}
			salt = base64.StdEncoding.EncodeToString(salt_bytes[:])
		}
		log.Printf("Salt: %s\n", salt)
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

	// Extract Rounds
	rounds_str := r.FormValue("rounds")
	var rounds int
	if len(rounds_str) > 0 {
		val, conv_err := strconv.Atoi(rounds_str)
		if conv_err != nil {
			log.Printf("Error parsing rounds")
			http.Error(w, "Error parsing rounds", http.StatusBadRequest)
			return
		} else {
			rounds = val
		}
	} else {
		rounds = 1
	}

	// Sleep
	log.Printf("Sleeping for %d seconds", SLEEP)
	time.Sleep(SLEEP * time.Second)

	// Hash Password
	var tohash []byte
	if len(salt_bytes) > 0 {
		tohash = append(salt_bytes, password_bytes...)
	} else {
		tohash = password_bytes
	}
	log.Printf("Hashing %d rounds\n", rounds)
	var hash_bytes [64]byte
	for i := 0; i < rounds; i++ {
		hash_bytes = sha512.Sum512(tohash)
		tohash = hash_bytes[:]
	}
	hash := base64.StdEncoding.EncodeToString(hash_bytes[:])
	log.Printf("Hash: %s\n", hash)

	// Reply
	res := ""
	if len(salt_bytes) > 0 {
		res += salt + SALT_SEP
	}
	res += hash
	log.Printf("Responding: %s", res)
	fmt.Fprintf(w, "%s\n", res)

}

func main() {

	// Parse CLI
	var port int
	if len(os.Args) > 2 {
		fmt.Printf("Too many args\n")
		fmt.Printf("Usage :\n%s [port]", os.Args[0])
	}
	if len(os.Args) == 2 {
		val, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("Error parsing port")
			os.Exit(1)
		} else {
			port = val
		}
	} else {
		port = DEFAULT_PORT
	}

	h := HashHandler{}
	log.Printf("Starting server on port %d\n", port)
	path := ":" + strconv.Itoa(port)
	manners.ListenAndServe(path, h)
	os.Exit(0)

}
