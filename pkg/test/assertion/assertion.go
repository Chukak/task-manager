package test

// todo: improve this package

import (
	"testing"
)

type CurrentTestObject struct {
	o *testing.T
}

var __current = CurrentTestObject{}

func SetT(t *testing.T) {
	__current.o = t
}

func CheckEqual(v1 interface{}, v2 interface{}) {
	if v1 != v2 {
		__current.o.Errorf("CheckEqual failed: expected v1 == v2, but v1 != v2; v1 = %v, v2 = %v",
			v1, v2)
	}
}

func CheckNotEqual(v1 interface{}, v2 interface{}) {
	if v1 == v2 {
		__current.o.Errorf("CheckNotEqual failed: expected v1 != v2, but v1 == v2; v1 = %v, v2 = %v",
			v1, v2)
	}
}

func CheckTrue(b bool) {
	if !b {
		__current.o.Errorf("CheckTrue failed: expected true, but received %v", b)
	}
}

func CheckFalse(b bool) {
	if b {
		__current.o.Errorf("CheckFalse failed: expected false, but received %v", b)
	}
}
