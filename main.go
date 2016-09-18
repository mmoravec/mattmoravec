package main

import (
    "net/http"
    "html/template"
    "fmt"
    "os"
)

func main() {
  port := os.Getenv("PORT")

  // if port == "" {
  //   log.Fatal("$PORT must be set")
  // }

  http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.URL)
    tmpl :=  template.Must(template.ParseFiles("index.html"))
    tmpl.Execute(w, nil)
  })
  http.ListenAndServe(":" + port, nil)
}
