package cybergrab

import (
	"errors"
	"runtime/debug"
)

const (
	MAX_DOWNLOADS = 30000
	MAX_WORKERS   = 30000
	WORKER_TIMEOUT_SECONDS = 5
)

///////////////////////////////
/// Public API
///////////////////////////////
var InvalidArgs error = errors.New("numWorkers/numDownloads too high")

func NewSpider(policy CrawlPolicy, numDownloads uint, numWorkers uint) (Spider, error) {
	if numDownloads > MAX_DOWNLOADS || numWorkers > MAX_WORKERS {
		return nil, InvalidArgs
	}

	d := newSimpleDownloader("downloads", numDownloads)
	s := newSimpleScheduler(policy, numWorkers, d)

	// set the max number of threads
	debug.SetMaxThreads(MAX_WORKERS + MAX_DOWNLOADS + 1)

	return &engine{s, d}, nil
}

///////////////////////////////

type engine struct {
	scheduler  scheduler
	downloader downloader
}

func (e *engine) Crawl(seedUrl string) error {
	e.scheduler.run(seedUrl)

	//TODO: proper error handling from the scheduler
	return nil
}
