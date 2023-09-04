package main

import (
	"encoding/json"
	"net/http"
	"os"
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
	users, err := GetAllUsers()
	if err != nil {
		panic(err)
	}
	b, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}
	w.Write(b)
}

type Users struct {
	Users []User `json:"users"`
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

func GetAllUsers() ([]Users, error) {
	b, err := os.ReadFile("accounts.json")
	if err != nil {
		panic(err)
	}
	var users []Users
	err = json.Unmarshal(b, &users)
	if err != nil {
		panic(err)
	}
	return users, nil
}
