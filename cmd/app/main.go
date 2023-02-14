package main

import (
	"log"
	"tableParser/internal/google_doc"
	"tableParser/internal/html_parser"
)

const (
	WEB_LINK = "https://confluence.hflabs.ru/pages/viewpage.action?pageId=1181220999"
)

func main() {
	contentTable, err := html_parser.ParseHtmlTable(WEB_LINK)
	if err != nil {
		log.Fatal(err.Error())
	}

	google_doc.MakeTableOnGoogleDocs(contentTable)
}
