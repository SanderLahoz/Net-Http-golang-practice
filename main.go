package main

import (
	"fmt"
	"net/http"
)

type welcome string

func (wc welcome) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to our server!")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login to our server!")
}

func main() {
	// Router
	router := http.NewServeMux()

	// Handler
	var wc welcome
	router.Handle("/", wc)

	// Handler Funcs
	router.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Logout Page!")
	})

	router.HandleFunc("/login", login)

	// server
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()
}
