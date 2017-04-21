package main

import (
	"fmt"
	"sync"
)

// Fetcher interface method
type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Cache of URL's we've already visited.
type Cache struct {
	visited map[string]bool
	sync.Mutex
}

// Results for a single crawled URL.
type Response struct {
	url  string
	body string
}

// Sets the given url in the cache.
// Returns false if it already existed.
func (cache *Cache) isVisitedOrSet(url string) bool {
	// Unlock for an other can access the map cache.visited.
	cache.Lock()
	_, ok := cache.visited[url]
	cache.visited[url] = true
	// Lock so only one goroutine at a time can access the map cache.visited.
	cache.Unlock()
	return !ok
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, ch chan Response, cache *Cache, wg *sync.WaitGroup) {
	defer wg.Done()
	// Don't fetch the same URL twice
	if depth <= 0 || !cache.isVisitedOrSet(url) {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Send success result to channel
	ch <- Response{url, body}
	for _, u := range urls {
		wg.Add(1)
		// Fetch URLs in parallel
		go Crawl(u, depth-1, fetcher, ch, cache, wg)
	}
	return
}

func readResult(ch chan Response) {
	for {
		select {
		case res := <-ch:
			fmt.Println("==============================")
			fmt.Printf("found: %s %q\n", res.url, res.body)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	var ch = make(chan Response)
	cache := &Cache{visited: make(map[string]bool)}

	wg.Add(1)
	go Crawl("http://golang.org/", 4, fetcher, ch, cache, &wg)
	go readResult(ch)
	wg.Wait()
	fmt.Println(cache)

	close(ch)
}

/* ----------- BEGIN IMPLEMENT FAKE DATA ----------- */

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("Not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		body: "The Go Programming Language",
		urls: []string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		body: "Packages",
		urls: []string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		body: "Package fmt",
		urls: []string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		body: "Package os",
		urls: []string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}

/* ----------- END IMPLEMENT FAKE DATA ----------- */
