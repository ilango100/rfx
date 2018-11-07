package main

import (
	"html/template"
	"log"
	"net/http"
)

var templ *template.Template

func backHandler(w http.ResponseWriter, r *http.Request) {
	if templ == nil {
		var err error
		templ, err = template.ParseGlob("*.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Server error occured"))
		}
	}
	w.Write([]byte(r.RequestURI))
	log.Println(r.RequestURI)
	log.Println(r.URL)
}
