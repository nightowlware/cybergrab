package cybergrab

import (
	_ "fmt"
	"testing"
)

func TestBasic(t *testing.T) {
	var engine Scheduler
	var downloader Downloader

	downloader = NewSimpleDownloader("downloads", 10)
	engine = NewSimpleScheduler(10, downloader)

	engine.run("http://www.cnn.com")
}
