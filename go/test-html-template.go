package main

import (
	"html/template"
	"log"
	"net/http"
)

// Data structure to hold dynamic content
type PageData struct {
	Title   string
	Name    string
	Message string
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Parse the template
	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}

	// Data to inject into the template
	data := PageData{
		Title:   "My Go Web App",
		Name:    "John Doe",
		Message: "Have a great day!",
	}

	// Execute the template and write to the response
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Starting server on :8080")

	// Start the server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

}
