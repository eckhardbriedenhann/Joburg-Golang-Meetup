package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alecthomas/template"
)

var (
	pagesTemp *template.Template
	indexTemp *template.Template
	address   = ":8081"
)

type Page struct {
	Message string
	Picture string
}

type indexRow struct {
	Display string
	Link    string
}

type Route struct {
	Endpoint string
	Temp     Page
}

var endpoints = []Route{
	Route{Endpoint: "/mindblown", Temp: Page{Message: "Boom!!!", Picture: "https://media.giphy.com/media/xT0xeJpnrWC4XWblEk/giphy.gif"}},
	Route{Endpoint: "/fin", Temp: Page{Message: "A bid you farewell.", Picture: "https://media.giphy.com/media/3otPoUkg3hBxQKRJ7y/giphy.gif"}},
}

func handleMagic(p Page) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pagesTemp.Execute(w, p)
	}
}

func main() {

	fmt.Printf("info: Starting server at %s", address)
	pagesTemp = template.Must(template.ParseFiles("template/page.html"))
	indexTemp = template.Must(template.ParseFiles("template/index.html"))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("template/images"))))
	var links []indexRow
	for i, e := range endpoints {
		links = append(links, indexRow{Display: fmt.Sprintf("Endpoint %d: %s", i, e.Endpoint), Link: fmt.Sprintf("http://localhost%s%s", address, e.Endpoint)})
		http.HandleFunc(e.Endpoint, handleMagic(e.Temp))
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		indexTemp.Execute(w, links)
	})
	log.Fatal(http.ListenAndServe(address, nil))

}
