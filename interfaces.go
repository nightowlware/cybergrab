package main

type LinkDispenser interface {
	pushUrl(url string)
	getUrl() string
	shutdown()
}

type Scheduler interface {
	run(seedUrl string)
	stop()
}

type Downloader interface {
	processNDownloads(N int)
	addDownload(url string)
	listDownloads() []string
}

type PageScrubber interface {
	run(url string)
}
