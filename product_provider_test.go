package main

import (
	"reflect"
	"testing"
)

func TestProductIDs(t *testing.T) {
	expected := []string{"1", "2"}
	provider := ProductProvider{Products: []Product{{ID: "1"}, {ID: "2"}}}
	actual := provider.ProductIDs()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("ProductIDs failed!\nexpected:%+v\nactual:  %+v\n", expected, actual)
	}
}

func TestGetProduct(t *testing.T) {
	expected := &Product{
		ID:   "123",
		Name: "Name123",
	}
	provider := ProductProvider{Products: []Product{
		{
			ID:   "123",
			Name: "Name123",
		},
		{
			ID:   "789",
			Name: "Name789",
		},
	}}
	actual := provider.GetProduct("123")
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("ProductIDs failed!\nexpected:%+v\nactual:  %+v\n", expected, actual)
	}
}
