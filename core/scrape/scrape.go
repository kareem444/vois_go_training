package scrape

import (
	"fmt"
	"log"
	"sync"

	"github.com/gocolly/colly"
)

type DataToScrap struct {
	Urls          []string
	TitleSelector string
}

type Scraper struct {
	DataToScrap DataToScrap
	Mutex       *sync.Mutex
}

type ScrappingResponse struct {
	Title    string `json:"titles"`
	Url      string `json:"url"`
	Selector string `json:"selector"`
}

func ScrapeURLs(s Scraper) ([]ScrappingResponse, error) {
	data := s.DataToScrap
	mutex := s.Mutex
	var result []ScrappingResponse

	c := colly.NewCollector(
		colly.Async(true),
	)

	c.OnHTML(data.TitleSelector, func(e *colly.HTMLElement) {
		mutex.Lock()
		defer mutex.Unlock()

		if e.Text == "" {
			return
		}

		result = append(result, ScrappingResponse{
			Title:    e.Text,
			Url:      e.Request.URL.String(),
			Selector: data.TitleSelector,
		})
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Error:", err)
	})

	for _, url := range data.Urls {
		err := c.Visit(url)
		if err != nil {
			return result, fmt.Errorf("failed to visit URL %s: %w", url, err)
		}
	}

	c.Wait()
	return result, nil
}
