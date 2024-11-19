package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	ch4 := make(chan int)

	go ChanRandomiser(5, 100, ch1)
	go ChanRandomiser(5, 100, ch2)
	go ChanRandomiser(5, 100, ch3)
	go ChanRandomiser(5, 100, ch4)
	merged := MergeChanels(ch1, ch2, ch3, ch4)
	fmt.Print("Meged chan: ")
	for v := range merged {
		fmt.Printf("%v ", v)
	}
}

func MergeChanels(chans ...<-chan int) <-chan int {
	merge := make(chan int)

	go func() {
		var wg sync.WaitGroup
		readFromCh := func(ch <-chan int) {
			defer wg.Done()
			for v := range ch {
				merge <- v
			}
		}
		for _, v := range chans {
			wg.Add(1)
			go readFromCh(v)
		}
		wg.Wait()
		close(merge)
	}()

	return merge
}

func ChanRandomiser(amount, maxValue int, randomCh chan<- int) {

	for i := 0; i < amount; i++ {
		num := rand.Intn(maxValue)
		randomCh <- num
	}
	close(randomCh)
}
