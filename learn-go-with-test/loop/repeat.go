package iteration

func Repeat(s string) string {
	var res string
	for i := 0; i < 5; i++ {
		res += s
	}
	return res
}