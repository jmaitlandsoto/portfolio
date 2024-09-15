package main

import (
    "log"
    "net/http"

    "github.com/jmaitlandsoto/portfolio"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", handlers.Home)
    mux.HandleFunc("/projects", handlers.Projects)

    // Serve static files
    fileServer := http.FileServer(http.Dir("./static"))
    mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

    log.Println("Starting server on :8080")
    err := http.ListenAndServe(":8080", mux)
    if err != nil {
        log.Fatal(err)
    }
}
