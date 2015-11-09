package cybergrab

import (
	_ "fmt"
)

const (
	channel_buffer_size = 100
)

type simpleLinkMgr struct {
	urls chan string
}

func newSimpleLinkMgr() *simpleLinkMgr {
	l := &simpleLinkMgr{}
	l.urls = make(chan string, channel_buffer_size)

	return l
}

func (this *simpleLinkMgr) pushUrl(url string) {
	this.urls <- url
}

func (this *simpleLinkMgr) getUrl() string {
	return <-this.urls
}

func (this *simpleLinkMgr) shutdown() {
	close(this.urls)
}
