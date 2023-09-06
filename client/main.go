package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	resp, err := http.Get("http://0.0.0.0:8080/user")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var user User
	json.Unmarshal(b, &user)
	fmt.Println(user)
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
