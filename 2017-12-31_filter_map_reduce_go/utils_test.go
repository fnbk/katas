package utils

import (
	"reflect"
	"strconv"
	"testing"
)

//
// Contains
//

func TestContains(t *testing.T) {
	slice := []int{1, 2, 3}
	if !Contains(slice, 1) {
		t.Errorf("Expected to contain 1")
	}
	if Contains(slice, 0) {
		t.Errorf("Expected to NOT contain 0")
	}
}

func TestContainsStruct(t *testing.T) {
	type A struct {
		B int
	}
	slice := []A{{1}, {2}, {3}}
	if !Contains(slice, A{1}) {
		t.Errorf("Expected to contain 1")
	}
}

func TestContainsPanicOnTypeMismatch(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	Contains([]struct{}{}, 1)
}

func TestContainsEmptySlice(t *testing.T) {
	slice := []int{}
	if Contains(slice, 0) {
		t.Errorf("Expected to NOT contain 0")
	}
}

//
// Remove
//

func TestRemove(t *testing.T) {
	slice := []int{1, 2, 3}
	result := Remove(slice, 2)
	expected := []int{1, 3}
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Remove() failed, expected:%v go:%v", expected, result)
	}
}

func TestRemoveEmptySlice(t *testing.T) {
	slice := []int{}
	result := Remove(slice, 2)
	expected := []int{}
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Remove() failed, expected:%v go:%v", expected, result)
	}
}

//
// Index
//

func TestIndexFound(t *testing.T) {
	slice := []int{1, 2, 3}
	result := Index(slice, 1)
	expected := int32(0)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Index() failed, expected:%v go:%v", expected, result)
	}
}

func TestIndexNotFound(t *testing.T) {
	slice := []int{1, 2, 3}
	result := Index(slice, 0)
	expected := int32(-1)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Index() failed, expected:%v go:%v", expected, result)
	}
}

func TestIndexFunction(t *testing.T) {
	type A struct {
		B int
	}
	findElement := func(val A) bool { return val.B == 3 }
	slice := []A{{1}, {2}, {3}}
	result := Index(slice, findElement)
	expected := int32(2)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Index() failed, expected:%v go:%v", expected, result)
	}
}

func TestIndexEmptySlice(t *testing.T) {
	slice := []int{}
	result := Index(slice, 1)
	expected := int32(-1)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Index() failed, expected:%v go:%v", expected, result)
	}
}

//
// Filter
//

func TestFilter(t *testing.T) {
	evenNumbers := func(val int, idx int32, slice []int) bool { return (val % 2) == 0 }
	slice := []int{1, 2, 3, 4}
	result := Filter(slice, evenNumbers).([]int)
	expected := []int{2, 4}

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Map() failed, expected:%v go:%v", expected, result)
	}
}

func TestFilterEmptySlice(t *testing.T) {
	evenNumbers := func(val int, idx int32, slice []int) bool { return (val % 2) == 0 }
	slice := []int{}
	result := Filter(slice, evenNumbers).([]int)
	expected := []int{}

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Map() failed, expected:%v go:%v", expected, result)
	}
}

//
// MapInPlace
//

func TestMapInPlace(t *testing.T) {
	type A struct {
		B int
	}
	double := func(val *A, idx int32, slice []*A) { *(slice[idx]) = A{val.B * 2} }
	slice := []*A{{1}, {2}, {3}}
	MapInPlace(slice, double)
	expected := []*A{{2}, {4}, {6}}

	if !reflect.DeepEqual(expected, slice) {
		t.Errorf("Map() failed, expected:%v go:%v", expected, slice)
	}
}

func TestMapInPlaceEmptySlice(t *testing.T) {
	type A struct {
		B int
	}
	double := func(val *A, idx int32, slice []*A) { *(slice[idx]) = A{val.B * 2} }
	slice := []*A{}
	MapInPlace(slice, double)
	expected := []*A{}

	if !reflect.DeepEqual(expected, slice) {
		t.Errorf("Map() failed, expected:%v go:%v", expected, slice)
	}
}

//
// Map
//

func TestMap(t *testing.T) {
	toString := func(val int, idx int32, slice []int) string {
		return strconv.Itoa(val)
	}
	slice := []int{1, 2, 3}
	result := Map(slice, toString).([]string)
	expected := []string{"1", "2", "3"}
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Map() failed, expected:%v go:%v", expected, result)
	}
}

func TestMapEmptySlice(t *testing.T) {
	toString := func(val int, idx int32, slice []int) string {
		return strconv.Itoa(val)
	}
	slice := []int{}
	result := Map(slice, toString).([]string)
	expected := []string{}
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Map() failed, expected:%v go:%v", expected, result)
	}
}

//
// Reduce
//

func TestReduceInt32(t *testing.T) {
	add := func(val int32, idx int32, slice []int32, acc int32) int32 { return acc + val }
	arr := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := Reduce(arr, add, int32(0)).(int32)
	expected := int32(55)
	if expected != result {
		t.Errorf("Reduce() failed expected:%d, got:%d", expected, result)
	}
}

func TestReduceString(t *testing.T) {
	concatenate := func(val string, idx int32, slice []string, acc string) string { return acc + val }
	arr := []string{"1 ", "2 ", "3 ", "4 "}
	result := Reduce(arr, concatenate, "").(string)
	expected := "1 2 3 4 "
	if expected != result {
		t.Errorf("Reduce() failed expected:%d, got:%d", expected, result)
	}
}

type MyAcc struct {
	A int
}

type MyStruct struct {
	A int
}

func TestReduceStruct(t *testing.T) {
	myReduce := func(val MyStruct, idx int32, slice []MyStruct, acc MyAcc) MyAcc {
		acc.A = acc.A + val.A
		return acc
	}
	arr := []MyStruct{{2}, {40}, {0}}
	result := Reduce(arr, myReduce, MyAcc{}).(MyAcc)
	expected := MyAcc{42}
	if expected != result {
		t.Errorf("Reduce() failed expected:%d, got:%d", expected, result)
	}
}

func TestReduceEmptySlice(t *testing.T) {
	add := func(val int, idx int32, slice []int, acc int) int { return acc + val }
	slice := []int{}
	result := Reduce(slice, add, 0).(int)
	expected := 0
	if expected != result {
		t.Errorf("Reduce() failed expected:%d, got:%d", expected, result)
	}
}

func TestReduceNilSlice(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	add := func(val int, idx int32, slice []int, acc int) int { return acc + val }
	Reduce(nil, add, 0)
}
