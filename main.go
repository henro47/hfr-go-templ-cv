package main

import (
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/henro47/hfr-go-templ-cv/components"
)

func main() {
	homePage := components.Index()
	pagesHandler := http.NewServeMux()
	pagesHandler.Handle("/", templ.Handler(homePage))
	log.Fatalln(http.ListenAndServe(":4000", pagesHandler))
}
