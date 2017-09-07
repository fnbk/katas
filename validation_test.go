package main

import "testing"

func TestValidateDatatype(t *testing.T) {
	// arrange
	testcases := []struct {
		Value    string
		Datatype string
		Result   string
	}{
		// string
		{Value: "abc", Datatype: "string", Result: ""},
		{Value: "1", Datatype: "string", Result: ""},
		{Value: "1.0", Datatype: "string", Result: ""},

		// int32
		{Value: "abc", Datatype: "int32", Result: "'abc' is not a valid datatype 'int32'"},
		{Value: "1", Datatype: "int32", Result: ""},
		{Value: "1.0", Datatype: "int32", Result: "'1.0' is not a valid datatype 'int32'"},

		// float32
		{Value: "abc", Datatype: "float32", Result: "'abc' is not a valid datatype 'float32'"},
		{Value: "1", Datatype: "float32", Result: "'1' is not a valid datatype 'float32'"},
		{Value: "1.0", Datatype: "float32", Result: ""},
		{Value: "3.402823e+38", Datatype: "float32", Result: ""},
		{Value: "3.402823e+39", Datatype: "float32", Result: "'3.402823e+39' is not a valid datatype 'float32'"},
		{Value: "4321567890123456789012345678901234567890.0", Datatype: "float32", Result: "'4321567890123456789012345678901234567890.0' is not a valid datatype 'float32'"},
	}

	// act
	for i, tt := range testcases {
		expected := tt.Result
		actual := ValidateDatatype(tt.Value, tt.Datatype)

		// assert
		if expected != actual {
			t.Errorf("\ncase:%d Value:%s Datatype:%s\nexpected:%s\nactual:%s\n", i, tt.Value, tt.Datatype, expected, actual)
		}
	}
}
