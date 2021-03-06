package cybergrab

///////////////////////////////
/// Public Interfaces
///////////////////////////////

type CrawlPolicy interface {
	ShouldDownload(url string) bool
	ShouldCrawl(url string) bool
}

type Spider interface {
	Crawl(seedURL string) error
}

///////////////////////////////

type linkDispenser interface {
	pushUrl(url string)
	getUrl() string
	shutdown()
}

type scheduler interface {
	run(seedUrl string)
	stop()
	getLinkDispenser() linkDispenser
	getDownloader() downloader
	getCrawlPolicy() CrawlPolicy
}

type downloader interface {
	processDownloads()
	addDownload(url string)
	listDownloads() []string
	getNumDownloads() uint
	shutdown()
}

type pageScrubber interface {
	run(url string) error
}
