package main

//
// CollectionManager
//

// * indexes
// * session management
// * collections

//
// Storer
//

// https://godoc.org/labix.org/v2/mgo
type Storer interface {
	// selector==nil => get all
	// returnValue: slice of values
	Get(selector, returnValue interface{}) error

	// selector==nil => exception
	// value: value or slice of values
	Put(selector, value interface{}) error

	// selector==nil => delete all?
	Delete(selector interface{}) error
}
