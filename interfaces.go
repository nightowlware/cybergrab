package main

type LinkDispenser interface {
	pushUrl(url string)
	getUrl() string
	shutdown()
}

type Scheduler interface {
	run(seedUrl string)
	stop()
	getLinkDispenser() LinkDispenser
	getDownloader() Downloader
	getCrawlPolicy() CrawlPolicy
}

type Downloader interface {
	processDownloads()
	addDownload(url string)
	listDownloads() []string
}

type PageScrubber interface {
	run(url string) error
}

type CrawlPolicy interface {
	ShouldDownload(url string) bool
	ShouldCrawl(url string) bool
}
