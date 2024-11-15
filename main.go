package main

import "fmt"

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	fmt.Println(FindUnMatch(slice1, slice2))
}

func FindUnMatch(s1, s2 []string) (s3 []string) {
	exist := make(map[string]interface{})
	for _, v := range s2 {
		exist[v] = nil
	}
	for _, v := range s1 {
		if _, ok := exist[v]; !ok {
			s3 = append(s3, v)
		}
	}
	return s3
}
