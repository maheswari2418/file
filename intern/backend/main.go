package main

import (
    "log"
    "net/http"
)

func main() {
    // Serve the index.html file at the root URL
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "index.html")
    })

    // Serve the styles.css file
    http.HandleFunc("/styles.css", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "styles.css")
    })

    // Serve the script.js file
    http.HandleFunc("/script.js", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "script.js")
    })

    log.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8081", nil))
}