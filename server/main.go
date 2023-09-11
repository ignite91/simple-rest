package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
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
	http.HandleFunc("/user", saveUserController)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func saveUserController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		user, err := manipulateData(b)
		if err != nil {
			panic(err)
		}
		err = saveOnFile(user)
		if err != nil {
			panic(err)
		}
		w.Write(b)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func manipulateData(b []byte) (User, error) {
	var user User
	err := json.Unmarshal(b, &user)
	if err != nil {
		panic(err)
	}
	user.SaveAT = time.Now()

	return user, nil
}

func saveOnFile(user User) error {
	userb, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("users.json", userb, 0644)
	if err != nil {
		panic(err)
	}
	return nil
}
