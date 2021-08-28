package main

import (
    "fmt"
    "log"
    "net/http"
)

var data map[string]string

func handler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "server.go")
}

func createHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Create request from %s!\n", r.URL.Path)
    if r.Method == "POST" {
        r.ParseForm()
        fmt.Fprintf(w, "%+v\n", r.PostForm)
        for key, val := range r.PostForm {
            data[key] = val[0]
        }
    } else {
        fmt.Fprintf(w, "Sorry we only handle POST requests\n")
    }
}

func readHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Read request from %s!\n", r.URL.Path)
    fmt.Fprintf(w, "%+v\n", data)
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
    data = make(map[string]string)

    fmt.Println("Starting issue tracker version 0.1")
    http.HandleFunc("/", handler)
    http.HandleFunc("/create", createHandler)
    http.HandleFunc("/read", readHandler)
    http.HandleFunc("/update", updateHandler)
    http.HandleFunc("/delete", deleteHandler)
    log.Fatal(http.ListenAndServe(":3000", nil))
}
