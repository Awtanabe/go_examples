package main

import "fmt"

func main() {
    foo()
}

// エラーになるケース
    func foo() {
        numbers := []int{0, 1, 2}
        fmt.Println(numbers[3])
    }


// 意図的に

// func foo() {
//     panic("panic")
// }