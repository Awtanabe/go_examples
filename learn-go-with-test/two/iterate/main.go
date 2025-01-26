package iteration

import "fmt"

func Repeat(init string) (string) {
	base := ""
	for i := 0; i <= 4; i++ {
		base += init
	}
	return base
}

func main() {
	fmt.Println(Repeat("a"))
}