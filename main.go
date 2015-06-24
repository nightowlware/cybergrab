package main

import (
	"fmt"
)

func main() {
	fmt.Println("CyberGrab, v1.0")

	var engine Scheduler
	var downloader Downloader

	downloader = NewSimpleDownloader("downloads")
	engine = NewSimpleScheduler(10, downloader)

	engine.run("http://www.cnn.com")

	//var input string
	//fmt.Scanln(&input)
	//panic("END")
}
