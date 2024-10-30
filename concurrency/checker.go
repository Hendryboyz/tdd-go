package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)
	for _, url := range urls {
		go func(u string) {
			// write to map directly could cause `fatal error: concurrent map writes`
			// results[u] = wc(u)
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for range len(urls) {
		r := <-resultChannel
		results[r.string] = r.bool
	}
	return results
}
