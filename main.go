package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/russross/blackfriday"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/markdown", generateMarkdown)
	http.Handle("/", http.FileServer(http.Dir("public")))

	fmt.Printf("Listening on port %s...", port)
	http.ListenAndServe(":"+port, nil)
}

func generateMarkdown(rw http.ResponseWriter, r *http.Request) {
	body := r.FormValue("body")
	markdown := blackfriday.MarkdownCommon([]byte(body))
	rw.Write(markdown)
}
