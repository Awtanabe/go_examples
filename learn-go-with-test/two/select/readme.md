

### channel

https://hori-ryota.com/blog/golang-channel-pattern/

## 参考
https://tech.yappli.io/entry/goroutine-select

サンプルコード

close(chan) をすると、 <-c で読み込まれる

```
package main


import (
  "fmt"
  "time"
)

func main() {
    start := time.Now()
    c := make(chan int)
    go func() {
        time.Sleep(2 * time.Second)
        close(c) //1
    }()

    fmt.Println("Blocking on read...")
    select {
    case <-c: //2
        fmt.Printf("Unblocked %v later.\n", time.Since(start))
    }
}
```