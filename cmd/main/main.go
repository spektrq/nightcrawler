package main

import (
	"sync"

	crawler "github.com/spektrq/nightcrawler/internal/crawler"
)

func main() {
	sitesChannel := make(chan string)
	crawedLinksChannel := make(chan string)
	pendingCountChannel := make(chan int)

	siteToCrawl := "https://crawler-test.com/"

	go func() {
		crawedLinksChannel <- siteToCrawl
	}()

	var wg sync.WaitGroup

	go crawler.ProcessCrawledLinks(sitesChannel, crawedLinksChannel, pendingCountChannel)
	go crawler.MonitorCrawling(sitesChannel, crawedLinksChannel, pendingCountChannel)

	var numCrawlerThreads = 100
	for i := 0; i < numCrawlerThreads; i++ {
		wg.Add(1)
		go crawler.CrawlWebpage(&wg, sitesChannel, crawedLinksChannel, pendingCountChannel)
	}

	wg.Wait()
}
