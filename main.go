package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Resume struct {
	Name       string   `json:"name"`
	Title      string   `json:"title"`
	Location   string   `json:"location"`
	Email      string   `json:"email"`
	Summary    string   `json:"summary"`
	Skills     []string `json:"skills"`
	Languages  []string `json:"languages"`
	Education  []struct {
		Degree      string `json:"degree"`
		Institution string `json:"institution"`
		Year        string `json:"year"`
		Field       string `json:"field"`
	} `json:"education"`
	Experience []struct {
		Company     string `json:"company"`
		Position    string `json:"position"`
		Years       string `json:"years"`
		Description string `json:"description"`
	} `json:"experience"`
	Certificates []struct {
		Name         string `json:"name"`
		Issuer       string `json:"issuer"`
		Year         string `json:"year"`
		CredentialId string `json:"credential_id"`
	} `json:"certificates"`
}

func main() {
	log.Println("📦 Resume web server starting on port 8080...")
	
	// Static files handler for CSS, JS, images etc.
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	
	// Main resume handler
	http.HandleFunc("/", resumeHandler)
	
	// Health check endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"healthy","service":"resume-app"}`))
	})
	
	log.Println("🚀 Server is running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("❌ Server failed: %v", err)
	}
}

func resumeHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("➡️  Request received from %s for %s", r.RemoteAddr, r.URL.Path)

	// Read resume data
	file, err := os.Open("resume.json")
	if err != nil {
		log.Printf("❌ Error opening resume.json: %v", err)
		http.Error(w, "Resume data not found", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	var resume Resume
	if err := json.NewDecoder(file).Decode(&resume); err != nil {
		log.Printf("❌ Error decoding resume.json: %v", err)
		http.Error(w, "Error parsing resume data", http.StatusInternalServerError)
		return
	}

	// Parse and execute template
	tmpl, err := template.ParseFiles("templates/resume.html")
	if err != nil {
		log.Printf("❌ Error parsing template: %v", err)
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}

	// Set content type
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if err := tmpl.Execute(w, resume); err != nil {
		log.Printf("❌ Error executing template: %v", err)
		http.Error(w, "Error rendering page", http.StatusInternalServerError)
		return
	}

	log.Printf("✅ Resume served successfully to %s", r.RemoteAddr)
}