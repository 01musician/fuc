package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// User represents a simple user struct for demonstration
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// In-memory data
var users = []User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob"},
}

// Get all users
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Get a single user
func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	for _, user := range users {
		if fmt.Sprintf("%d", user.ID) == id {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

// Add a new user
func addUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}

// Main function
func main() {
	http.HandleFunc("/users", getUsers)        // GET /users
	http.HandleFunc("/user", getUser)         // GET /user?id=<id>
	http.HandleFunc("/user/add", addUser)     // POST /user/add

	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

