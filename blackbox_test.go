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

type simpleCrawlPolicy struct{}

func (scp simpleCrawlPolicy) ShouldDownload(url string) bool {
	return strings.HasSuffix(url, "jpg") ||
		strings.HasSuffix(url, "gif") ||
		strings.HasSuffix(url, "pdg")
}

func (scp simpleCrawlPolicy) ShouldCrawl(url string) bool {
	return !scp.ShouldDownload(url)
}

/////////////////////////////////////////////////////////////

// TestMain is a setup function for the rest of the
// tests in this file.
func TestMain(m *testing.M) {
	// setup a signal handler so that "Ctrl-\" shows a stack-trace.
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

	os.Exit(m.Run())
}

func TestMaxGoRoutines(t *testing.T) {
	var spider cybergrab.Spider

	spider, err := cybergrab.NewSpider(simpleCrawlPolicy{}, cybergrab.MAX_DOWNLOADS, cybergrab.MAX_WORKERS)
	if err == nil {
		spider.Crawl("http://www.salami.com")
	} else {
		fmt.Println(err)
	}
}

func TestMaxWorkers(t *testing.T) {
	var spider cybergrab.Spider

	spider, err := cybergrab.NewSpider(simpleCrawlPolicy{}, 100, cybergrab.MAX_WORKERS)
	if err == nil {
		spider.Crawl("http://www.quickmeme.com")
	} else {
		fmt.Println(err)
	}
}

func TestMaxDownloads(t *testing.T) {
	var spider cybergrab.Spider

	spider, err := cybergrab.NewSpider(simpleCrawlPolicy{}, cybergrab.MAX_DOWNLOADS, 1000)
	if err == nil {
		spider.Crawl("http://www.spam.com")
	} else {
		fmt.Println(err)
	}
}
