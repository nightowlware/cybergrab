package main

import (
	"fmt"
)

func main() {
	fmt.Println("CyberGrab, v0.1")

	var engine Scheduler
	var downloader Downloader

	numdownloads should be in the downloader
	downloader = NewSimpleDownloader("downloads")
	engine = NewSimpleScheduler(10, 1000, downloader)

	engine.run("http://www.cnn.com")

	//var input string
	//fmt.Scanln(&input)
	//panic("END")
}
