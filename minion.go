package cybergrab

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"reflect"
	"strings"
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
	} else {
		// parse the current page for links
		resp, err := http.Get(url)
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

				fmt.Println("t.Data is: ", t.Data)

				// if we find downloadable content
				if t.Data == "img" || t.Data == "link" {
					fmt.Println("Attr is: ", reflect.TypeOf(t.Attr[0]), t.Attr)

					for _, attr := range t.Attr {
						if attr.Key == "src" {
							href_link := attr.Val
							// deal with relative links
							if !strings.HasPrefix(href_link, "http://") {
								href_link = url + href_link
							}

							if this.scheduler.getCrawlPolicy().ShouldDownload(href_link) {
								this.scheduler.getDownloader().addDownload(href_link)
							}
						}
					}
				}

				// if we find an a link that can be followed
				if t.Data == "a" {
					// search for href
					for _, a := range t.Attr {
						if a.Key == "href" {
							href_link := a.Val
							// deal with relative links
							if !strings.HasPrefix(href_link, "http://") {
								href_link = url + href_link
							}

							if this.scheduler.getCrawlPolicy().ShouldCrawl(href_link) {
								this.scheduler.getLinkDispenser().pushUrl(href_link)
							}
						}
					}
				}
			}
		}
	}

	return nil
}
