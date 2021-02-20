package racer

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecodTimeout = 10 * time.Second

func measureResponseTime(url string) time.Duration {
	startTime := time.Now()
	http.Get(url)
	return time.Since(startTime)
}

func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, tenSecodTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {

	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("time out for waiting for %s and %s", a, b)
	}
}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}
