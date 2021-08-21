package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "server.go")
}

func createHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Create request from %s!\n", r.URL.Path)
    fmt.Fprintf(w, "%+v", r)
}

func readHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Read request from %s!\n", r.URL.Path)
    fmt.Fprintf(w, "%+v", r)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Update request from %s!\n", r.URL.Path)
    fmt.Fprintf(w, "%+v", r)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Delete request from %s!\n", r.URL.Path)
    fmt.Fprintf(w, "%+v", r)
}

func main() {
    fmt.Println("Starting issue tracker version 0.1")
    http.HandleFunc("/", handler)
    http.HandleFunc("/create", createHandler)
    http.HandleFunc("/read", readHandler)
    http.HandleFunc("/update", updateHandler)
    http.HandleFunc("/delete", deleteHandler)
    log.Fatal(http.ListenAndServe(":3000", nil))
}
