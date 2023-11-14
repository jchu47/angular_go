package server

import (
	"fmt"
	"io"
	"net/http"
	"angular_go/backend/controller"
)

// CURRENTLY WORKING ON CONFIGURING GET ROUTE HANDLERS TO GET ALL RECORDS

func HandleGetRequest(w http.ResponseWriter, r *http.Request) {
	controller.getAllAlgos()
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


