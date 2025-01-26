package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

// type Buffer struct {
// 	buf      []byte // contents are the bytes buf[off : len(buf)]
// 	off      int    // read at &buf[off], write at &buf[len(buf)]
// 	lastRead readOp // last read operation, so that Unread* can work correctly.
// }
// はWrite関数を持っているので

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct {}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}


// mock化
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i -- {
		sleeper.Sleep()
		fmt.Fprintf(out, strconv.Itoa(i))
	}
	sleeper.Sleep()
	fmt.Fprintf(out, finalWord)
}




func main() {
	s := &DefaultSleeper{}
	Countdown(os.Stdout, s)
}