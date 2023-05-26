package main

import (
	"log"
	"net/http"
)

var shortToLong = make(map[string]string)

func shorten(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	long := req.FormValue("long")
	if long == "" {
		http.Error(res, "Bad Request", http.StatusBadRequest)
		return
	}

	short := GenerateShort()
	for _, ok := shortToLong[short]; ok; short = GenerateShort() {
	}

	shortToLong[short] = long
	WriteLinked(res, short)
	log.Println(short, long)
}

func mapper(res http.ResponseWriter, req *http.Request) {
	short := req.URL.Path[len("/m/"):]
	http.Redirect(res, req, shortToLong[short], http.StatusMovedPermanently)
}

func main() {
	log.Println("main called")

	var fs http.Handler = http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/submit", shorten)
	http.HandleFunc("/m/", mapper)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Println(err)
	}
}

//next step is to store active routes in file on exit, and read active routes on start
