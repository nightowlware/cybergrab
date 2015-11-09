package cybergrab

import (
	_ "fmt"
	"testing"
)

func _TestBasic(t *testing.T) {
	var engine scheduler
	var downloader downloader

	downloader = newSimpleDownloader("downloads", 10)
	engine = newSimpleScheduler(10, downloader)

	engine.run("http://www.cnn.com")
}
