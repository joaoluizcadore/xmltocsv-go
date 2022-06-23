package service

import (
	"testing"
)

func TestValidateTemplateRow(t *testing.T) {
	table := []struct {
		input       string
		expectError bool
	}{
		{input: "COLUMN_01;//@TEST;ATTR", expectError: false},
		{input: "COLUMN_01;//@TEST;", expectError: true},
		{input: "COLUMN_01;//@TEST", expectError: true},
		{input: "COLUMN_01;//@TEST;ATTRXXX", expectError: true},
		{input: "COLUMN_01;//@TEST;TEXT", expectError: false},
		{input: "COLUMN_01;//@TEST;ATTR;TEMP", expectError: true},
	}
	for _, test := range table {
		err := validateTemplateRow(test.input)
		if test.expectError && err == nil {
			t.Errorf("Input: %v -> didn't get error, but should!", test.input)
		}
		if !test.expectError && err != nil {
			t.Errorf("Input %v -> should not get erro, but got %v", test.input, err.Error())
		}
	}

}
