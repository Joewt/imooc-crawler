package view

import (
	"testing"
	"html/template"
	"joewt.com/joe/learngo/crawler/frontend/model"
	"os"
)

func TestIndex(t *testing.T) {
	template := template.Must(
		template.ParseFiles("template.html"))

	page := model.SearchResult{}

	err := template.Execute(os.Stdout,page)
	if err != nil {
		panic(err)
	}



}
