package main

import (
	"fmt"
	"reflect"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//
// Mongo
//

type KeyValueReplacerMongo struct {
}

func NewKeyValueReplacerMongo(executer CollectionExecuter) KeyValueReplacer {
	return nil
}

func (self *KeyValueReplacerMongo) Get(selector, returnValue interface{}) error {
	return nil
}

func (self *KeyValueReplacerMongo) Put(selector, inputValue interface{}) error {
	return nil
}

//
// replace-key-value-store/executer/operator
//

// * operate on single value (if many found then error) // very strict

// GET
//  single selector, single value, error (found?)
//  Beispiel: err = Get(selector, &value)
//  guards/expections:
//   selector==nil, selector==bson.M{}
//	 input: value!=reflect.Slice, (need: pointer of Gpd)
//   result: value==&Gpd{}, value==nil (need: pointer of Gpd)
//
func KeyValueReplacerGet(executer CollectionExecuter, selector, returnValue interface{}) error {
	var err error

	//
	// guard
	//

	//
	// if returnValue.Kind() != reflect.Ptr || resultv.Elem().Kind() == reflect.Slice {
	// 	panic("result argument must NOT be a slice address")
	// }

	//
	// selector
	//

	var findSelector bson.M
	if selector != nil {
		findSelector, err = buildUpsertSelector(selector)
		if err != nil {
			return err
		}
		if reflect.DeepEqual(findSelector, bson.M{}) {
			return fmt.Errorf("selector ist leeres Struct")
		}
	} else {
		findSelector = bson.M{}
	}

	//
	// query
	//

	sliceType := reflect.SliceOf(reflect.ValueOf(returnValue).Type())
	returnValueSlice := reflect.MakeSlice(sliceType, 0, 0)

	fn := func(collection *mgo.Collection) error {
		err := collection.Find(findSelector).All(&returnValueSlice)
		if err != nil {
			return err
		}

		len := returnValueSlice.Len()
		if len > 1 {
			return fmt.Errorf("Anzahl der gefunden Dokumente is größer als 1. Daten zum entsprechenden Key sind mehrfach vorhanden.")
		}
		if len == 1 {
			*returnValue = returnValueSlice.Index(0)
		}

		return nil
	}

	//
	// execute
	//

	err = executer.Execute(fn)
	if err != nil {
		return err
	}

	return nil
}

// PUT (upsert)
//  single selector, single value, error (found?)
//  Beispiel: err = Put(selector, &value)
//  guards/expections:
//   selector==nil, selector==bson.M{}
//	 input: value!=reflect.Slice, (need: pointer of Gpd)
//   result: value==&Gpd{}, value==nil
//
func KeyValueReplacerPut(executer CollectionExecuter, selector, value interface{}) error {
	//
	// guard
	//

	if reflect.DeepEqual(selector, nil) {
		return fmt.Errorf("kein gueltiger selector: nil")
	}

	//
	// selector
	//

	upsertSelector, err := buildUpsertSelector(selector)
	if err != nil {
		return err
	}
	if reflect.DeepEqual(upsertSelector, bson.M{}) {
		return fmt.Errorf("kein selector angegeben")
	}

	updateValue := bson.M{"$set": value}

	//
	// query
	//

	fn := func(collection *mgo.Collection) error {
		changeInfo, err := collection.Upsert(upsertSelector, updateValue)
		if err != nil {
			return err
		}
		// bei einem Put sollte nur ein Dokument betroffen sein
		if changeInfo.Matched > 1 {
			errMsg := "Anzahl der gefunden Dokumente is größer als 1. Daten zum entsprechenden Key sind mehrfach vorhanden."
			return fmt.Errorf(errMsg)
		}
		return nil
	}

	//
	// execute
	//

	err = executer.Execute(fn)
	if err != nil {
		return err
	}

	return nil
}
