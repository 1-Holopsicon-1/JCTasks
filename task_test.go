package main

import (
	"testing"
)

func TestChanRandomiser(t *testing.T) {
	maxValue := 100
	count := 10
	randomCh := make(chan int)
	go ChanRandomiser(count, maxValue, randomCh)
	counter := 0
	for v := range randomCh {
		if v > 100 {
			t.Errorf("Max value more then sended %v", v)
		}
		counter++
	}
	if counter != 10 {
		t.Errorf("Not enough items %v", counter)
	}
}
