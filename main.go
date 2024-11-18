package main

import (
	"fmt"
	"math/rand"
)

func main() {
	maxValue := 100
	count := 10
	randomCh := make(chan int)
	go ChanRandomiser(count, maxValue, randomCh)
	for v := range randomCh {
		fmt.Println(v)
	}
}

func ChanRandomiser(amount, maxValue int, randomCh chan<- int) {

	for i := 0; i < amount; i++ {
		num := rand.Intn(maxValue)
		randomCh <- num
	}
	close(randomCh)
}
