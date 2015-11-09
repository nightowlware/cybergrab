package cybergrab

import (
	"fmt"
)

type SimpleScheduler struct {
	linkDispenser    LinkDispenser
	downloader       Downloader
	numPageScrubbers int
}

func NewSimpleScheduler(numPageScrubbers int, downloader Downloader) *SimpleScheduler {
	scheduler := &SimpleScheduler{}
	scheduler.downloader = downloader
	scheduler.numPageScrubbers = numPageScrubbers
	return scheduler
}

func (this *SimpleScheduler) run(seedUrl string) {
	fmt.Printf("Starting a web-crawl @ seedUrl: %s\n", seedUrl)

	// lazy initialization
	if this.linkDispenser == nil {
		this.linkDispenser = NewSimpleLinkMgr()
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

func (this *SimpleScheduler) stop() {
	fmt.Println("Stopping")
	this.linkDispenser.shutdown()
	this.linkDispenser = nil
}

func (this *SimpleScheduler) getLinkDispenser() LinkDispenser {
	return this.linkDispenser
}

func (this *SimpleScheduler) getDownloader() Downloader {
	return this.downloader
}

func (this *SimpleScheduler) getCrawlPolicy() CrawlPolicy {
	return simplePolicy{}
}

type simplePolicy struct{}

func (s simplePolicy) ShouldCrawl(url string) bool {
	return true
}

func (s simplePolicy) ShouldDownload(url string) bool {
	return false
}
