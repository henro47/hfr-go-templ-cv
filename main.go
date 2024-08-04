package main

import (
	"log"
	"net/http"
	"os"
	"path"

	"github.com/a-h/templ"
	"github.com/henro47/hfr-go-templ-cv/components"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("couldn't get current working directory %v", err)
	}
	homePage := components.Index()
	pagesHandler := http.NewServeMux()
	pagesHandler.Handle("/media/", http.StripPrefix("/media/", http.FileServer(http.Dir(path.Join(wd, "/media/")))))
	pagesHandler.Handle("/", templ.Handler(homePage))
	log.Fatalln(http.ListenAndServe(":8080", pagesHandler))
}
