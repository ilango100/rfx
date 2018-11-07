package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
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

	if r.Method == http.MethodPost {
		if err := r.ParseMultipartForm(1024 * 1024 * 512); err != nil {
			log.Println(err)
		}
		fls, ok := r.MultipartForm.File["file"]
		if !ok {
			log.Println("file not found")
		}
		for i := range fls {
			dfl, err := os.OpenFile(filepath.Join(path, fls[i].Filename), os.O_APPEND|os.O_CREATE, 0644)
			if err != nil {
				log.Println(err)
			}
			fl, err := fls[i].Open()
			if err != nil {
				fmt.Println(err)
			}
			io.Copy(dfl, fl)
			fl.Close()
			dfl.Close()
		}
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
