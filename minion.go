package main

import (
	"fmt"
)

type PageMinion struct {
	downloader Downloader
	linkdispenser LinkDispenser
}

func NewPageMinion(l LinkDispenser, d Downloader) *PageMinion {
	m := &PageMinion{}
	m.downloader = d
	m.linkdispenser = l
	return m
}

func invalidUrl(url string) bool {
	return url == ""
}

// This function *must* feed (via addDownload) the downloader at least one url to download,
// otherwise the downloader will block indefinitely for it to supply one (bad).
func (this *PageMinion) run(url string) {
	fmt.Println("PageMinion: Scrubbing page: " + url)

	if invalidUrl(url) {
		this.downloader.addDownload("INVALID")
	} else {
		this.downloader.addDownload(url)
	}

	this.linkdispenser.pushUrl("http://www.reddit.com")
}
