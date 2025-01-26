### 進捗

- 定数

https://andmorefine.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world#ding-shu

### Goをインストールする

リンター

brew install golangci/tap/golangci-lint

### Hello world

- テストの手順
  - xx_test.goファイルの作成
  - Testで始まる関数を作成
  - テストライブラリを利用して got とwant を比較して t.Errorfで出力
    - ここの結果でテストが判定してくれる

- テストフロー
  - テストをかく
  - エラーになる
  - 実態をかく
  - テストを通す
  - リファクタリングする

- 定数
  - const 定数名 = "文字列"
  - const englishHelloPrefix = "Hello, "

### 整数

Add関数を作成する

### 反復、繰り返し

ベンチマーク

```
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
			Repeat("a")
	}
}


go test -bench=.
```


### 配列とスライス

- 配列

[5]intを期待する関数に [4]intを渡せない

```
# array [array.test]
./array_test.go:18:14: cannot use numbers (variable of type []int) as [5]int value in argument to Sum
FAIL    array [build failed]
```

- スライス

任意のサイズを持つ

- カバレッジ

```
go test -cover
```

- make

notionの実装メモ まとめの、
make メモリ確保 slice に結構まとめられた


### 構造体 interface

- interface Shape

type Shape interface {
  Area() float64
}

- テーブルテストの記載

### ポインター エラー

Wallet構造体 預金、引き出しの学習

- 預金(Deposit)でポインタにしないと変更できないことを学習

- フィールドは大文字だと外部呼び出し、小文字だと内部呼び出し

- castの学習

want := Bitcoin(10)

- error

err.Error()で、errorr.New("エラー")のエラーの内容を表示できる

- 未チェックのエラーを検出できるライブラリ

### map

- mapをレシーバーにすることが可能

```
type Dictionary map[string]string

Dictionary{"test": "this is a test"}

func (d Dictionary) Search(key string) string {
	return d[key]
}
```

- mapのキーがあるかの判定

```
	val, ok := d[key]
	if !ok {
		return "", errors.New("not found")
	}

```

## concurecy


- go func() {} ()でループするときに引数

```
for _, url := range urls {
  go func(u string) {
     map[u] = wc(u)
  } (url)
}
```

- go test -raceで競合状態を見れる


### sync 同期

安全？
=> wait を使うのはメイン go routineが先に終わってサブルーチン

- メイン go routineがサブルーチンより先に終わる

```
package main

import (
	"fmt"
)

func main() {
	fmt.Println("main")
	go hoge() //goroutine

}

func hoge() {
	fmt.Println("hoge")
}
```


- サブルーチンが終わるまでメインルーチンを待つ

```
// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)

	fmt.Println("main")
	go hoge(&wg) //goroutine
	wg.Wait()
}

func hoge(wg *sync.WaitGroup) {
	fmt.Println("hoge")
	wg.Done()
}

```