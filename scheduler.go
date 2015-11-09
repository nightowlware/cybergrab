package cybergrab

import (
	"fmt"
)

type simpleScheduler struct {
	linkDispenser    linkDispenser
	downloader       downloader
	numPageScrubbers int
}

func newSimpleScheduler(numPageScrubbers int, downloader downloader) *simpleScheduler {
	scheduler := &simpleScheduler{}
	scheduler.downloader = downloader
	scheduler.numPageScrubbers = numPageScrubbers
	return scheduler
}

func (this *simpleScheduler) run(seedUrl string) {
	fmt.Printf("Starting a web-crawl @ seedUrl: %s\n", seedUrl)

	// lazy initialization
	if this.linkDispenser == nil {
		this.linkDispenser = newSimpleLinkMgr()
		this.linkDispenser.pushUrl(seedUrl)
	}

	// launch N PageScrubbers, each in their own goroutine
	for i := 0; i < this.numPageScrubbers; i++ {
		go func() {
			// getUrl() is a blocking receive on a channel
			PageMinion{this}.run(this.linkDispenser.getUrl())
		}()
	}

	// perform downloads from the downloader - blocking until done or timeout.
	this.downloader.processDownloads()
}

func (this *simpleScheduler) stop() {
	fmt.Println("Stopping")
	this.linkDispenser.shutdown()
	this.linkDispenser = nil
}

func (this *simpleScheduler) getLinkDispenser() linkDispenser {
	return this.linkDispenser
}

func (this *simpleScheduler) getDownloader() downloader {
	return this.downloader
}

func (this *simpleScheduler) getCrawlPolicy() CrawlPolicy {
	return simplePolicy{}
}

type simplePolicy struct{}

func (s simplePolicy) ShouldCrawl(url string) bool {
	return true
}

func (s simplePolicy) ShouldDownload(url string) bool {
	return false
}
