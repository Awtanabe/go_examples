package main

import (
	"bytes"
	"fmt"
)


func Greet(writer *bytes.Buffer, name string) {
	// fmt.Printf("Hello, %s", name)
	fmt.Fprintf(writer, "Hello, %s", name)
}


// func Fprintf(w io.Writer,に 違和感があったけど、interfaceだからその後に Writeで書き込み処理などがあるから
// Writerでもおかしくないか
// Writer => 実行時 *bytes.Buffer(bufを保持したもの。編集されれるデータ)だけど、書き込むメソッドを保持してるから書き込む側ではある 
// ```
// func Greet(writer *bytes.Buffer, name string) {
//     fmt.Fprintf(writer, "Hello, %s", name)
// }
// ```
// type Writer interface {
// 	Write(p []byte) (n int, err error)
// }
