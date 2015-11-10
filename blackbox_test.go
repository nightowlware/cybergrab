package cybergrab_test

import (
	"bitbucket/cybergrab"
	_ "fmt"
	"testing"
)

type simpleCrawlPolicy struct {
}

func (scp simpleCrawlPolicy) ShouldDownload(url string) bool {
	return true
}

func (scp simpleCrawlPolicy) ShouldCrawl(url string) bool {
	return true
}

func TestBasic(t *testing.T) {
	var spider cybergrab.Spider

	spider, err := cybergrab.NewSpider(simpleCrawlPolicy{}, 10, 10)
	if err == nil {
		spider.Crawl("http://www.cnn.com")
	}
}
