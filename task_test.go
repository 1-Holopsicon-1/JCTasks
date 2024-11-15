package main

import (
	"reflect"
	"testing"
)

func TestFindUnMatch(t *testing.T) {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	unmached := []string{"apple", "cherry", "43", "lead", "gno1"}

	if s3 := FindUnMatch(slice1, slice2); !reflect.DeepEqual(unmached, s3) {
		t.Errorf("Output %q not equal to expected %q", s3, unmached)

	}
}
