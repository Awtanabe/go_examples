package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

// WebsiteCheckerはチェックの処理をDIすることで依存関係を減らせる
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	// urlが何個になるかわからないからかな？
	results := make(map[string]bool)
	resultChannel := make(chan result)

	// 取得は並列処理
	for _, url := range urls {
		go func (u string) {
			// 書き込み
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	// 呼び出しはループ
	for i := 0; i < len(urls); i ++ {
		result := <- resultChannel
		results[result.string] = result.bool
	}

	return results
}