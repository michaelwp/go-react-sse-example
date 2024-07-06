package main

import (
    "fmt"
    "net/http"
    "time"
)

func sseHandler(w http.ResponseWriter, r *http.Request) {
    // Set headers for SSE
    w.Header().Set("Content-Type", "text/event-stream")
    w.Header().Set("Cache-Control", "no-cache")
    w.Header().Set("Connection", "keep-alive")

    flusher, ok := w.(http.Flusher)
    if !ok {
        http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
        return
    }

    // Send initial message
    fmt.Fprintf(w, "data: %s\n\n", "Connected to SSE server")
    flusher.Flush()

    ticker := time.NewTicker(5 * time.Second)
    defer ticker.Stop()

    for {
        select {
        case t := <-ticker.C:
            fmt.Fprintf(w, "data: %s\n\n", t.String())
            flusher.Flush()
        case <-r.Context().Done():
            return
        }
    }
}

func main() {
    // Serve the static files from the React app
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/", fs)

    // SSE endpoint
    http.HandleFunc("/events", sseHandler)

    fmt.Println("Server started at :8080")
    http.ListenAndServe(":8080", nil)
}