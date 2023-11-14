package controller
import (
	"fmt"
	"os"
	"database/sql"
	"net/http"
	"encoding/json"
	"angular_go/backend/models"
)

func ConnectToDB() {
	connStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return db
}

func GetAllAlgos(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got / request!")

	// Get the database connection.
	db, err := ConnectToDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Query the database for all Algo records.
	rows, err := db.Query("SELECT * FROM algos")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Iterate over the results and return them to the client.
	var algos []models.Algo
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

