package main

import (
	"log"

	"go-scraper/src/config"
	"go-scraper/src/crawler"
)

func main() {
	err := config.Load()
	if err != nil {
		panic(err)
	}

	title, err := crawler.Run("https://progy.com.br")
	if err != nil {
		panic(err)
	}

	log.Printf("Finish with value: %s", title)
}
