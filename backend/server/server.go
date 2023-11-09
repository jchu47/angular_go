package server

import (
	"fmt"
	"io"
	"net/http"
)

func HandleGetRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request!\n")
	io.WriteString(w, "Website has rendered!\n")
}

func SetupRoutes() {
	http.HandleFunc("/get", HandleGetRequest)
}

func StartServer() {
	SetupRoutes()
	fmt.Println("Server listening on 3001")
	http.ListenAndServe(":8080", nil)
}


