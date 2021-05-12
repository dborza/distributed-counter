package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
    "sync/atomic"
)

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })

    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hi")
    })

    var ops uint64

    http.HandleFunc("/inc", func(w http.ResponseWriter, r *http.Request){
        atomic.AddUint64(&ops, 1)
        fmt.Fprintf(w, "Count: %d", ops)
    })

    http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Count: %d", ops)
    })

    log.Fatal(http.ListenAndServe(":8081", nil))

}
