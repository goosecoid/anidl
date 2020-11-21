package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/spf13/cobra"
)

type Show struct {
	Name string
	URL  string
}

func init() {
	rootCmd.AddCommand(searchCommand)
}

var (
	searchCommand = &cobra.Command{
		Use:   "search",
		Short: "Search command",
		Long:  "Enter the name of series/movie to search",
		Run: func(cmd *cobra.Command, args []string) {

			keyword, _ := cmd.Flags().GetString("keyword")

			output := search(keyword)

			for k := range output {
				count++
				fmt.Fprintf(os.Stdout, "[%d]: %s\n", count, k)
			}
		},
	}

	count int

	keyword string
)

func init() {
	searchCommand.Flags().StringVarP(&keyword, "keyword", "k", "", "Key to search with")
}

func search(keyword string) (output map[string]Show) {
	c := colly.NewCollector()

	shows := make(map[string]Show)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		title := e.Attr("title")

		if !strings.HasPrefix(link, "/anime/") {
			return
		}

		if !strings.Contains(strings.ToLower(title), keyword) {
			return
		}

		if len(shows) == 10 {
			return
		}

		shows[title] = Show{
			Name: title,
			URL:  link,
		}
	})

	c.Visit(fmt.Sprintf("https://www4.gogoanime.pro/search?keyword=%s", keyword))

	return shows
}
