package main

import (
	"reflect"
	"testing"
)

func TestProductIDs(t *testing.T) {
	expected := []string{"1", "2"}
	provider := ProductProvider{IDs: expected}
	actual := provider.ProductIDs()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("ProductIDs failed!\nexpected:%+v\nactual:  %+v\n", expected, actual)
	}
}
