package models

type Algo struct {
	ID int64
	// UserID int64
	Name string
	Resource string
}

type RequestBody struct {
	Name     string `json:"name"`
	Resource string `json:"resource"`
}

// MAYBE LATER

// type Account struct {
// 	ID int64
// 	Name string
// 	Password string
// }