package main

//
// partial-update-key-value-store/executer/operator
//

type KeyValueUpdater interface {
	Get(selector, returnValue interface{}) error
	Put(selector, inputValue interface{}) error
}

// GET
//  multiple selector, multiple value, count, error
//  Beispiel: count, err = Get(selectors, &values)
//  guards/expections:
//   selector==nil, selector==bson.M{}
//   value==Gpd{}, value==nil (need: pointer of slice of Gpd)
//   SelectAll() not possible
//
func KeyValueUpdaterGet(executer CollectionExecuter, selector, returnValue interface{}) error {
	return nil
}

// PUT
//  single selector, single value, count, error
//  Beispiel: count, err = Put(selector, &value)
func KeyValueUpdaterPut(executer CollectionExecuter, selector, value interface{}) error {
	return nil
}
