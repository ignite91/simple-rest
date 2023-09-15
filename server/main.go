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
	http.HandleFunc("/user", userController)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func userController(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getAllUsers(w, r)
	case "POST":
		saveUsers(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}
func getAllUsers(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("users.json")
	if err != nil {
		panic(err)
	}
	w.Write(data)
}
func saveUsers(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	user, err := manipulateData(b)
	if err != nil {
		panic(err)
	}
	userb, err := saveOnFile(user)
	if err != nil {
		panic(err)
	}
	w.Write(userb)
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

func saveOnFile(user User) ([]byte, error) {
	data, err := os.ReadFile("users.json")
	if err != nil {
		panic(err)
	}
	var users []User
	err = json.Unmarshal(data, &users)
	if err != nil {
		panic(err)
	}
	users = append(users, user)
	usersb, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}
	os.WriteFile("users.json", usersb, 0644)
	if err != nil {
		panic(err)
	}
	return usersb, nil
}
