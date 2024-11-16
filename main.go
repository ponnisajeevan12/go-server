package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Member struct {
	Name  string `json:"name"`
	Role  string `json:"role"`
	Email string `json:"email"`
}

var members = []Member{
	{Name: "Alice", Role: "Developer", Email: "alice@example.com"},
	{Name: "Bob", Role: "Designer", Email: "bob@example.com"},
	{Name: "Charlie", Role: "Project Manager", Email: "charlie@example.com"},
}

func main() {
	// Serve static files from the "static" directory
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Endpoint to get member details by name
	http.HandleFunc("/member", getMemberDetails)

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func getMemberDetails(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	for _, member := range members {
		if member.Name == name {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(member)
			return
		}
	}
	http.Error(w, "Member not found", http.StatusNotFound)
}
