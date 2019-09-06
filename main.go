package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

const BumsKimCom = "bumskim.com"

func main() {

	r := mux.NewRouter()
	r.PathPrefix("/.well-known/acme-challenge/").Handler(
		http.StripPrefix("/.well-known/acme-challenge/", http.FileServer(http.Dir("./letsencrypt/"))))
	http.Handle("/", r)
	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:80",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println(srv.ListenAndServe())
	// TODO : need ssl certification for https
	//        it has bug (https -> http redirect fail cause of https certification
	//go func() {
	//	if err := http.ListenAndServe(":443", http.HandlerFunc(redirectTLS)); err != nil {
	//		log.Fatalf("ListenAndServe error: %v", err)
	//	}
	//}()

	//log.Println("server start!")
	////http.HandleFunc("/", TestController)
	//err := http.ListenAndServe(":80", nil)
	//if err != nil {
	//	log.Println(err)
	//}
}

func redirectTLS(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://"+BumsKimCom+r.RequestURI, http.StatusTemporaryRedirect)
}

func TestController(w http.ResponseWriter, r *http.Request) {
	html := `
<html>
<style>* {margin:0; padding:0;}</style>
<body>
<div>
<video autoplay muted loop>
    <source src="https://d2qguwbxlx1sbt.cloudfront.net/TextInMotion-VideoSample-1080p.mp4" type="video/mp4">
</video>
</div>
</body>
</html>`
	_ = r
	_, err := fmt.Fprintln(w, html)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
