package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", TestController)
	log.Println("server start!")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Println(err)
	}
}

func TestController(w http.ResponseWriter, r *http.Request) {
	_ = r
	_, err := fmt.Fprintln(w, "<h1>Hello World!</h1>")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
