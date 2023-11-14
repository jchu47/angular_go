package controller

import (
	"fmt"
	"database/sql"
	"os"
	_ "github.com/lib/pq"
	"net/http"
	"encoding/json"
	"angular_go/backend/models"
)

func ConnectToDB() (*sql.DB, error) {
	connStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func GetAllAlgos(w http.ResponseWriter, r *http.Request) {
	fmt.Println("recieved /get request!")

	// Get the database connection.
	db, err := ConnectToDB()
	if err != nil {
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

	// Iterate over the results and return them to the client.
	var algos []models.Algo
	//.Next will itereate over each row and return false when nothing left
	for rows.Next() {
		var algo models.Algo
		err := rows.Scan(&algo.ID, &algo.Name, &algo.Resource)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		algos = append(algos, algo)
	}

	// Return the algos to the client as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(algos)
}

