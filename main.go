package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

type Dork struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Category    string   `json:"category"`
	Description string   `json:"description"`
	Template    string   `json:"template"`
	TargetTypes []string `json:"targetTypes"`
}

type GeneratedDork struct {
	Name        string
	Category    string
	Description string
	Query       string
}

type PageData struct {
	Target      string
	TargetType  string
	Category    string
	Results     []GeneratedDork
	Categories  []string
	ResultCount int
}

var dorks []Dork

func loadDorks(filename string) ([]Dork, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var loaded []Dork
	err = json.Unmarshal(file, &loaded)
	if err != nil {
		return nil, err
	}

	return loaded, nil
}

func contains(list []string, item string) bool {
	for _, v := range list {
		if strings.EqualFold(v, item) {
			return true
		}
	}
	return false
}

func getCategories(dorks []Dork) []string {
	seen := make(map[string]bool)
	var categories []string

	for _, d := range dorks {
		if !seen[d.Category] {
			seen[d.Category] = true
			categories = append(categories, d.Category)
		}
	}

	return categories
}

func generateDorks(target, targetType, category string, dorks []Dork) []GeneratedDork {
	var results []GeneratedDork

	for _, d := range dorks {
		if !contains(d.TargetTypes, targetType) {
			continue
		}

		if category != "" && category != "All" && !strings.EqualFold(d.Category, category) {
			continue
		}

		query := strings.ReplaceAll(d.Template, "{target}", target)

		results = append(results, GeneratedDork{
			Name:        d.Name,
			Category:    d.Category,
			Description: d.Description,
			Query:       query,
		})
	}

	return results
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	data := PageData{
		Categories: getCategories(dorks),
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Template render hatası", http.StatusInternalServerError)
	}
}

func generateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	target := strings.TrimSpace(r.FormValue("target"))
	targetType := strings.TrimSpace(r.FormValue("targetType"))
	category := strings.TrimSpace(r.FormValue("category"))

	results := generateDorks(target, targetType, category, dorks)

	tmpl := template.Must(template.ParseFiles("templates/results.html"))
	data := PageData{
		Target:      target,
		TargetType:  targetType,
		Category:    category,
		Results:     results,
		ResultCount: len(results),
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Results template render hatası", http.StatusInternalServerError)
	}
}

func main() {
	var err error
	dorks, err = loadDorks("data/dorks.json")
	if err != nil {
		log.Fatal("dorks.json yüklenemedi:", err)
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/generate", generateHandler)

	log.Println("Server çalışıyor: http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server başlatılamadı:", err)
	}
}