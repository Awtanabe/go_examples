### 参考

https://zenn.dev/hsaki/books/golang-context/viewer/intro


### contextの概要

- 役割は3つ
  - 処理の締め切りを伝播
  - キャンセル信号の伝播
  - リクエストスコープ値の伝播

- context の意義
  - 複数のゴルーチンを跨ぐ例

### 処理が複数個のゴルーチンを跨ぐ例

下記のハンドラはmainのゴルーチンとは別のゴルーチンを立てて処理を行なっている
※意識しないでも、複数のゴルーチンを利用している
※異なるゴールーチン間での情報共有は、ロックを使ってメモリを共有するよりも、チャネルを使った伝達を使うべし

=> main1
  => sub1
  => sub2
=> 情報伝達が難しい

```
func main() {
	// ハンドラ関数の定義
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #1!\n")
	}
	h2 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #2!\n")
	}

	http.HandleFunc("/", h1) // /にきたリクエストはハンドラh1で受け付ける
	http.HandleFunc("/endpoint", h2) // /endpointにきたリクエストはハンドラh2で受け付ける

	// サーバー起動
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

### 途中で不要になるケース (context 利用しないdone)

generatorで無限にループさせている
=> <- numでgeneratoerパターンで送信をしつ付けていて
=> doneで受信が来たらクローズしてbreak

※気になったところ、generatorでデータを生成し続けるで余ると思ったけど、
ブロッキング機能があるので余らない

```
var wg sync.WaitGroup

// キャンセルされるまでnumをひたすら送信し続けるチャネルを生成
func generator(done chan struct{}, num int) <-chan int {
	out := make(chan int)
	go func() {
		defer wg.Done()

	LOOP:
		for {
			select {
			case <-done: // doneチャネルがcloseされたらbreakが実行される
				break LOOP
			case out <- num: // キャンセルされてなければnumを送信
			}
		}

		close(out)
		fmt.Println("generator closed")
	}()
	return out
}

func main() {
	done := make(chan struct{})
	gen := generator(done, 1)

	wg.Add(1)

	for i := 0; i < 5; i++ {
		fmt.Println(<-gen)
	}
	close(done) // 5回genを使ったら、doneチャネルをcloseしてキャンセルを実行

	wg.Wait()
}
```

### context

contextを利用するケース
・cancelを用いたケース
・ctx.Done()利用

```
var wg sync.WaitGroup

func generator(done chan struct{}, num int) <-chan int {
func generator(ctx context.Context, num int) <-chan int {
	out := make(chan int)
	go func() {
		defer wg.Done()

	LOOP:
		for {
			select {
			case <-done:
			case <-ctx.Done():
				break LOOP
			case out <- num:
			}
		}

		close(out)
		fmt.Println("generator closed")
	}()
	return out
}

func main() {
	done := make(chan struct{})
	gen := generator(done, 1)
	ctx, cancel := context.WithCancel(context.Background())
	gen := generator(ctx, 1)

	wg.Add(1)

	for i := 0; i < 5; i++ {
		fmt.Println(<-gen)
	}
	close(done)
	cancel()

	wg.Wait()
}

```

## 直列、並列のセクション

キャンセルされると、直列、並列だろうが連動して終了する

### 直列なゴルーチン

直列ってのは、ネストされて順々に実行
・ctx1はキャンセルされない

```

func main() {
	ctx0 := context.Background()

	ctx1, _ := context.WithCancel(ctx0)
	// G1
	go func(ctx1 context.Context) {
		ctx2, cancel2 := context.WithCancel(ctx1)

		// G2-1
		go func(ctx2 context.Context) {
			// G2-2
			go func(ctx2 context.Context) {
				select {
				case <-ctx2.Done():
					fmt.Println("G2-2 canceled")
				}
			}(ctx2)

			select {
			case <-ctx2.Done():
				fmt.Println("G2-1 canceled")
			}
		}(ctx2)

		cancel2()

		select {
		case <-ctx1.Done():
			fmt.Println("G1 canceled")
		}

	}(ctx1)

	time.Sleep(time.Second)
}

```

### 並列なゴルーチン


```
func main() {
	ctx0 := context.Background()

	ctx1, cancel1 := context.WithCancel(ctx0)
	// G1-1
	go func(ctx1 context.Context) {
		select {
		case <-ctx1.Done():
			fmt.Println("G1-1 canceled")
		}
	}(ctx1)

	// G1-2
	go func(ctx1 context.Context) {
		select {
		case <-ctx1.Done():
			fmt.Println("G1-2 canceled")
		}
	}(ctx1)

	cancel1()

	time.Sleep(time.Second)
}

```

⭐️直列、並列のセクションは、「ゴールーチンの生死を制御するcontextが同じであるので、キャンセルタイミングも当然連動すること」を伝えたいみたい


### 兄弟関係にあるcontextの場合

並列と変わらない気がする

・context.Background()の親のcontext(何もないcontext)から、withContextを作成している


### 親子関係にあるcontextの場合

親のコンテキストがキャンセルされるろ子のコンテキストも終わる
=> 子供のキャンセルを明示的にしていないけど、子のコンテキストもキャンセルされる


## Deadlineメソッド

### チャネルをを使ったケース

time.Afterで任意のタイミングでチェネル操作


### コンテキストを利用したケース

WithDeadlineで可能

```
// 初期化contextと時間
context.WithDeadline(context.Background(), time.Now().Add(time.Second))
```

### Errメソッド

- cancelかdeadlineかで区別できない
- contextの2種類のエラー変数
  - var Canceled = errors.New("context canceled")
  - var DeadlineExceeded error = deadlineExceededError{}
- Errメソッドの特徴
  - contextがキャンセルされない場合nil
  - 明示的にキャンセル => calceledエラー
  - タイムアウト => DeadlineExceeded

```
select {
case <-ctx.Done():
    if err := ctx.Err(); errors.Is(err, context.Canceled) {
        // キャンセルされていた場合
        fmt.Println("canceled")
    } else if errors.Is(err, context.DeadlineExceeded) {
        // タイムアウトだった場合
        fmt.Println("deadline")
    }
}
```

### Cause

キャンセルがmain由来からタスク由来かわからない


- before: context.WithCancel(context.Background())
- after: ctx, cancel := context.WithCancelCause(context.Background())
  - cancel(errors.New("canceled by main func"))

```
func subTask(ctx context.Context) {
	defer wg.Done()

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
		fmt.Println(context.Cause(ctx)) // ここで理由を表示できる
	case <-doSomething():
		fmt.Println("subtask done")
	}
}
```

## クリーンナップ処理

- こんな感じでできる

```
  ctx, cancel := context.WithCancel(context.Background())
	stop := context.AfterFunc(ctx, func() {
		fmt.Println("ctx cleanup done")
	})
	defer stop()
```

## Valueメソッド

- 渡す

```
  // キャンセルもある
  ctx, cancel := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, "userID", 2)
	ctx = context.WithValue(ctx, "authToken", "xxxxxxxx")
	ctx = context.WithValue(ctx, "traceID", 3)
	gen := generator(ctx, 1)
```

- 使う

```
		userID, authToken, traceID := ctx.Value("userID").(int), ctx.Value("authToken").(string), ctx.Value("traceID").(int)

```


### 実用例

https://zenn.dev/hsaki/books/golang-context/viewer/usage

- 認証で必要なデータをcontextに入れる
- sessionIDもcontextに入れる
- DBリクエストでタイムアウトの設定