package main

import(
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Resume struct{
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

func main(){
	http.HandleFunc("/",resumeHandler)
	log.Println("Sunucu çalışıyor: http://localhost:8080")
	http.ListenAndServe(":8080",nil)
}

func resumeHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("resume.json")
	if err != nil {
		http.Error(w, "Dosya bulunamadı", 500)
		return
	}
	defer file.Close()

	var res Resume
	if err := json.NewDecoder(file).Decode(&res); err != nil {
		http.Error(w, "JSON çözümleme hatası", 500)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/resume.html"))
	tmpl.Execute(w, res)
}
