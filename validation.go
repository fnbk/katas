package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func ValidateDatatype(value, datatype string) string {
	switch datatype {
	case "int32":
		_, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			return fmt.Sprintf("'%s' is not a valid datatype 'int32'", value)
		}
	case "float32":
		match, _ := regexp.MatchString("(\\.)", value)
		if !match {
			return fmt.Sprintf("'%s' is not a valid datatype 'float32'", value)
		}
		_, err := strconv.ParseFloat(value, 32)
		if err != nil {
			return fmt.Sprintf("'%s' is not a valid datatype 'float32'", value)
		}
	}
	return ""
}
