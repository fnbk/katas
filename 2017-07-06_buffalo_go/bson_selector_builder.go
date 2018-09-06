package main

import (
	"fmt"
	"reflect"

	"gopkg.in/mgo.v2/bson"
)

//
// BsonSelectorBuilder()
//

// * creates mongodb selector (mongo db selector builder)
// * match single attribute as ID
// * match multiple attributes as ID
// * match multiple values as ID (match slices)
// * match structs
// * consume Bson annotated struct, produce Bson.M{}

// supports: string
// does not support: int, slice, embedded struct ...
func buildUpsertSelector(selector interface{}) (bson.M, error) {
	upsertSelector := bson.M{}

	val := reflect.ValueOf(selector)
	for i := 0; i < val.NumField(); i++ {

		field := val.Type().Field(i)

		tag := field.Tag.Get("bson")
		if len(tag) == 0 {
			continue
		}

		// die einfache Zuweisung mittels val.Field(i) liefert reflect.Value,
		// und dies hat zur Folge, dass Vergleiche nich mehr funktionieren,
		// denn string ist ungleich reflect.Value
		//
		switch val.Field(i).Kind() {
		case reflect.String:
			strVal := val.Field(i).String()
			if len(strVal) != 0 {
				upsertSelector[tag] = strVal
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			intVal := val.Field(i).Int()
			if intVal != 0 {
				upsertSelector[tag] = intVal
			}
		case reflect.Struct, reflect.Slice:
			// ignore structs and slices as selectors
		default:
			typeName := val.Field(i).Type().Name()
			return bson.M{}, fmt.Errorf("wrong kind:%s", typeName)
		}
	}

	return upsertSelector, nil
}
