package main

import (
	"fmt"
)

type SimpleDownloader struct {
	folderName string
	urlChannel chan string
}

func NewSimpleDownloader(folderName string) *SimpleDownloader {
	sd := &SimpleDownloader{}
	sd.folderName = folderName
	sd.urlChannel = make(chan string)
	return sd
}

// Blocking call - will not return until N pages are downloaded
func (this *SimpleDownloader) processNDownloads(N int) {
	for i := 0; i < N; i++ {
		downloadPage(<-this.urlChannel)
	}
}

func (this *SimpleDownloader) addDownload(url string) {
	this.urlChannel <- url
}

func (this *SimpleDownloader) listDownloads() []string {
	return []string{"not implemented yet"}
}

func downloadPage(url string) {
	fmt.Println("Pretending to download:", url)
}
