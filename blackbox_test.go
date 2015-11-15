package cybergrab_test

import (
	"bitbucket/cybergrab"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	"testing"
)

type simpleCrawlPolicy struct {
}

func (scp simpleCrawlPolicy) ShouldDownload(url string) bool {
	return strings.HasSuffix(url, "jpg") ||
		strings.HasSuffix(url, "gif") ||
		strings.HasSuffix(url, "pdg")
}

func (scp simpleCrawlPolicy) ShouldCrawl(url string) bool {
	return true
}

func _TestMaxGoRoutines(t *testing.T) {
	var spider cybergrab.Spider

	//spider, err := cybergrab.NewSpider(simpleCrawlPolicy{}, cybergrab.MAX_DOWNLOADS, cybergrab.MAX_WORKERS)
	spider, err := cybergrab.NewSpider(simpleCrawlPolicy{}, cybergrab.MAX_DOWNLOADS, 2000)
	if err == nil {
		spider.Crawl("http://phapit.com")
	} else {
		fmt.Println(err)
	}
}

func TestDownloaderExits(t *testing.T) {
	var spider cybergrab.Spider

	spider, err := cybergrab.NewSpider(simpleCrawlPolicy{}, cybergrab.MAX_DOWNLOADS, 10000)
	if err == nil {
		spider.Crawl("http://www.porngifs.com")
	} else {
		fmt.Println(err)
	}

	go func() {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGQUIT)
		buf := make([]byte, 1<<20)
		for {
			<-sigs
			runtime.Stack(buf, true)
			log.Printf("=== received SIGQUIT ===\n*** goroutine dump...\n%s\n*** end\n", buf)
		}
	}()
}
