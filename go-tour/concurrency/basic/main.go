package main

import (
	"fmt"
	"sync"
)

// ダメなケース
// func Say() {
// 	fmt.Println("hello")
// }
// メイン関数がSay関数の実行が完了する前に、main関数が完了してしまう
// func main() {
// 	for i := 0; i < 10; i++ {
// 		go Say()
// 	}
// }

// okなケース
//
func Say(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("hello")
}


func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		// ポインタを渡すのは、Addで内部状況を変えているからかな？
		go Say(&wg)
	}
	// 同期処理が goroutineを待つ
	wg.Wait()
}