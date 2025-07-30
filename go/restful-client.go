package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// User represents a simple user struct
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	baseURL := "http://localhost:8080"

	// GET all users
	resp, err := http.Get(baseURL + "/users")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("GET /users response:", string(body))

	// GET a single user
	resp, err = http.Get(baseURL + "/user?id=1")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ = ioutil.ReadAll(resp.Body)
	fmt.Println("GET /user response:", string(body))

	// POST a new user
	newUser := User{ID: 3, Name: "Charlie"}
	data, _ := json.Marshal(newUser)
	resp, err = http.Post(baseURL+"/user/add", "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ = ioutil.ReadAll(resp.Body)
	fmt.Println("POST /user/add response:", string(body))
}

