package cybergrab

import ()

const (
	channel_buffer_size = 100
)

type LinkMgr struct {
	urls chan string
}

func NewLinkMgr() *LinkMgr {
	l := &LinkMgr{}
	l.urls = make(chan string, channel_buffer_size)

	return l
}

func (l *LinkMgr) pushUrl(url string) {
	l.urls <- url
}

func (l *LinkMgr) getUrl() string {
	return <- l.urls
}

