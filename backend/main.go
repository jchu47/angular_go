package main
// currently working on being able to import packets from my own repo... this is stupid, it's mine...  models.go now contains the schema

import (
	"fmt"
	"io"
	"net/http"
	"github.com/jchu47/angular_go/models"
)

func main() {
	fmt.Println("Whattup")
	StartServer()
}

// BELOW CODE WAS PREVIOUSLY IN SERVER.GO, moved due to sucking at Go

// test with localhost:8080/get on POSTMAN
func HandleGetRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request!\n")
	// add code to retrieve records from PostgreSQL
	io.WriteString(w, "Website has rendered!\n")
}

// runs the route handlers at /get endpoint
func SetupRoutes() {
	http.HandleFunc("/get", HandleGetRequest)
}

// starts the http server, runs the route handlers, starts server on port 8080.
// nil (null)
func StartServer() {
	SetupRoutes()
	fmt.Println("Server listening on 3001")
	http.ListenAndServe(":8080", nil)
}