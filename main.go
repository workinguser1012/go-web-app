package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

type PageData struct {
	Title       string
	CurrentPage string
}

func render(w http.ResponseWriter, tmpl string, data PageData) {
	err := templates.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	render(w, "home.html", PageData{Title: "Home", CurrentPage: "home"})
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "about.html", PageData{Title: "About", CurrentPage: "about"})
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "contact.html", PageData{Title: "Contact", CurrentPage: "contact"})
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"ok"}`))
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/health", healthHandler)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}