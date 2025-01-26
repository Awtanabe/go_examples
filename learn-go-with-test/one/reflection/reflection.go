package reflection

import (
	"reflect"
)


func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)
	// xの値をコールバック関数に渡す
	// String()だと文字列以外の場合はエラーになるはず
	for i:= 0; i < val.NumField(); i ++ {
		field :=val.Field(i)
		fn(field.String())
	}
}