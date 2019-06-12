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
	html := "<h5><div><video controls loop><source src'https://d2qguwbxlx1sbt.cloudfront.net/TextInMotion-VideoSample-1080p.mp4' type='video/mp4'></video></div></h5>"
	_ = r
	_, err := fmt.Fprintln(w, html)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
