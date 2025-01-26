package main

import (
	"fmt"
)

func main() {

	doWork()

}

func doWork() {
	for i := 1; i <= 10; i++ {
		fmt.Println("loop ...", i)
	}
}
