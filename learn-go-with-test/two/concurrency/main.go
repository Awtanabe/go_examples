package concurrency

// これは関数でチェックするやつ
type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebSites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	resultChaneel:= make(chan result)
	for _, url := range urls {
		go func(u string) {
			resultChaneel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <- resultChaneel
		results[result.string] = result.bool
	}
	return results
}