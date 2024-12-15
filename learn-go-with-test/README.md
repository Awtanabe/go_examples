### 参考

https://andmorefine.gitbook.io/learn-go-with-tests/go-fundamentals/install-go

## この記事で意識されていること

- リファクタリング
- テスト


## 気付き

- package単位でグループ化される
  - 本体のファイルとテストのファイルで共有される
    - ex: package iteration
      - iteration_test.go, iteration.goでは関数が共される
- *testing.Bはベンチマーク
  - 速度計測？
  - loopの箇所に記述
  - 実行
    - go test -bench="." 

```
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
			Repeat("a")
	}
}

BenchmarkRepeat-10      14355256                83.69 ns/op

// 14355256回数実行 b.Nが  14355256までテストした 10^8 だから　2秒以下なら余裕
// ナノ秒単位
```

### hello worldセクション

- 開発の規律(サイクル)について
- テストを書く
  - t.Runも利用した
- if const switchなどの基本構文
- 変数、定数の宣言

### loopセクション

- for文のテスト
- ベンチマークの理解
- *testing.Bはベンチマーク
  - 速度計測？
  - loopの箇所に記述
  - 実行
    - go test -bench="." 

```
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
			Repeat("a")
	}
}

BenchmarkRepeat-10      14355256                83.69 ns/op

// 14355256回数実行 b.Nが  14355256までテストした 10^8 だから　2秒以下なら余裕
// ナノ秒単位
```

### Structs_methods_interfaces

- テーブル駆動テスト
  - TestAreaで

### pointer

- レシーバポインターによる挙動の確認
- エラー
  - errors.New
- エラーのテストが足りていないかチェック
  - https://qiita.com/Sekky0905/items/a47e83e7f97d5311e72b