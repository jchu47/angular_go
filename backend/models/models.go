package models

import (
	"database/sql"
	"fmt"
	"os"
)

func ConnectToDB() {
	// Create a database connection string.
	connStr := os.Getenv("DATABASE_URL")
	// "postgres://kfexssfq:3E-cEMRK3wFxs-kKf872lJWEll7Ts3xh@suleiman.db.elephantsql.com/kfexssfq"

	// Open a connection to the database.
	db, err := sql.Open("postgres", connStr)
	if err != nil {
			panic(err)
	}

	// Close the connection when we're done.
	defer db.Close()

	// Query the database.
	rows, err := db.Query("SELECT * FROM algos")
	if err != nil {
			panic(err)
	}

	// Iterate over the results.
	for rows.Next() {
			var algo Algo
			err := rows.Scan(&algo.ID, &algo.Name, &algo.Resource)
			if err != nil {
					panic(err)
			}
			// Print the algo information.
			fmt.Println(algo.ID, algo.Name, algo.Resource)
	}
}

type Algo struct {
	ID int64
	// UserID int64
	Name string
	Resource string
}


// MAYBE LATER

// type Account struct {
// 	ID int64
// 	Name string
// 	Password string
// }