package checks

// todo: improve this package
import (
	"fmt"
	"reflect"
	"testing"
)

type currentTestObject struct {
	o *testing.T
}

var currentT = currentTestObject{o: nil}

// SetT setting a *testing.T object as global for this module
func SetT(t *testing.T) {
	currentT.o = t
}

func checkType(val1 reflect.Value, val2 reflect.Value) bool {
	ok := val1.IsValid() && val2.IsValid()
	if ok {
		ok = val1.Type() == val2.Type()
	}
	return ok
}

func compare(val1 reflect.Value, val2 reflect.Value) bool {
	var ok bool = false
	switch val1.Kind() {
	case reflect.Uint64:
		ok = val1.Interface().(uint64) == val2.Interface().(uint64)
	case reflect.Int64:
		ok = val1.Interface().(int64) == val2.Interface().(int64)
	case reflect.Float64:
		ok = val1.Interface().(float64) == val2.Interface().(float64)
	case reflect.String:
		ok = val1.Interface().(string) == val2.Interface().(string)
	case reflect.Bool:
		ok = val1.Interface().(bool) == val2.Interface().(bool)
	case reflect.Array:
		ok = val1.Len() == val2.Len()
		for i := 0; ok && i < val1.Len(); i++ {
			ok = compare(val1.Index(i), val2.Index(i))
		}
	case reflect.Slice:
		ok = val1.IsNil() != val2.IsNil()
		if !ok {
			ok = val1.Len() == val2.Len()
			if ok {
				for i := 0; ok && i < val1.Len(); i++ {
					ok = compare(val1.Index(i), val2.Index(i))
				}
			}
		}
	case reflect.Ptr:
		ok = val1.Pointer() == val2.Pointer()
	// todo map, func if needeed
	default:
		ok = false
	}
	return ok
}

func prepareType(v reflect.Value) reflect.Value {
	switch v.Kind() {
	case reflect.Uint:
		return reflect.ValueOf(uint64(v.Interface().(uint)))
	case reflect.Uint8:
		return reflect.ValueOf(uint64(v.Interface().(uint8)))
	case reflect.Uint16:
		return reflect.ValueOf(uint64(v.Interface().(uint16)))
	case reflect.Uint32:
		return reflect.ValueOf(uint64(v.Interface().(uint32)))
	case reflect.Int:
		return reflect.ValueOf(int64(v.Interface().(int)))
	case reflect.Int8:
		return reflect.ValueOf(int64(v.Interface().(int8)))
	case reflect.Int16:
		return reflect.ValueOf(int64(v.Interface().(int16)))
	case reflect.Int32:
		return reflect.ValueOf(int64(v.Interface().(int32)))
	case reflect.Float32:
		return reflect.ValueOf(float64(v.Interface().(float32)))
	default:
	}
	return v
}

func deepComparion(v1 interface{}, v2 interface{}) bool {
	if v1 == nil || v2 == nil {
		return v1 == v2
	}

	val1 := prepareType(reflect.ValueOf(v1))
	val2 := prepareType(reflect.ValueOf(v2))
	ok := checkType(val1, val2)

	if ok {
		ok = compare(val1, val2)
	}
	return ok
}

// CheckEqual compare two values and print error if not equal.
func CheckEqual(v1 interface{}, v2 interface{}) {
	if !deepComparion(v1, v2) {
		if currentT.o != nil {
			currentT.o.Errorf("CheckEqual failed: expected v1 == v2, but v1 != v2; v1 = %v, v2 = %v",
				v1, v2)
		} else {
			fmt.Println(fmt.Errorf("The global testing object is nil! " +
				"Use SetT(*testing.T) function in your test function before all the tests."))
		}
	}
}

// CheckNotEqual compare two values and pritn error if equal.
func CheckNotEqual(v1 interface{}, v2 interface{}) {
	if deepComparion(v1, v2) {
		if currentT.o != nil {
			currentT.o.Errorf("CheckNotEqual failed: expected v1 != v2, but v1 == v2; v1 = %v, v2 = %v",
				v1, v2)
		} else {
			fmt.Println(fmt.Errorf("The global testing object is nil! " +
				"Use SetT(*testing.T) function in your test function before all the tests."))
		}
	}
}

// CheckTrue check value of argument and print error if false.
func CheckTrue(b bool) {
	if !b {
		if currentT.o != nil {
			currentT.o.Errorf("CheckTrue failed: expected true, but received %v", b)
		} else {
			fmt.Println(fmt.Errorf("The global testing object is nil! " +
				"Use SetT(*testing.T) function in your test function before all the tests."))
		}
	}
}

// CheckFalse check value of argument and print error if true.
func CheckFalse(b bool) {
	if b {
		if currentT.o != nil {
			currentT.o.Errorf("CheckFalse failed: expected false, but received %v", b)
		} else {
			fmt.Println(fmt.Errorf("The global testing object is nil! " +
				"Use SetT(*testing.T) function in your test function before all the tests."))
		}
	}
}
