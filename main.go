package main

import "fmt"

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}
	fmt.Println(FindMatch(a, b))
}

func FindMatch(a1, a2 []int) (bool, []int) {
	var a3 []int
	found := false
	for _, v := range a2 {
		for _, v2 := range a1 {
			if v2 == v {
				a3 = append(a3, v2)
			}
		}
	}
	if len(a3) != 0 {
		found = true
	}
	return found, a3
}
