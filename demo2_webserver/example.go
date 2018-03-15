package main

import (
    "net/http"
    "log"

    "github.com/gorilla/mux"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello World!\n"))
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", YourHandler)

    log.Fatal(http.ListenAndServe(":8080", r))
}
