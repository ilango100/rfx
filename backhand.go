package main

import (
	"html/template"
	"io/ioutil"
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
			return
		}
	}

	list, err := ioutil.ReadDir(r.RequestURI[1:])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	templ.ExecuteTemplate(w, "files.html", View{
		Title: r.RequestURI[1:],
		List:  list,
	})
}
