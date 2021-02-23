package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"

	"github.com/ofstudio/advice-wcs-life/advices"
)

//go:embed templates/index.tmpl
var indexTmpl string

//go:embed assets/*
var assetsDir embed.FS

func main() {

	var t, err = template.New("foo").Parse(indexTmpl)
	if err != nil {
		log.Panic("Error parsing template")
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		a := advices.GetAdvice()
		log.Printf("%s %s %s %s", r.RemoteAddr, r.Method, r.RequestURI, r.UserAgent())
		w.WriteHeader(200)
		w.Header().Add("Content-type", "text/html")
		t.Execute(w, a)
	}

	http.HandleFunc("/", handler)
	http.Handle("/assets/", http.FileServer(http.FS(assetsDir)))
	log.Print("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}
