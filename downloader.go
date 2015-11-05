package main

import (
	"fmt"
	"net/http"
	"io"
	"os"
	"strings"
)

type SimpleDownloader struct {
	folderName string
	urlChannel chan string
	numDownloads int
}

func NewSimpleDownloader(folderName string, numDownloads int) *SimpleDownloader {
	sd := &SimpleDownloader{}
	sd.folderName = folderName
	sd.urlChannel = make(chan string)
	sd.numDownloads = numDownloads

	os.MkdirAll(folderName, 0777)
	return sd
}

// Blocking call - will not return until N pages are downloaded,
// or if urlChannel is closed.
func (this *SimpleDownloader) processDownloads() {
	for i := 0; i < this.numDownloads; i++ {
		this.downloadPage(<-this.urlChannel)
	}
}

func (this *SimpleDownloader) addDownload(url string) {
	this.urlChannel <- url
}

func (this *SimpleDownloader) listDownloads() []string {
	return []string{"not implemented yet"}
}

func (this *SimpleDownloader) downloadPage(url string) {
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
