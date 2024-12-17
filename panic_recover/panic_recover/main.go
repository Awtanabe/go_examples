package main

import (
	"errors"
	"fmt"
)


func main() {
	err := bar()

	if err != nil {
		fmt.Printf("%#v\n", err)
		return
	}
	fmt.Println("Normal End")
}

func bar() (err error) {
	defer func() {
		if rec := recover(); rec != nil {
			fmt.Printf("Type: %T, Value: %#v\n", rec, rec)
			err = fmt.Errorf("Recoverd from %w", rec)
		}
	}()

	foo()
	return
}

func foo() {
	panic(errors.New("Panic!"))
}