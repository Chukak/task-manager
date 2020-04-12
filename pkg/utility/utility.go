package utility

import (
	"reflect"
)

// FuncVoidType implements function type without returning value.
type FuncVoidType func()

// BindedFunction implements type after calling 'Bind' function.
// Represent base functions as simple functions 'func ()'.
type BindedFunction struct {
	F FuncVoidType
}

// Bind function to 'BindedFunction' object.
func Bind(fn interface{}, params ...interface{}) BindedFunction {
	var arr []reflect.Value
	for _, v := range params {
		arr = append(arr, reflect.ValueOf(v))
	}

	return BindedFunction{F: func() {
		reflect.ValueOf(fn).Call(arr)
	}}
}
