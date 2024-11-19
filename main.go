package main

import "fmt"

func main() {
	chfl := Conveyor(1, 2, 3, 4, 5)
	for v := range chfl {
		fmt.Printf("%v %T \n", v, v)
	}
}

func Conveyor(nums ...int) <-chan float64 {
	ch1 := make(chan uint8)
	ch2 := make(chan float64)
	go func() {
		for _, v := range nums {
			ch1 <- uint8(v)
		}
		close(ch1)
	}()
	go func() {
		for v := range ch1 {
			ch2 <- float64(v * v * v)
		}
		close(ch2)
	}()
	return ch2
}
