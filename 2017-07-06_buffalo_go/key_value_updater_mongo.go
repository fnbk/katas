package main

type KeyValueUpdaterMongo struct {
}

func NewKeyValueUpdaterMongo(executer CollectionExecuter) KeyValueUpdater {
	return nil
}

//
// Updater
//

// * operate on multiple values

// GET
//  single selector, single value, error
//  Beispiel: err = Get(selector, &values)
//  guards/expections:
//   selector==nil, selector==bson.M{}
//	 input: value==reflect.Slice, (need: pointer of []Gpd)
//   result: value==&[]Gpd{}, value==nil
//
func KeyValueUpdaterGet(executer CollectionExecuter, selector, returnValue interface{}) error {
}

// PUT (update all found)
//  single selector, single value, error (found?)
//  Beispiel: err = Put(selector, &value)
//  guards/expections:
//   selector==nil, selector==bson.M{}
//	 input: value!=reflect.Slice, (need: pointer of Gpd)
//   result: value==&Gpd{}, value==nil
//
func KeyValueUpdaterPut(executer CollectionExecuter, selector, value interface{}) error {
}
