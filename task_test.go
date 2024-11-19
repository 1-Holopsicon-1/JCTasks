package main

import (
	"reflect"
	"testing"
)

func TestConveyor(t *testing.T) {
	output := Conveyor(2, 2, 3, 3)
	var checkType float64
	for v := range output {
		if reflect.TypeOf(v) != reflect.TypeOf(checkType) {
			t.Errorf("Error of type")
		}
	}
}
