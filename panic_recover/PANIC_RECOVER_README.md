### 参考記事

https://zenn.dev/spiegel/books/error-handling-in-golang/viewer/panics

- ポイント
  - recoverでpanic()の引数にいれたものがそのまま出てくる
  - 文字列もだし、error.New()もだし