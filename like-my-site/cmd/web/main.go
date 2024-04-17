package main

import (
    //"fmt"
    "log"
    "net/http"
    "html/template"
    "sync"
)

var mu sync.Mutex
var count int = 0

func handlerLikeMe(w http.ResponseWriter, r *http.Request) {
    page, err := template.ParseFiles(
        "ui/html/home.tmpl.html",
        "ui/html/like.tmpl.html",
        "ui/html/like-button.tmpl.html",
        "ui/html/liked-button.tmpl.html",
        "ui/html/counter.tmpl.html",
        )
    if err != nil {
        log.Fatal(err.Error())
    }
    mu.Lock()
    err = page.ExecuteTemplate(w, "base", count)
    mu.Unlock()
    if err != nil {
        log.Fatal(err.Error())
    }
}

func handlerLike(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    count++
    mu.Unlock()
    t, err := template.ParseFiles(
        "ui/html/no-like.tmpl.html",
        "ui/html/liked-button.tmpl.html",
        "ui/html/counter.tmpl.html",
        )
    if err != nil {
        log.Fatal(err.Error())
    }
    mu.Lock()
    err = t.ExecuteTemplate(w, "no-like", count)
    mu.Unlock()
    if err != nil {
        log.Fatal(err.Error())
    }
}

func handlerNoLike(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    count--
    mu.Unlock()
    t, err := template.ParseFiles(
        "ui/html/like.tmpl.html",
        "ui/html/like-button.tmpl.html",
        "ui/html/counter.tmpl.html",
        )
    if err != nil {
        log.Fatal(err.Error())
    }
    mu.Lock()
    err = t.ExecuteTemplate(w, "like", count)
    mu.Unlock()
    if err != nil {
        log.Fatal(err.Error())
    }
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", handlerLikeMe)
    mux.HandleFunc("/like", handlerLike)
    mux.HandleFunc("/no-like", handlerNoLike)

    log.Println("Listening on :8080")
    err := http.ListenAndServe(":8080", mux)
    log.Fatal(err)
}
