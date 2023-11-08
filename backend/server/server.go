package server

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func handleGetRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request!\n")
	io.WriteString(w, "Website has rendered!\n")
}

func setupRoutes() {
	http.HandleFunc("/get", handleGetRequest)
}

func StartServer() {
	setupRoutes()
	fmt.Println("Server listening on 3001")
	http.ListenAndServe(":8080", nil)
}


