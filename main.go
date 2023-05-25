package main

import (
	"log"
	"net/http"
)

func test(res http.ResponseWriter, req *http.Request) {
	log.Println("test called")
	log.Println(req.Header)
	log.Println(req.Body)
	log.Println(req.Method)
	log.Println(req.FormValue("longlink"))
}

func main() {
	log.Println("main called")

	var fs http.Handler = http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/submit", test)

	http.HandleFunc("/test/", test)
	http.Handle("/google", http.RedirectHandler("http://www.google.com", http.StatusMovedPermanently))
	log.Println("route(s) handled")

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Println(err)
	}
}
