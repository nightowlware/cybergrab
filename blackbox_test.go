package cybergrab_test

import (
	_ "fmt"
	"testing"
	"bitbucket/cybergrab"
)

func TestBasic(t *testing.T) {
	var spider cybergrab.Spider

	spider = cybergrab.NewSpider()
	spider.Crawl("www.seed.com")
}
