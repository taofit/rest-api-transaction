// +build ignore

package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "pong")
}

func main() {
    http.HandleFunc("/ping", handler)
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("I am root"))
    })
    log.Fatal(http.ListenAndServe(":8080", nil))
}