package server

import (
	"log"
	"net/http"
)


func indexHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("<h1>Welcome to my web server!</h1>"))
}

func ServeHttp() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", indexHandler)
    log.Fatal(http.ListenAndServe(":8080", mux))
}