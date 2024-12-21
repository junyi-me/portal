package main

import "net/http"

func handleGetHome(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        w.WriteHeader(http.StatusOK)
        renderHome(w)
    default:
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusMethodNotAllowed)
        w.Write([]byte(`{"message": "Method not allowed"}`))
    }
}

func handleComments(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        w.WriteHeader(http.StatusOK)
        renderComments(w)
    default:
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusMethodNotAllowed)
        w.Write([]byte(`{"message": "Method not allowed"}`))
    }
}

