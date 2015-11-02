package main

import (
	"fmt"
)

type PageMinion struct {
	downloader Downloader
}

func NewPageMinion(d Downloader) *PageMinion {
	m := &PageMinion{}
	m.downloader = d
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
}
