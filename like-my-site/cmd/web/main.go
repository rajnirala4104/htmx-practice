package main

import (
    "fmt"
    "log"
    "net/http"
)

func handlerLikeMe(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "It works")
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", handlerLikeMe)

    log.Println("Listening on :8080")
    err := http.ListenAndServe(":8080", mux)
    log.Fatal(err)
}
