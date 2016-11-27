// Redirect Google

package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8001", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://google.com", 301)
	//http.Redirect(w, r, "http://google.com", http.StatusMovedPermanently)
}
