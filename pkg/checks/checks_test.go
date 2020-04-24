package checks

import (
	"reflect"
	"testing"
)

func TestPrepareNumbers(t *testing.T) {
	{
		var val8 int8 = 1
		var val16 int16 = 1
		var val32 int32 = 1
		var val int = 1
		var res int64 = 1

		resType := reflect.ValueOf(res)
		val8Type := reflect.ValueOf(val8)
		// convert valType to resType
		if resType.Type() != prepareType(val8Type).Type() {
			t.Errorf("(val8) res type != prepared type; %v != %v", resType.Type(), prepareType(val8Type).Type())
		}
		val16Type := reflect.ValueOf(val16)
		if resType.Type() != prepareType(val16Type).Type() {
			t.Errorf("(val16) res type != prepared type; %v != %v", resType.Type(), prepareType(val16Type).Type())
		}
		val32Type := reflect.ValueOf(val32)
		if resType.Type() != prepareType(val32Type).Type() {
			t.Errorf("(val32) res type != prepared type; %v != %v", resType.Type(), prepareType(val32Type).Type())
		}
		valType := reflect.ValueOf(val)
		if resType.Type() != prepareType(valType).Type() {
			t.Errorf("(val) res type != prepared type; %v != %v", resType.Type(), prepareType(valType).Type())
		}
	}
	{
		var val8 uint8 = 1
		var val16 uint16 = 1
		var val32 uint32 = 1
		var val uint = 1
		var res uint64 = 1

		resType := reflect.ValueOf(res)
		val8Type := reflect.ValueOf(val8)
		// convert valType to resType
		if resType.Type() != prepareType(val8Type).Type() {
			t.Errorf("(val8) res type != prepared type; %v != %v", resType.Type(), prepareType(val8Type).Type())
		}
		val16Type := reflect.ValueOf(val16)
		if resType.Type() != prepareType(val16Type).Type() {
			t.Errorf("(val16) res type != prepared type; %v != %v", resType.Type(), prepareType(val16Type).Type())
		}
		val32Type := reflect.ValueOf(val32)
		if resType.Type() != prepareType(val32Type).Type() {
			t.Errorf("(val32) res type != prepared type; %v != %v", resType.Type(), prepareType(val32Type).Type())
		}
		valType := reflect.ValueOf(val)
		if resType.Type() != prepareType(valType).Type() {
			t.Errorf("(val) res type != prepared type; %v != %v", resType.Type(), prepareType(valType).Type())
		}
	}
	{
		var val32 float32 = 1
		var res float64 = 1

		resType := reflect.ValueOf(res)
		val32Type := reflect.ValueOf(val32)
		// convert valType to resType
		if resType.Type() != prepareType(val32Type).Type() {
			t.Errorf("(val8) res type != prepared type; %v != %v", resType.Type(), prepareType(val32Type).Type())
		}
	}
}

func TestCheckTypes(t *testing.T) {
	{
		var val1 int = -1
		var val2 uint = 1

		val1Type := reflect.ValueOf(val1)
		val2Type := reflect.ValueOf(val2)

		if checkType(val1Type, val2Type) {
			t.Errorf("(int, uint) val1 == val2; %v == %v", val1Type, val2Type)
		}

		var val3 int = 1
		val3Type := reflect.ValueOf(val3)
		if !checkType(val1Type, val3Type) {
			t.Errorf("(int, int) val1 != val2; %v != %v", val1Type, val3Type)
		}
	}
	{
		var val1 *int = nil
		var val2 *int = nil

		val1Type := reflect.ValueOf(val1)
		val2Type := reflect.ValueOf(val2)

		if !checkType(val1Type, val2Type) {
			t.Errorf("(nil, nil) val1 != val2; %v != %v", val1Type, val2Type)
		}
	}
}

func TestCompare(t *testing.T) {
	{
		var val1 int64 = 1008234
		var val2 int64 = 1008234

		val1Type := reflect.ValueOf(val1)
		val2Type := reflect.ValueOf(val2)

		if !compare(val1Type, val2Type) {
			t.Errorf("(int64, int64) val1 != val2; %v != %v", val1, val2)
		}
	}
	{
		var val1 float64 = 1.63454
		var val2 float64 = 1.63454

		val1Type := reflect.ValueOf(val1)
		val2Type := reflect.ValueOf(val2)

		if !compare(val1Type, val2Type) {
			t.Errorf("(float64, float64) val1 != val2; %v != %v", val1, val2)
		}
	}
	{
		var val1 uint64 = 7777
		var val2 uint64 = 88888

		val1Type := reflect.ValueOf(val1)
		val2Type := reflect.ValueOf(val2)

		if compare(val1Type, val2Type) {
			t.Errorf("(uint64, uint64) val1 == val2; %v == %v", val1, val2)
		}
	}
	{
		var val1 bool = true
		var val2 bool = false

		val1Type := reflect.ValueOf(val1)
		val2Type := reflect.ValueOf(val2)

		if compare(val1Type, val2Type) {
			t.Errorf("(bool, bool) val1 == val2; %v == %v", val1, val2)
		}
	}
	{
		var val1 string = "aaabbbcccddd"
		var val2 string = "aaabbbcccddd"

		val1Type := reflect.ValueOf(val1)
		val2Type := reflect.ValueOf(val2)

		if !compare(val1Type, val2Type) {
			t.Errorf("(string, string) val1 != val2; %v != %v", val1, val2)
		}
	}
	{
		var val1 *int32 = new(int32)
		var val2 *int32 = new(int32)
		*val1 = 12345
		val2 = val1

		val1Type := reflect.ValueOf(val1)
		val2Type := reflect.ValueOf(val2)

		if !compare(val1Type, val2Type) {
			t.Errorf("(int32, int32) val1 != val2; %v != %v", val1, val2)
		}
	}
	{
		val1 := []string{"a", "bb", "ccc", "dddd"}
		val2 := []string{"aaaa", "bbb", "cc", "d"}

		val1Type := reflect.ValueOf(val1)
		val2Type := reflect.ValueOf(val2)

		if compare(val1Type, val2Type) {
			t.Errorf("([]string, []string) val1 == val2; %v == %v", val1, val2)
		}
	}
	{
		val1 := make([]string, 4)
		val2 := make([]string, 4)

		val1Type := reflect.ValueOf(val1)
		val2Type := reflect.ValueOf(val2)

		if !compare(val1Type, val2Type) {
			t.Errorf("([]string, []string) val1 != val2; %v != %v", val1, val2)
		}
	}
}

func TestDeepComparison(t *testing.T) {
	if !deepComparion(1, 1) {
		t.Errorf("(int, int) val1 != val2; %v != %v", 1, 1)
	}
	if deepComparion(1, -54) {
		t.Errorf("(int, int) val1 == val2; %v == %v", 1, -54)
	}
	if !deepComparion("aace", "aace") {
		t.Errorf("(string, string) val1 != val2; %v != %v", "aace", "aace")
	}
	if deepComparion(true, false) {
		t.Errorf("(bool, bool) val1 == val2; %v == %v", true, false)
	}
	if deepComparion(1, true) {
		t.Errorf("(int, bool) val1 == val2; %v == %v", 1, true)
	}
	if !deepComparion(nil, nil) {
		t.Errorf("(nil, nil) val1 != val2; %v != %v", 1, 1)
	}
}
