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

	path := pathify(r.RequestURI)
	finfo, err := os.Stat(path)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Content cannot be found"))
		return
	}
	if !finfo.IsDir() {
		http.ServeFile(w, r, path)
		return
	}

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

	if path == "." {
		path = ""
	}
	list, err := listDir(path)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if path == "" {
		path = "Root"
	}
	err = templ.ExecuteTemplate(w, "files.html", View{
		Title: path,
		List:  list,
	})
	if err != nil {
		log.Println(err)
	}
}
