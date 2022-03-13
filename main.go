package main

import (
	"fmt"
	"sync"
)

// Fetcher is a type that can fetch URLs and
// return their body as a string and a slice of URLs found on that page
type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on the page
	Fetch(url string) (body string, urls []string, err error)
}

type Result struct {
	body  string
	urls  []string
	error error
}

// Get returns a Result from the cache for a given key
func (c *Cache) Get(key string) (Result, bool) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	item, ok := c.items[key]
	return item, ok
}

// Set adds a Result to the cache for a given key
func (c *Cache) Set(key string, value Result) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.items[key] = value
}

type Cache struct {
	mtx   sync.Mutex
	items map[string]Result
}

// Crawl uses fetcher recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, cache *Cache) {
	// TODO: Fetch URLs in parallel
	// TODO: Don't fetch the same URL twice
	// this implementation doesn't do either
	if depth <= 0 {
		return
	}
	if _, ok := cache.Get(url); ok {
		// fmt.Printf("\t<<%s is cached\n", url)
		return
	}
	body, urls, err := fetcher.Fetch(url)
	cache.Set(url, Result{body, urls, err})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s, %q\n", url, body)
	for _, u := range urls {
		Crawl(u, depth-1, fetcher, cache)
	}
}

func main() {
	cache := &Cache{sync.Mutex{}, make(map[string]Result)}
	Crawl("https://golang.org/", 4, fetcher, cache)
}

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
