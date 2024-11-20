package main

import (
	"fmt"
	"time"
)

func main() {
	var wg = NewSemaphoreWait(3)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("Горутина %d начала работу\n", i)
			time.Sleep(time.Second)
			fmt.Printf("Горутина %d завершила работу\n", i)
		}(i)
	}

	wg.Wait()
	fmt.Println("Все горутины завершены.")
}

type SemaphoreWait struct {
	semaphore chan struct{}
	count     int
}

func NewSemaphoreWait(n int) *SemaphoreWait {
	return &SemaphoreWait{
		semaphore: make(chan struct{}, n),
		count:     0,
	}
}

func (wg *SemaphoreWait) Add(delta int) {
	wg.count += delta
	for i := 0; i < delta; i++ {
		wg.semaphore <- struct{}{}
	}
}

func (wg *SemaphoreWait) Done() {
	wg.count--
	<-wg.semaphore
}

func (wg *SemaphoreWait) Wait() {
	for wg.count > 0 {
		time.Sleep(10 * time.Millisecond)
	}
}
