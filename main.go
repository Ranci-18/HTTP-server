package main

import (
       "fmt"
       "net/http"
)

func main() {
     // define a handler function
     http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
         fmt.Fprintf(w, "Hello, this is a Go web server!")
    })

    // start the server on port 3002
    fmt.Println("Starting server at : 3002")
    err := http.ListenAndServe(":3002", nil)
    if err != nil {
        fmt.Println("Error starting server:", err)
    }
}