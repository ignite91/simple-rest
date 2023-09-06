package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/user", getAllUsersController)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func getAllUsersController(w http.ResponseWriter, r *http.Request) {
	users, err := GetUser()
	if err != nil {
		panic(err)
	}
	b, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}
	w.Write(b)
}

type User struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	LastName string    `json:"lastname"`
	Age      int       `json:"age"`
	Active   bool      `json:"active"`
	Account  []Account `json:"account"`
	SaveAT   time.Time `json:"saveat"`
}

type Account struct {
	ID      int     `json:"id"`
	Number  int     `json:"number"`
	Type    string  `json:"type"`
	Balance float64 `json:"balance"`
}

func GetUser() (User, error) {
	acc := Account{
		ID:      1,
		Number:  111,
		Type:    "type-account",
		Balance: 9999999999,
	}
	user := User{
		ID:       1,
		Name:     "name",
		LastName: "lastname",
		Age:      5,
		Active:   true,
		Account:  []Account{acc, acc, acc, acc, acc, acc, acc, acc, acc, acc, acc, acc, acc, acc, acc, acc, acc, acc},
		SaveAT:   time.Now(),
	}

	return user, nil
}
