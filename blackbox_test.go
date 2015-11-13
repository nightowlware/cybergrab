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

	spider, err := cybergrab.NewSpider(simpleCrawlPolicy{}, 100, 100)
	if err == nil {
		spider.Crawl("http://www.cnn.com/2015/11/12/politics/u-s-airstrike-targets-jihadi-john-in-syria/index.html")
	}
}
