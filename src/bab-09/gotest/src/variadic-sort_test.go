package main

import "testing"

func TestDoSort(t *testing.T) {
	result := DoSort(8, 5, 1, 4)
	if result[0] != 1 {
		t.Errorf("Sort was incorrect, got: %v, want: %d.", result, 1)
	}
}
