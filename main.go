package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Resume struct {
	Name      string   `json:"name"`
	Title     string   `json:"title"`
	Summary   string   `json:"summary"`
	Skills    []string `json:"skills"`
	Experience []struct {
		Company  string `json:"company"`
		Position string `json:"position"`
		Years    string `json:"years"`
	} `json:"experience"`
}

func main() {
	log.Println("📦 Resume web server starting on port 8080...")
	http.HandleFunc("/", resumeHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("❌ Server failed: %v", err)
	}
}

func resumeHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("➡️  Request received from %s", r.RemoteAddr)

	file, err := os.Open("resume.json")
	if err != nil {
		log.Printf("❌ Error opening resume.json: %v", err)
		http.Error(w, "Dosya bulunamadı", 500)
		return
	}
	defer file.Close()

	var res Resume
	if err := json.NewDecoder(file).Decode(&res); err != nil {
		log.Printf("❌ Error decoding resume.json: %v", err)
		http.Error(w, "JSON çözümleme hatası", 500)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/resume.html"))
	if err := tmpl.Execute(w, res); err != nil {
		log.Printf("❌ Error executing template: %v", err)
	}
}
