package main

import (
	"log"
	"net/http"
)

func test(res http.ResponseWriter, req *http.Request) {
	log.Println("test called")
	log.Println(req.Header)
	log.Println(req.Body)
}

func main() {
	log.Println("main called")

	http.HandleFunc("/test/", test)
	http.Handle("/google", http.RedirectHandler("http://www.google.com", 301))
	log.Println("route(s) handled")

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Println(err)
	}
}
