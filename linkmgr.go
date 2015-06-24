package main

import (
	//"fmt"
)

const (
	channel_buffer_size = 100
)

type SimpleLinkMgr struct {
	urls chan string
}

func NewSimpleLinkMgr() *SimpleLinkMgr {
	l := &SimpleLinkMgr{}
	l.urls = make(chan string, channel_buffer_size)

	return l
}

func (this *SimpleLinkMgr) pushUrl(url string) {
	this.urls <- url
}

func (this *SimpleLinkMgr) getUrl() string {
	return <-this.urls
}

func (this *SimpleLinkMgr) shutdown() {
	close(this.urls)
}
