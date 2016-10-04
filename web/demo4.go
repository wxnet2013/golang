package main

import (
  "fmt"
  "net/http"
  "log"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
  http.HandleFunc("/", handler)
  err := http.ListenAndServe("127.0.0.1:9000", nil)
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}
