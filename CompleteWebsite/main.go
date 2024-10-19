package main

import (
	"html/template"
	"net/http"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*html"))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", nil)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "contact.html", nil)
}

func abouHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "about.html", nil)
}

func main() {
	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))
	http.HandleFunc("/about", abouHandler)
	http.HandleFunc("/contact", contactHandler)

	http.HandleFunc("/", homeHandler)
	http.ListenAndServe(":9999", nil)

}
