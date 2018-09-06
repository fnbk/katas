package main

//
// KeyValueReplacer
//

type KeyValueReplacer interface {
	Get(selector, returnValue interface{}) error
	Put(selector, inputValue interface{}) error
}
