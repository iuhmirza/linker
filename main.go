package main

import (
	"log"
	"net/http"
)

var shortToLong = make(map[string]string)

func shorten(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, "Method Not Allowed", http.StatusMethodNotAllowed)
	}

	long := req.FormValue("long")
	if long == "" {
		http.Error(res, "Bad Request", http.StatusBadRequest)
	}

	short := GenerateShort()
	for _, ok := shortToLong[short]; ok; short = GenerateShort() {
	}

	shortToLong[short] = long
	WriteLinked(res, short)
	http.Handle("/"+short, http.RedirectHandler(long, http.StatusPermanentRedirect))
	log.Println(short, long)
}

func main() {
	log.Println("main called")

	var fs http.Handler = http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/submit", shorten)

	log.Println("route(s) handled")

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Println(err)
	}
}

//next step is to store active routes in file on exit, and read active routes on start
