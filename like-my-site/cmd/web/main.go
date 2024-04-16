package main

import (
    //"fmt"
    "log"
    "net/http"
    "html/template"
)

func handlerLikeMe(w http.ResponseWriter, r *http.Request) {
    page, err := template.ParseFiles(
        "ui/html/home.tmpl.html",
        "ui/html/like.tmpl.html",
        "ui/html/like-button.tmpl.html",
        "ui/html/liked-button.tmpl.html",
        )
    if err != nil {
        log.Fatal(err.Error())
    }
    err = page.ExecuteTemplate(w, "base", nil)
    if err != nil {
        log.Fatal(err.Error())
    }
}

func handlerLike(w http.ResponseWriter, r *http.Request) {
    t, err := template.ParseFiles(
        "ui/html/no-like.tmpl.html",
        "ui/html/liked-button.tmpl.html",
        )
    if err != nil {
        log.Fatal(err.Error())
    }
    err = t.ExecuteTemplate(w, "no-like", nil)
    if err != nil {
        log.Fatal(err.Error())
    }
}

func handlerNoLike(w http.ResponseWriter, r *http.Request) {
    t, err := template.ParseFiles(
        "ui/html/like.tmpl.html",
        "ui/html/like-button.tmpl.html",
        )
    if err != nil {
        log.Fatal(err.Error())
    }
    err = t.ExecuteTemplate(w, "like", nil)
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
