package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

type Result struct {
  URL string;
  Title string;
}

var (
  target string
  results []Result
)

func filterDuplicates() {
  keys := make(map[string]bool)
  tempList := []Result{}

  for _, item := range results {
    if _, value := keys[item.URL]; !value{
      keys[item.URL] = true
      tempList = append(tempList, item)
    }
  }
  results = tempList
}

func getUrls() {
  c := colly.NewCollector()
  size := 0
  results = make([]Result, size)

  c.OnHTML("a[href]", func (e *colly.HTMLElement) {
    if strings.Contains(e.Attr("href"), target) && !strings.Contains(e.Attr(("href")), "search") {
      size++
      results = append(results, Result{Title: e.Attr("title"), URL: e.Request.AbsoluteURL(e.Attr("href"))})
    }
  })

  c.Visit(fmt.Sprintf("https://gogoanime.pro/search?keyword=%s", target))
}

func main() {
  flag.StringVar(&target, "name", "", "Name of the show")
  flag.Parse()

  getUrls()
  filterDuplicates()

  for index, item := range results {
    fmt.Printf("%d: %s\n", index+1, item.Title)
  }
}
