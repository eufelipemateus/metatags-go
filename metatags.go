package metatags

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Get html from remote url
func getHtmmlFromURL(url string) string {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	content, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	html := string(content)

	return html
}

// Get metatags list from html text.
func GetMetatagsFromHTML(html string) map[string]string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader((html)))
	if err != nil {
		log.Fatal(err)
	}

	metas := make(map[string]string)

	doc.Find("meta").Each(func(i int, s *goquery.Selection) {
		if name, exists := s.Attr("name"); exists {
			value, _ := s.Attr("content")
			metas[name] = value
		}
	})

	return metas
}

// Get metatagsb list from url
func GetMetatagsFromURL(url string) map[string]string {
	html := getHtmmlFromURL(url)
	return GetMetatagsFromHTML(html)
}
