package main

import (
	"testing"
)

func TestMergeChanels(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	go func() {
		for i := 5; i < 10; i++ {
			ch2 <- i
		}
		close(ch2)
	}()
	go func() {
		for i := 10; i < 15; i++ {
			ch3 <- i
		}
		close(ch3)
	}()
	merged := MergeChanels(ch1, ch2, ch3)
	expected := make(map[int]interface{})
	for i := 0; i < 15; i++ {
		expected[i] = ""
	}
	var findUnmatched bool
	for v := range merged {
		_, findUnmatched = expected[v]
		if !findUnmatched {
			t.Errorf("Error cant find existing variable")
		}
	}

}
