package controller
import (
	"fmt"
	"angular_go/backend/models"
	"strconv"
)


func getAllAlgos() {
	fmt.Printf("got / request!\n")

  // Get the database connection.
  models.ConnectToDB()

  // Get the algo ID from the request query parameters.
  algoID, err := strconv.Atoi(r.URL.Query().Get("id"))
  if err != nil {
    fmt.Println(err)
    return
  }

  // Query the database for the algo record with the given ID.
  var algo Algo
  err = db.QueryRow("SELECT * FROM algos WHERE id = $1", algoID).Scan(&algo.ID, &algo.Name, &algo.Resource)
  if err != nil {
    fmt.Println(err)
    return
  }

  // Return the algo to the client as JSON.
  json.NewEncoder(w).Encode(algo)
}