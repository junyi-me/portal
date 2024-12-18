package main

import "net/http"

func main() {
    http.HandleFunc("/", handleGetHome)
    http.ListenAndServe(":8080", nil)
}

