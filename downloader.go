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
	// make sure the url buffer is big enough so that none of the worker
	// go routines block on input.
	sd.urlChannel = make(chan string, MAX_WORKERS)
	sd.numDownloads = numDownloads

	os.MkdirAll(folderName, 0777)
	return sd
}

func (sd *simpleDownloader) shutdown() {
	close(sd.urlChannel)
}

func (sd *simpleDownloader) getNumDownloads() uint {
	return sd.numDownloads
}

// Blocking call - will not return until N pages are downloaded,
// or if urlChannel is closed.
// Reads numDownloads' worth of Urls, and downloads each Url
// in its separate goroutine.
func (sd *simpleDownloader) processDownloads() {
	var wait_group sync.WaitGroup

	for i := uint(0); i < sd.numDownloads; i++ {
		wait_group.Add(1)
		go func() {
			sd.downloadUrl(<-sd.urlChannel)
			wait_group.Done()
		}()
	}

	// wait for all the downloading goroutines to finish
	wait_group.Wait()
}

func (sd *simpleDownloader) addDownload(url string) {
	sd.urlChannel <- url
}

func (sd *simpleDownloader) listDownloads() []string {
	return []string{"not implemented yet"}
}

func (sd *simpleDownloader) downloadUrl(url string) {
	arr := strings.Split(url, "/")
	name := sd.folderName + "/" + arr[len(arr)-1]

	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// attempt to create a file based on name
	file, err := os.Create(name)
	if err != nil {
		return
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Downloaded url: ", url)
	}
}
