package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

//View defines single view of the page
type View struct {
	Title string
	List  []os.FileInfo
}

var templ *template.Template

func backHandler(w http.ResponseWriter, r *http.Request) {
	if templ == nil {
		var err error
		templ, err = template.ParseGlob("*.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Template cannot be read"))
			log.Println(err)
			return
		}
	}

	list, err := listDir(r.RequestURI)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = templ.ExecuteTemplate(w, "files.html", View{
		Title: r.RequestURI[1:],
		List:  list,
	})
	if err != nil {
		log.Println(err)
	}
}
