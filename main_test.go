//main_test.go

package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	total := Sum(5, 4)
	if total != 9 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 9)
	}

}
