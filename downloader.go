package cybergrab

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
)

type simpleDownloader struct {
	folderName   string
	urlChannel   chan string
	numDownloads uint
}

func newSimpleDownloader(folderName string, numDownloads uint) *simpleDownloader {
	sd := &simpleDownloader{}
	sd.folderName = folderName
	sd.urlChannel = make(chan string)
	sd.numDownloads = numDownloads

	os.MkdirAll(folderName, 0777)
	return sd
}

// Blocking call - will not return until N pages are downloaded,
// or if urlChannel is closed.
// Reads numDownloads' worth of Urls, and downloads each Url
// in its separate goroutine.
func (this *simpleDownloader) processDownloads() {
	var wait_group sync.WaitGroup

	for i := uint(0); i < this.numDownloads; i++ {
		wait_group.Add(1)
		go func() {
			this.downloadUrl(<-this.urlChannel)
		}()
	}

	// wait for all the downloading goroutines to finish
	wait_group.Wait()
}

func (this *simpleDownloader) addDownload(url string) {
	this.urlChannel <- url
}

func (this *simpleDownloader) listDownloads() []string {
	return []string{"not implemented yet"}
}

func (this *simpleDownloader) downloadUrl(url string) {
	fmt.Println("Downloading page: ", url)

	arr := strings.Split(url, "/")
	name := this.folderName + "/" + arr[len(arr)-1]

	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println(err)
	}
}
