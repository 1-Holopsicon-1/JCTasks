package main

import (
	"reflect"
	"testing"
)

func TestSliceExample(t *testing.T) {
	slice := []int{2, 4, 6, 5}
	oddNums := append(slice[:3], slice[3+1:]...)
	if !reflect.DeepEqual(SliceExample(slice), oddNums) {
		t.Errorf("Output %q not equal to expected %q", oddNums, SliceExample(slice))
	}

}

func TestAddElement(t *testing.T) {
	slice := []int{2, 4, 6, 5}
	newSlice := []int{2, 4, 6, 5, 7}
	if !reflect.DeepEqual(AddElement(slice, 7), newSlice) {
		t.Errorf("Output %q not equal to expected %q", newSlice, AddElement(slice, 7))
	}
}

func TestCopySlice(t *testing.T) {
	slice := []int{2, 4, 6, 5}
	copySl := CopySlice(slice)
	if !reflect.DeepEqual(copySl, slice) {
		t.Errorf("Output %q not equal to expected %q", copySl, CopySlice(slice))
	}
	copySl = append(copySl, 7)
	if reflect.DeepEqual(copySl, slice) {
		t.Errorf("New slice changed with original")
	}
}

func TestRemoveElement(t *testing.T) {
	slice := []int{2, 4, 6, 5}
	output := RemoveElement(slice, 1)
	expected := []int{2, 6, 5}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("Output %q not equal to expected %q", output, expected)
	}
}
