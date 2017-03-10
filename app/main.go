package main
import (
    "fmt"
    "log"
"net/http" )
func messageHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hello NY Times")
}
func main() {
    http.HandleFunc("/hello", messageHandler)
    log.Println("Listening...")
    http.ListenAndServe(":8080", nil)
}

