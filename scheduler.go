package cybergrab

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type simpleScheduler struct {
	linkDispenser    linkDispenser
	downloader       downloader
	numPageScrubbers uint
	policy           CrawlPolicy
	_count           *int32
}

func newSimpleScheduler(cp CrawlPolicy, numPageScrubbers uint, downloader downloader) *simpleScheduler {
	scheduler := &simpleScheduler{}
	scheduler.downloader = downloader
	scheduler.numPageScrubbers = numPageScrubbers
	scheduler.policy = cp
	scheduler._count = new(int32)
	*scheduler._count = 0
	return scheduler
}

func (this *simpleScheduler) run(seedUrl string) {
	fmt.Printf("Starting a web-crawl @ seedUrl: %s\n", seedUrl)

	// lazy initialization
	if this.linkDispenser == nil {
		this.linkDispenser = newSimpleLinkMgr(this.downloader.getNumDownloads())
		this.linkDispenser.pushUrl(seedUrl)
	}

	// asynchronously launch N PageScrubbers, each in their own goroutine
	go func() {
		waitGroup := sync.WaitGroup{}
		waitGroup.Add(int(this.numPageScrubbers))

		for i := uint(0); i < this.numPageScrubbers; i++ {
			go func() {
				// getUrl() is a (timeout) blocking receive on a channel
				pageMinion{this}.run(this.linkDispenser.getUrl())
				waitGroup.Done()
				atomic.AddInt32(this._count, 1)
				fmt.Println("COUNT OF DONES:::::::::::::::::::::::::::::::", *this._count)
			}()
		}

		// wait until all the scrubbers have finished running,
		// then shutdown the linkDispenser, causing the downloader to
		// quit (safely).
		waitGroup.Wait()
		this.stop()
	}()

	// synchronously perform downloads from the downloader - blocking until done or timeout.
	this.downloader.processDownloads()
}

func (this *simpleScheduler) stop() {
	fmt.Println("Stopping Scheduler")

	this.linkDispenser.shutdown()
	this.linkDispenser = nil

	this.downloader.shutdown()
	this.downloader = nil
}

func (this *simpleScheduler) getLinkDispenser() linkDispenser {
	return this.linkDispenser
}

func (this *simpleScheduler) getDownloader() downloader {
	return this.downloader
}

func (this *simpleScheduler) getCrawlPolicy() CrawlPolicy {
	return this.policy
}
