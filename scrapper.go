package main

import (
	"fmt"
	"sync"

	"github.com/gocolly/colly"
)

type Index struct {
	name, url string
}

func collectIndexes() []Index {
	var indices []Index
	var wg sync.WaitGroup

	c := colly.NewCollector(
		colly.AllowedDomains("eltopoblindado.com"),
	)

	wg.Add(1)

	c.OnHTML("li.cat-item", func(e *colly.HTMLElement) {
		indexItem := Index{
			name: e.ChildText("a"),
			url:  e.ChildAttr("a", "href"),
		}
		indices = append(indices, indexItem)
	})

	c.OnScraped(func(r *colly.Response) {
		wg.Done()
	})

	c.Visit("https://eltopoblindado.com/indice/")

	wg.Wait()

	return indices

}

func main() {
	collectedIndices := collectIndexes()

	for _, idx := range collectedIndices {
		fmt.Printf("Name: %s, URL: %s\n", idx.name, idx.url)
	}
}
