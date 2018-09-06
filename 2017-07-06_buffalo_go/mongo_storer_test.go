package main

// import (
// 	"reflect"
// 	"testing"
//
// 	"gopkg.in/mgo.v2/bson"
// )
//
// type Cowboy struct {
// 	UUID string `bson:"uuid"`
// 	// Age   int      `bson:"age"`
// 	// Names []string `bson:"names"`
// }
//
// func TestBuildUpdatenSelector(t *testing.T) {
// 	//
// 	// arrange
// 	//
//
// 	rider := Cowboy{
// 		UUID: "123",
// 		// Age:   2,
// 		// Names: []string{"a", "b"},
// 	}
//
// 	expected := bson.M{
// 		"uuid": "123",
// 	}
//
// 	//
// 	// act
// 	//
//
// 	actual, err := buildUpsertSelector(rider)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	//
// 	// assert
// 	//
//
// 	e := expected["uuid"]
// 	a := actual["uuid"]
// 	if e != a {
// 		t.Errorf("\nexpected:%s\n  actual:%s\n", e, a)
// 		eVal := reflect.ValueOf(e)
// 		eTypeName := eVal.Type().Name()
// 		t.Errorf("\neTypeName:%s\n", eTypeName)
// 		aVal := reflect.ValueOf(a)
// 		aTypeName := aVal.Type().Name()
// 		t.Errorf("\naTypeName:%s\n", aTypeName)
// 	}
//
// 	if !reflect.DeepEqual(expected, actual) {
// 		t.Errorf("\nexpected:%+v\n  actual:%+v\n", expected, actual)
// 	}
// }
//
// func TestBuildUpdatenSelectorEmpty(t *testing.T) {
// 	//
// 	// arrange
// 	//
//
// 	rider := Cowboy{}
// 	expected := bson.M{}
//
// 	//
// 	// act
// 	//
//
// 	actual, err := buildUpsertSelector(rider)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	//
// 	// assert
// 	//
//
// 	if !reflect.DeepEqual(expected, actual) {
// 		t.Errorf("\nexpected:%+v\n  actual:%+v\n", expected, actual)
// 	}
// }
