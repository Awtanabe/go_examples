### 使い方

```
// 正常系
➜  test_cobra go run main.go     
Hello from the root command!

// 引数あり
go run main.go greet John
Hello, John!

// 失敗
➜  test_cobra go run main.go fail
Error: this command failed intentionally
Usage:
  app fail [flags]

Flags:
  -h, --help   help for fail

Error: this command failed intentionally
exit status 1
```