package race

import (
	"fmt"
	"net/http"
	"time"
)

const tenSecondsTimeout = 10 * time.Second

func Racer(urlA, urlB string) (winner string, err error) {
	// `select` allows you to wait on multiple channels
	return ConfigurableRacer(urlA, urlB, tenSecondsTimeout)
}

func ConfigurableRacer(urlA string, urlB string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(urlA):
		return urlA, nil
	case <-ping(urlB):
		return urlB, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", urlA, urlB)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
