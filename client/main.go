package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type User struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	LastName string    `json:"lastname"`
	Age      int       `json:"age"`
	Active   bool      `json:"active"`
	Money    float64   `json:"money"`
	SaveAT   time.Time `json:"saveat"`
}

func main() {
	user := User{
		ID:       1,
		Name:     "name",
		LastName: "lastName",
		Age:      99,
		Active:   true,
		Money:    33999.22,
	}
	b, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	bReader := bytes.NewReader(b)
	resp, err := http.Post("http://0.0.0.0:8080/user", "POST", bReader)
	if err != nil {
		panic(err)
	}
	log.Println("status CODE: ", resp.StatusCode)
}
