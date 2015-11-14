package cybergrab_test

import (
	"bitbucket/cybergrab"
	_ "fmt"
	"strings"
	"testing"
)

type simpleCrawlPolicy struct {
}

func (scp simpleCrawlPolicy) ShouldDownload(url string) bool {
	return strings.HasSuffix(url, "jpg") ||
		strings.HasSuffix(url, "gif")
}

func (scp simpleCrawlPolicy) ShouldCrawl(url string) bool {
	return true
}

func TestBasic(t *testing.T) {
	var spider cybergrab.Spider

	spider, err := cybergrab.NewSpider(simpleCrawlPolicy{}, 10000, 10000)
	if err == nil {
		spider.Crawl("http://www.pawg.site/")
	}
}
