package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var originalSlice []int
	for i := 0; i < 10; i++ {
		originalSlice = append(originalSlice, rand.Intn(100))
	}

	fmt.Println(SliceExample(originalSlice))
	fmt.Println(AddElement(originalSlice, 5))
	fmt.Println(CopySlice(originalSlice))
	fmt.Println(RemoveElement(originalSlice, 2))

}

func SliceExample(orSl []int) []int {
	var newSl []int
	for _, v := range orSl {
		if v%2 == 0 {
			newSl = append(newSl, v)
		}
	}
	return newSl
}

func AddElement[T any](slice []T, el ...T) []T {
	totalLength := len(slice) + len(el)
	if totalLength <= cap(slice) {
		slice = slice[:totalLength]
	} else {
		newSlice := make([]T, totalLength)
		copy(newSlice, slice)
		slice = newSlice
	}
	copy(slice[len(slice)-len(el):], el)
	return slice
}

func CopySlice[T any](slice []T) []T {
	var sliceRet []T
	for _, v := range slice {
		sliceRet = append(sliceRet, v)
	}
	return sliceRet
}

func RemoveElement[T any](slice []T, index int) []T {
	sliceRt := slice[:index]
	sliceRt = append(sliceRt, slice[index+1:]...)
	return sliceRt
}
