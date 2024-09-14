package main

import (
    "fmt"
    "time"
    "net/http"
    "log"
    "path/filepath"
)

// dprint func def delays text print
func dprint(text string, delay time.Duration) {
    for _, letter := range text {
        fmt.Print(string(letter))
        time.Sleep(delay)
    }
}

//logreq func def logs http request
func logreq(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        log.Printf("Request Log: %s %s %s %s", r.Method, r.URL.Path, r.RemoteAddr, time.Since(start))
    })
}

func main() {

    // Port Config
    port := ":80"

    //dprint title
    delay := 50 * time.Millisecond
    dprint(" _   _   _       _   _     \n| |_| |_| |_ ___| |_|_|___ \n|   |  _|  _| . | . | |   |\n|_|_|_| |_| |  _|___|_|_|_|\n            |_|\n       Server Binary", delay)
    fmt.Println()
    fmt.Println()

    //serv root
    dir, err := filepath.Abs(".")
    if err != nil {
        log.Fatal(err)
    }

    //req handler
    fileServer := http.FileServer(http.Dir(dir))
    http.Handle("/", logreq(http.StripPrefix("/", fileServer)))

    //server init
    log.Println("Starting server on ", port)
    err = http.ListenAndServe(port, nil)
    if err != nil {
        log.Fatal(err)
    }
}
