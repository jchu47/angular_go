package controller

import (
	"angular_go/backend/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

// asterick is a pointer bc sql.DB is a large data structure and shoould be passed by referrence to it's memory address
func ConnectToDB() (*sql.DB, error) {
	connStr := os.Getenv("DATABASE_URL")
	//if the below line is successful, db = a pointer to sql.DB and err = nil
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// w & r from net/http standard lib. w: sends response to client, r: reads from client
func GetAllAlgos(w http.ResponseWriter, r *http.Request) {
	fmt.Println("recieved /get request!")

	// Get the database connection.
	db, err := ConnectToDB()
	if err != nil {
		// if error connecting, return an http error message
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// defer to close after all below functions execute.  Good practice to make sure it closes in case there was an error when connecting to db.
	defer db.Close()

	// SELECT all, duhhhhh
	rows, err := db.Query("SELECT * FROM algos")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// create empty, dynamic slice to store Algo objectes
	var algos []models.Algo
	//.Next will itereate over each row and return false when nothing left
	for rows.Next() {
		// create a Algo object for each iteration
		var algo models.Algo
		// scan to read from each iteration into the empty Algo object
		err := rows.Scan(&algo.ID, &algo.Name, &algo.Resource)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// push to Algo slice above the for loop
		algos = append(algos, algo)
	}

	// Return the Algo slice to the client as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(algos)
}

func PostAlgo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("recieved /post request!")

	db, err := ConnectToDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Extract the name and resource from the request
	var requestBody models.RequestBody
    err = json.NewDecoder(r.Body).Decode(&requestBody)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
	name := requestBody.Name
	resource := requestBody.Resource
	// Insert record
	result, err := db.Exec("INSERT INTO algos (Name, Resource) VALUES ($1, $2)", name, resource)
	if err != nil {
		fmt.Println("error inserting", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}