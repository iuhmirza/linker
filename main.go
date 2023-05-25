package main

import (
	"log"
	"net/http"
)

var shortToLong = make(map[string]string)

func shorten(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		res.WriteHeader(http.StatusMethodNotAllowed)
	}
	long := req.FormValue("longlink")
	var short string = GenerateShort()
	for _, ok := shortToLong[short]; ok; short = GenerateShort() {
	}
	log.Println(short)
	shortToLong[short] = long
}

func main() {
	log.Println("main called")

	var fs http.Handler = http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/submit", shorten)

	http.Handle("/google", http.RedirectHandler("http://www.google.com", http.StatusMovedPermanently))
	log.Println("route(s) handled")

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Println(err)
	}
}
