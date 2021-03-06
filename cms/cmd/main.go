package main

import (
  "net/http"
  "block-go/cms"
)

func main() {
  http.HandleFunc("/", cms.ServeIndex)
  http.HandleFunc("/new", cms.HandleNew)
  http.HandleFunc("/page/", cms.ServePage)
  http.HandleFunc("/post/", cms.ServePost)

  http.ListenAndServe(":3000", nil)
}
