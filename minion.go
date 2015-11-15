package cybergrab

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
	"time"
)

type pageMinion struct {
	scheduler scheduler
}

func invalidUrl(url string) bool {
	return url == ""
}

// Instructs this minion to process the given url.i
//
// This minion can do any combination of:         }
// 1. Push a new url into the download queue
// 2. Push a new url into the link dispenser queuei, for other minions to run()
// 3. Do nothing
//                                                }
// Returns error or nil for success
func (this pageMinion) run(url string) error {
	fmt.Println("pageMinion: Scrubbing page: " + url)

	if invalidUrl(url) {
		return fmt.Errorf("URL <%s> is invalid, ignoring.", url)
	}

	// set a timeout for http.Get()
	timeout := time.Duration(WORKER_TIMEOUT_SECONDS * time.Second)
	client := http.Client {
		Timeout: timeout,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	z := html.NewTokenizer(resp.Body)

	// iterate through all hrefs and decide whether to
	// download the page, push the page into the queue for
	// other PageScrubbers to process, or both.
	for tt := z.Next(); tt != html.ErrorToken; tt = z.Next() {
		switch {
		case tt == html.StartTagToken:
			t := z.Token()

			// skip tags that are neither links to follow or downloadable content
			if !(t.Data == "a" || t.Data == "img" || t.Data == "video" || t.Data == "audio") {
				continue
			}

			for _, attr := range t.Attr {
				href_link := attr.Val

				// deal with relative links
				if !strings.HasPrefix(href_link, "http://") &&
					!strings.HasPrefix(href_link, "https://") {
					href_link = url + href_link
				}

				// is this a link tag?
				if attr.Key == "href" {
					if this.scheduler.getCrawlPolicy().ShouldCrawl(href_link) {
						this.scheduler.getLinkDispenser().pushUrl(href_link)
					}
				}

				if attr.Key == "src" {
					if this.scheduler.getCrawlPolicy().ShouldDownload(href_link) {
						this.scheduler.getDownloader().addDownload(href_link)
					}
				}
			}
		}
	}

	return nil
}
