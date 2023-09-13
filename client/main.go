package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	t := time.Now()
	for i := 0; i < 10000; i++ {
		//getUsers()
		saveUser()
	}
	fmt.Println("Elapsed: ", time.Since(t).Seconds())
}

func getUsers() {
	_, err := http.Get("http://0.0.0.0:8080/user")
	if err != nil {
		panic(err)
	}
}

func saveUser() {
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
	_, err = http.Post("http://0.0.0.0:8080/user", "POST", bReader)
	if err != nil {
		panic(err)
	}
}
