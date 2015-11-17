package cybergrab

import (
	"fmt"
	"sync"
)

type simpleScheduler struct {
	linkDispenser    linkDispenser
	downloader       downloader
	numPageScrubbers uint
	policy           CrawlPolicy
}

func newSimpleScheduler(cp CrawlPolicy, numPageScrubbers uint, downloader downloader) *simpleScheduler {
	scheduler := &simpleScheduler{}
	scheduler.downloader = downloader
	scheduler.numPageScrubbers = numPageScrubbers
	scheduler.policy = cp
	return scheduler
}

func (ss *simpleScheduler) run(seedUrl string) {
	fmt.Printf("Starting a web-crawl @ seedUrl: %s\n", seedUrl)

	// lazy initialization
	if ss.linkDispenser == nil {
		ss.linkDispenser = newSimpleLinkMgr(ss.downloader.getNumDownloads())
		ss.linkDispenser.pushUrl(seedUrl)
	}

	// asynchronously launch N PageScrubbers, each in their own goroutine
	go func() {
		waitGroup := sync.WaitGroup{}
		waitGroup.Add(int(ss.numPageScrubbers))

		for i := uint(0); i < ss.numPageScrubbers; i++ {
			go func() {
				// getUrl() is a (timeout) blocking receive on a channel
				pageMinion{ss}.run(ss.linkDispenser.getUrl())
				waitGroup.Done()
			}()
		}

		// wait until all the scrubbers have finished running,
		// then shutdown the linkDispenser, causing the downloader to
		// quit (safely).
		waitGroup.Wait()
		ss.stop()
	}()

	// synchronously perform downloads from the downloader - blocking until done or timeout.
	ss.downloader.processDownloads()
}

func (ss *simpleScheduler) stop() {
	fmt.Println("Stopping Scheduler")

	ss.linkDispenser.shutdown()
	ss.linkDispenser = nil

	ss.downloader.shutdown()
	ss.downloader = nil
}

func (ss *simpleScheduler) getLinkDispenser() linkDispenser {
	return ss.linkDispenser
}

func (ss *simpleScheduler) getDownloader() downloader {
	return ss.downloader
}

func (ss *simpleScheduler) getCrawlPolicy() CrawlPolicy {
	return ss.policy
}
