package main

import (
	"net/http"
	"joewt.com/joe/learngo/crawler/frontend/controler"
)

func main() {
	http.Handle("/",http.FileServer(
		http.Dir("crawler/frontend/view"),
	))
	http.Handle("/search",
		controler.CreateSearchResultHandler(
			"crawler/frontend/view/template.html",
		))
	err := http.ListenAndServe(":80",nil)
	if err != nil {
		panic(err)
	}
}
