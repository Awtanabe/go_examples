package select_test

import (
	"fmt"
	"net/http"
	"time"
)

// func Racer(a, b string) (winner string) {
// 	aDucration := mearsureResponseTime(a)
// 	bDucration := mearsureResponseTime(b)

// 	if aDucration < bDucration {
// 		return a
// 	}
// 	return b
// }

func Racer(a, b string , timeout time.Duration) (winner string, error error) {

	select {
	case <-ping(a):
		  return a, nil
	case <-ping(b):
			return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
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

func mearsureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}