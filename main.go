package main

import (
  "fmt"
  "net/http"
)

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "HEEEEEEE!!")
  })

  fmt.Println("Server is running on port 1212...")
  http.ListenAndServe(":1212", nil)
}
