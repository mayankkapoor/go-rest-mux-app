// main_test.go

package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	total := Sum(5, 4)
	expected := 9
	if total != expected {
		t.Errorf("Sum was incorrect, got: %d, expected: %d.", total, expected)
	}

}
