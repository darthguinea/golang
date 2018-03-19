package main

import (
  "io"
  "fmt"
  "net/http"
)

func main() {
  serverPort := "8080"
  http.HandleFunc("/", root)
  http.HandleFunc("/run", run)
  if err := http.ListenAndServe(":" + serverPort, nil);err != nil {
  }
}

func root(w http.ResponseWriter, r *http.Request) {
  io.WriteString(w, "Hello world!")
}

func run(w http.ResponseWriter, r *http.Request) {
}
