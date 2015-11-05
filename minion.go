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

// Instructs this minion to process the given url. 
//
// This minion can do any combination of:
// 1. Push a new url into the download queue
// 2. Push a new url into the link dispenser queue, for other minions to run()
// 3. Do nothing
func (this *PageMinion) run(url string) {
	fmt.Println("PageMinion: Scrubbing page: " + url)

	if invalidUrl(url) {
		fmt.Println("URL <", url, "> is invalid, ignoring.")
	} else {
		this.downloader.addDownload(url)
	}

	this.linkdispenser.pushUrl("http://www.reddit.com")
}
