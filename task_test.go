package main

import (
	"reflect"
	"testing"
)

func TestFindMatch(t *testing.T) {
	arr1 := []int{65, 3, 58, 678, 64}
	arr2 := []int{64, 2, 3, 43}
	arr3 := []int{64, 3}
	if outB, outAr := FindMatch(arr1, arr2); !reflect.DeepEqual(outAr, arr3) || !outB {
		t.Errorf("Output %v, %t not equal to expected %v", outAr, outB, arr3)

	}
	arr1 = []int{65, 3, 58, 678, 64}
	arr2 = []int{77, 77, 77, 77}
	arr3 = []int{}
	if outB, outAr := FindMatch(arr1, arr2); reflect.DeepEqual(outAr, arr3) || outB {
		t.Errorf("Output %v, %t not equal to expected %v", outAr, outB, arr3)

	}
}
