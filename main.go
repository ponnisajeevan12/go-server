package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Member struct {
	Name      string `json:"name"`
	StudentID string `json:"studentid"`
	Image     string `json:"image"`  // Image file name
	GitHub    string `json:"github"` // GitHub profile link
}

var members = []Member{
	{Name: "Abhishek", StudentID: "500227881", Image: "a.png", GitHub: "https://github.com/AbhishekMittal93/"},
	{Name: "Dipansu", StudentID: "500228887", Image: "d.png", GitHub: "hhttps://github.com/Dipansu/"},
	{Name: "Swaroop", StudentID: "500228657", Image: "s.png", GitHub: "https://github.com/swaroopkrishna91/"},
	{Name: "Piyush", StudentID: "500227709", Image: "p.png", GitHub: "https://github.com/PiyushIyer123/"},
	{Name: "Rose", StudentID: "500227214", Image: "r.png", GitHub: "https://github.com/rsauri/"},
	{Name: "Ponni", StudentID: "500228122", Image: "po.png", GitHub: "https://github.com/ponnisajeevan12/"},
	{Name: "Nikhitha", StudentID: "500229719", Image: "n.png", GitHub: "https://github.com/nikhitha44/"},
}

func main() {
	// Serve static files from the "static" directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Endpoint to get member details by name
	http.HandleFunc("/member", getMemberDetails)

	// Set the root URL to serve the homepage (index.html)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})

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
