package utils

import "reflect"

//
// Array-Hilfsfunktionen
//

//
// Überprüft ob ein Slice ein bestimmtes Element enthält
// (inkl. reflect Deep.Equal) und liefert true bzw. false zurück
//
func Contains(slice, element interface{}) bool {
	in := reflect.ValueOf(slice)
	if in.Kind() != reflect.Slice {
		panic("Contains: not slice")
	}

	elem := reflect.ValueOf(element)
	if elem.Type().Kind() != in.Type().Elem().Kind() {
		panic("Contains: 'slice' type and 'element' type do not match")
	}

	for i := 0; i < in.Len(); i++ {
		if reflect.DeepEqual(in.Index(i).Interface(), element) {
			return true
		}
	}
	return false
}

//
// Entfernt ein bestimmtes Element aus einem Slice
//
func Remove(slice, element interface{}) interface{} {
	in := reflect.ValueOf(slice)
	if in.Kind() != reflect.Slice {
		panic("Remove: not slice")
	}

	elem := reflect.ValueOf(element)
	if elem.Type().Kind() != in.Type().Elem().Kind() {
		panic("Remove: 'slice' type and 'element' type do not match")
	}

	idx := Index(slice, element)

	if idx >= 0 {
		out := reflect.MakeSlice(reflect.SliceOf(elem.Type()), 0, in.Len()-1)
		for j := 0; j < in.Len(); j++ {
			if int(idx) != j {
				out = reflect.Append(out, in.Index(j))
			}
		}
		return out.Interface()
	}

	return slice
}

//
// Liefert den Index eines bestimmten Elements im Slice zurück
//
func Index(slice, element interface{}) int32 {
	in := reflect.ValueOf(slice)
	if in.Kind() != reflect.Slice {
		panic("Index: not slice")
	}

	var fn, elem reflect.Value
	if reflect.ValueOf(element).Type().Kind() == reflect.Func {
		fn = reflect.ValueOf(element)
	} else {
		elem = reflect.ValueOf(element)
	}

	if fn.IsValid() {
		if fn.Type().NumIn() != 1 {
			panic("Index: the input function needs to have one input parameter only")
		}
		if fn.Type().NumOut() != 1 {
			panic("Index: the input function needs to have one return argument only")
		}
		if fn.Type().In(0).Kind() != in.Type().Elem().Kind() {
			panic("Index: the input function's first argument type does not match the type of the 'slice' argument")
		}
		if fn.Type().Out(0).Kind() != reflect.Bool {
			panic("Index: the input function's return type does not equal 'Bool'")
		}
	}

	if elem.IsValid() {
		if elem.Type().Kind() != in.Type().Elem().Kind() {
			panic("Index: 'slice' type and 'element' type do not match")
		}
	}

	for i := 0; i < in.Len(); i++ {
		if fn.IsValid() {
			val := in.Index(i)
			ins := []reflect.Value{val}
			good := fn.Call(ins)[0]
			if good.Bool() {
				return int32(i)
			}
		} else {
			if reflect.DeepEqual(in.Index(i).Interface(), element) {
				return int32(i)
			}
		}
	}
	return -1
}

//
// Filter-Map-Reduce Funktionen (siehe Implementierungen für Beispiele)
//

// Filter:
// Filter ein übergebenes Objekt nach bestimmten Kriterien filtern und das (mengenmäßig reduzierte)
// Ergebnis zurück

// Map:
// Führt ein Mapping eines übergebenen Objekts auf ein anderes Objekt durch und gibt dieses zurück

// MapInPlace:
// Führt ein Mapping eines übergebenen Objekts durch, dabei wird das Objekt selbst geändert (in-place)

// Reduce:
// Aggregiert ein übergebenes Objekt anhand bestimmter Kriterien und gibt das Ergebnis zurück

type FilterFunc func(slice, function interface{}) interface{}
type MapInPlaceFunc func(slice, function interface{})
type MapFunc func(slice, function interface{}) interface{}
type ReduceFunc func(slice, function, accumulator interface{}) interface{}

//
// Struct, an dem die Filter-, Map- und Reduce-Funktionen aufgerufen werden können,
// die Verdrahtung mit den Funktions-Implementierungen geschieht in der Factory-Methode
//
type FilterMapReduce struct {
	filterFunc     FilterFunc
	mapFunc        MapFunc
	mapInPlaceFunc MapInPlaceFunc
	reduceFunc     ReduceFunc
}

//
// Factory
//
func New() FilterMapReduce {
	return FilterMapReduce{
		filterFunc:     filterImpl,
		mapFunc:        mapImpl,
		mapInPlaceFunc: mapInPlaceImpl,
		reduceFunc:     reduceImpl,
	}
}

//
// Standard-Verdrahtung der Funktionen (wieder-) herstellen
//
func Reset() {
	SetFilterFunc(filterImpl)
	SetMapFunc(mapImpl)
	SetMapInPlaceFunc(mapInPlaceImpl)
	SetReduceFunc(reduceImpl)
}

//
// Setzen/Überschreiben der Funktions-Verdrahtung
//
func SetFilterFunc(fn FilterFunc) {
	std.filterFunc = fn
}

func SetMapFunc(fn MapFunc) {
	std.mapFunc = fn
}

func SetMapInPlaceFunc(fn MapInPlaceFunc) {
	std.mapInPlaceFunc = fn
}

func SetReduceFunc(fn ReduceFunc) {
	std.reduceFunc = fn
}

// siehe golang log package: https://golang.org/src/log/log.go
var std = New()

func Filter(slice, function interface{}) interface{} {
	return std.filterFunc(slice, function)
}

func MapInPlace(slice, function interface{}) {
	std.mapInPlaceFunc(slice, function)
}

func Map(slice, function interface{}) interface{} {
	return std.mapFunc(slice, function)
}

func Reduce(slice, function, accumulator interface{}) interface{} {
	return std.reduceFunc(slice, function, accumulator)
}

// Beispiel:
//  evenNumbersFunc := func(val int, idx int32, slice []int) bool { return (val % 2) == 0 }
//  slice := []int{1, 2, 3, 4}
//  result := Filter(slice, evenNumbersFunc).([]int) // {2, 4}
func filterImpl(slice, function interface{}) interface{} {
	//
	// input
	//

	in := reflect.ValueOf(slice)
	if in.Kind() != reflect.Slice {
		panic("Filter: not slice")
	}

	//
	// function
	//

	fn := reflect.ValueOf(function)

	if fn.Kind() != reflect.Func {
		panic("Filter: not a function")
	}

	if fn.Type().NumIn() != 3 {
		panic("Filter: the number of the function's input parameters is not correct")
	}

	if fn.Type().NumOut() != 1 {
		panic("Filter: the number of the function's output parameters is not correct")
	}

	if fn.Type().In(2).Kind() != reflect.Slice {
		panic("Filter: the type of the 'slice' argument (3rd) needs to be a Slice")
	}

	if fn.Type().In(0).Kind() != fn.Type().In(2).Elem().Kind() {
		panic("Filter: the type of the 'val' argument (1st) needs to be of the same type as the 'slice' argument (3rd)")
	}

	if fn.Type().In(1).Kind() != reflect.Int32 {
		panic("Filter: the type of the 'idx' argument (2nd) needs to be int32")
	}

	if fn.Type().Out(0).Kind() != reflect.Bool {
		panic("Filter: the type of the return argument needs to be bool")
	}

	//
	// output
	//

	out := reflect.MakeSlice(reflect.SliceOf(fn.Type().In(0)), 0, in.Len())

	for i := 0; i < in.Len(); i++ {
		val := in.Index(i)
		idx := reflect.ValueOf(int32(i))
		ins := []reflect.Value{val, idx, in}
		good := fn.Call(ins)[0]
		if good.Bool() {
			out = reflect.Append(out, val)
		}
	}

	return out.Interface()
}

// Beispiel:
//  type A struct {
//  	B int
//  }
//  double := func(val *A, idx int32, slice []*A) { *(slice[idx]) = A{val.B * 2} }
//  slice := []*A{{1}, {2}, {3}}
//  MapInPlace(slice, double)
func mapInPlaceImpl(slice, function interface{}) {
	//
	// input
	//

	in := reflect.ValueOf(slice)
	if in.Kind() != reflect.Slice {
		panic("Filter: not slice")
	}

	//
	// function
	//

	fn := reflect.ValueOf(function)

	if fn.Kind() != reflect.Func {
		panic("MapInPlace: not a function")
	}

	if fn.Type().NumIn() != 3 {
		panic("MapInPlace: the number of the function's input parameters is not correct")
	}

	if fn.Type().NumOut() != 0 {
		panic("MapInPlace: the number of the function's output parameters is not correct")
	}

	if fn.Type().In(2).Kind() != reflect.Slice {
		panic("MapInPlace: the type of the 'slice' argument (3rd) needs to be a Slice")
	}

	if fn.Type().In(0).Kind() != fn.Type().In(2).Elem().Kind() {
		panic("MapInPlace: the type of the 'val' argument (1st) needs to be of the same type as the 'slice' argument (3rd)")
	}

	if fn.Type().In(1).Kind() != reflect.Int32 {
		panic("MapInPlace: the type of the 'idx' argument (2nd) needs to be int32")
	}

	//
	// in place
	//

	for i := 0; i < in.Len(); i++ {
		val := in.Index(i)
		idx := reflect.ValueOf(int32(i))
		ins := []reflect.Value{val, idx, in}
		fn.Call(ins)
	}
}

// Beispiel:
//  toStringFunc := func(val int, idx int32, slice []int) string { return strconv.Itoa(val) }
//  slice := []int{1, 2, 3}
//  result := Map(slice, toStringFunc).([]string) // {"1", "2", "3"}
func mapImpl(slice, function interface{}) interface{} {
	//
	// input
	//

	in := reflect.ValueOf(slice)
	if in.Kind() != reflect.Slice {
		panic("Filter: not slice")
	}

	//
	// function
	//

	fn := reflect.ValueOf(function)

	if fn.Kind() != reflect.Func {
		panic("Map: not a function")
	}

	if fn.Type().NumIn() != 3 {
		panic("Map: the number of the function's input parameters is not correct")
	}

	if fn.Type().NumOut() != 1 {
		panic("Map: the number of the function's output parameters is not correct")
	}

	if fn.Type().In(2).Kind() != reflect.Slice {
		panic("Map: the type of the 'slice' argument (3rd) needs to be a Slice")
	}

	if fn.Type().In(0).Kind() != fn.Type().In(2).Elem().Kind() {
		panic("Map: the type of the 'val' argument (1st) needs to be of the same type as the 'slice' argument (3rd)")
	}

	if fn.Type().In(1).Kind() != reflect.Int32 {
		panic("Map: the type of the 'idx' argument (2nd) needs to be int32")
	}

	//
	// output
	//

	out := reflect.MakeSlice(reflect.SliceOf(fn.Type().Out(0)), in.Len(), in.Len())

	for i := 0; i < in.Len(); i++ {
		val := in.Index(i)
		idx := reflect.ValueOf(int32(i))
		ins := []reflect.Value{val, idx, in}
		result := fn.Call(ins)[0]
		out.Index(i).Set(result)
	}

	return out.Interface()
}

// zektoren2 := Reduce(zektoren1, myReduce, SliceAccumulator{}).(SliceAccumulator).Slice
// func Reduce(slice interface{}, function interface{}, zero interface{}) interface{} { return nil }
// func reduceA(acc int32, val string, idx, slice []string) int32 { return 1 }

// * type-assert by the caller
// Beispiel:
//  addFunc := func(acc int, val int, idx int32, slice []int) int { return acc + val }
//	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//  result := Reduce(slice, addFunc, 0).(int)
func reduceImpl(slice, function, accumulator interface{}) interface{} {
	//
	// input
	//

	in := reflect.ValueOf(slice)
	if in.Type().Kind() != reflect.Slice {
		panic("Reduce: not slice")
	}

	//
	// function
	//

	fn := reflect.ValueOf(function)

	if fn.Kind() != reflect.Func {
		panic("Reduce: not a function")
	}

	if fn.Type().NumIn() != 4 {
		panic("Reduce: the number of the function's input parameters is not correct")
	}

	if fn.Type().NumOut() != 1 {
		panic("Reduce: the number of the function's output parameters is not correct")
	}

	if fn.Type().In(2).Kind() != reflect.Slice {
		panic("Reduce: the type of the 'slice' argument (3rd) needs to be a Slice")
	}

	if fn.Type().In(0).Kind() != fn.Type().In(2).Elem().Kind() {
		panic("Reduce: the type of the 'val' argument (1st) needs to be of the same type as the 'slice' argument (3rd)")
	}

	if fn.Type().In(1).Kind() != reflect.Int32 {
		panic("Reduce: the type of the 'idx' argument (2nd) needs to be int32")
	}

	if fn.Type().In(2).Elem().Kind() != in.Type().Elem().Kind() {
		panic("Reduce: the type of the function-'slice' argument (3rd) needs to be of the same type as the Reduce-'slice' argument (1st)")
	}

	if fn.Type().In(3).Kind() != fn.Type().Out(0).Kind() {
		panic("Reduce: the type of the 'acc' argument (4th) needs to be of the same type as the return argument")
	}

	//
	// accumulator
	//

	acc := reflect.ValueOf(accumulator)

	if acc.Type().Kind() != fn.Type().In(3).Kind() {
		panic("Reduce: the type of the 'accumulator' argument (4th) needs to be of the same type as the 'acc' argument (4th)")
	}

	//
	// output
	//

	for i := 0; i < in.Len(); i++ {
		val := in.Index(i)
		idx := reflect.ValueOf(int32(i))
		ins := []reflect.Value{val, idx, in, acc}
		acc = fn.Call(ins)[0]
	}

	return acc.Interface()
}
