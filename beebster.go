package main

import (
        //        "./isearch"
        "html/template"
        "io"
        "net/http"
)

func aboutHandler(w http.ResponseWriter, r *http.Request) {
        //t, _ := template.New("some template").ParseFile("about.html")
        //        t, _ = t.ParseFiles("about.html", nil)
        t.Execute(w, abtemplate)
}
func main() {
        http.HandleFunc("/about/", aboutHandler)
        http.ListenAndServe(":4242", nil)
        //      isearch.isearch()

}
