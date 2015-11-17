package cybergrab

import (
	_ "fmt"
	"time"
)

type simpleLinkMgr struct {
	urls chan string
}

func newSimpleLinkMgr(bufferSize uint) *simpleLinkMgr {
	l := &simpleLinkMgr{}
	l.urls = make(chan string, 100*bufferSize)

	return l
}

func (slm *simpleLinkMgr) pushUrl(url string) {
	timeout := createTimeout(WORKER_TIMEOUT_SECONDS)

	// skip the receive on the urls channel if it takes too long
	select {
	case slm.urls <- url:
	case <-timeout:
	}
}

// getUrl was initially an indefinitely-blocking receive on a channel,
// but that was changed later to incorporate a timeout.
func (slm *simpleLinkMgr) getUrl() string {

	timeout := createTimeout(WORKER_TIMEOUT_SECONDS)

	select {
	case url := <-slm.urls:
		return url
	case <-timeout:
		// return an invalid url in the case of a timeout on a channel receive
		return ""
	}
}

func (slm *simpleLinkMgr) shutdown() {
	close(slm.urls)
}

// createTimeout returns a simple channel that will have a value
// inserted into it after seconds amount of time
func createTimeout(seconds uint) <-chan bool {
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(time.Duration(seconds) * time.Second)
		timeout <- true
	}()

	return timeout
}
