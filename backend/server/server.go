package server

import (
	"angular_go/backend/controller"
	"fmt"
	"net/http"
)

// CURRENTLY WORKING ON CONFIGURING GET ROUTE HANDLERS TO GET ALL RECORDS
func SetupRoutes() {
	http.HandleFunc("/get", controller.GetAllAlgos)
	http.HandleFunc("/post", controller.PostAlgo)
	// http.HandleFunc("/update", controller.GetAllAlgos)
	// http.HandleFunc("/delete", controller.GetAllAlgos)
}

func StartServer() {
	SetupRoutes()
	fmt.Println("Server listening on 3001")
	http.ListenAndServe(":8080", nil)
}


