package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", hello)
    log.Fatal(http.ListenAndServe(":8079", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello World2!")
}

