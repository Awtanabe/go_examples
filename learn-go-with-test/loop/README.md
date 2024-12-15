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
