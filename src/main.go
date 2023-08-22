package main

import (
	"fmt"
	"go-scraper/src/cli"
	"go-scraper/src/config"
	"go-scraper/src/crawler"
	"go-scraper/src/lib"
	"log"
	"os"
)

func RunCrawler() {
	for _, domain := range config.GetDomains() {
		log.Printf("Crawling through %s...\n", domain.Title)
		title, err := crawler.Run(domain)
		if err != nil {
			panic(err)
		}

		log.Printf("Value is: %s", title)
	}
}

func RunFetch() {
	data, err := lib.FetchNews(lib.SearchQuery{
		Subject: "startups AND Latin American",
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Data: %s", data.Status)
}

func main() {
	defer os.Exit(0)

	err := config.Load()
	if err != nil {
		panic(err)
	}

	cli := cli.CommandLine{}
	cli.Run()
}
