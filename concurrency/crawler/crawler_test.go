package crawler

import (
	"fmt"
	"testing"
)

func TestCrawler(t *testing.T) {
	ch := make(chan string)
	go Crawl("http://golang.org/", 4, fetcher, ch)

	// print results from channel
	for value := range ch {
		fmt.Print(value)
	}
}
