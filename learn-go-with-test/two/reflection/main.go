package reflection_test

import (
	"fmt"
	"reflect"
)

func walk(x interface{}, fn func(input string)) {
	t := reflect.TypeOf(x)
	value := reflect.ValueOf(x)
	switch t.Kind() {
	case reflect.String:
		fn(value.String())
	case reflect.Struct:
		fn(value.Field(0).String())
	}
}

func main() {

	walk("aaa", func(name string) { fmt.Println("xxx", name) })

	walk(struct{ Yahoo string }{"Yahoo"}, func(name string) { fmt.Println("xxx", name) })
	fmt.Println("Hello, 世界")
}
