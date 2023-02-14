package html_parser

import (
	"errors"

	"github.com/gocolly/colly/v2"
)

const (
	NOT_MODIFIED = 304
)

func ParseHtmlTable(link string) (table [][]string, err error) {
	collector := colly.NewCollector()
	collector.OnResponse(func(r *colly.Response) {

		if r.StatusCode == NOT_MODIFIED {
			table, err = nil, errors.New("NOT_MODIFIED")
			return
		}
		collector.OnHTML("table.confluenceTable", func(e *colly.HTMLElement) {
			colsCount := len(e.ChildTexts("th"))
			table = make([][]string, colsCount)

			e.ForEach("tr", func(i int, el *colly.HTMLElement) {
				for k, v := range el.ChildTexts("th") {
					table[k] = append(table[k], v)
				}

				for k, v := range el.ChildTexts("td") {
					table[k] = append(table[k], v)
				}
			})
		})
	})

	collector.Visit(link)
	return
}
