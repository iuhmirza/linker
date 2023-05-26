package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var linkedPageTemplate = filepath.Join("templates", "linked.html")
var linkedPage = template.Must(template.New("linked").ParseFiles(linkedPageTemplate))

func WriteLinked(res http.ResponseWriter, short string) {
	res.Header().Set("Content-Type", "text/html")
	err := linkedPage.Execute(res, short)
	if err != nil {
		log.Println(err)
	}
}
