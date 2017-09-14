package main

import (
	"reflect"
	"testing"
)

func TestGetStructure(t *testing.T) {
	expected := &Structure{
		ID:   "123",
		Name: "Name123",
	}
	provider := StructureProvider{Structures: []Structure{
		{
			ID:   "123",
			Name: "Name123",
		},
		{
			ID:   "789",
			Name: "Name789",
		},
	}}
	actual := provider.GetStructure("123")
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("ProductIDs failed!\nexpected:%+v\nactual:  %+v\n", expected, actual)
	}
}
