package main

import (
    "fmt"
    "log"
    "net/http"
)

type Bounds struct {
    topLeft int
    botRight int
}

func (b *Bounds) String() string {
    return fmt.Sprintf("{%v, %v}", b.topLeft, b.botRight)
}

func parseBounds(boundsS string) (Bounds, error) {
    return Bounds{0, 0}, nil
}
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!\n", r.URL.Path[1:])
    bounds, err := parseBounds(r.URL.Path[1:])
    if err != nil {
        fmt.Fprintf(w, "error\n")
    }
    fmt.Fprintf(w, "%v\n", bounds)
}

func main() {
    fmt.Println("Starting house-viewer version 0.1")
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":3000", nil))
}
