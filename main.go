package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"

	"github.com/ofstudio/advice-wcs-life/advices"
)

//go:embed "templates"
var templatesFS embed.FS

//go:embed "assets"
var assetsFS embed.FS

var tmpl = template.Must(template.ParseFS(templatesFS, "templates/*.gohtml"))

func handle(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s %s", r.RemoteAddr, r.Method, r.RequestURI, r.UserAgent())
	err := tmpl.ExecuteTemplate(w, "index.gohtml", advices.GetAdvice())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Print(err)
	}
}

func main() {
	http.HandleFunc("/", handle)
	http.Handle("/assets/", http.FileServer(http.FS(assetsFS)))
	log.Print("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
