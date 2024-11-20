package main

import (
	"sync"
	"testing"
	"time"
)

func TestSemaphoreWait(t *testing.T) {
	const maxGoroutines = 3
	const totalGoroutines = 10

	var wg = NewSemaphoreWait(maxGoroutines)
	var mu sync.Mutex
	var counter int

	for i := 0; i < totalGoroutines; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(100 * time.Millisecond)
			mu.Lock()
			counter++
			mu.Unlock()
		}(i)
	}
	wg.Wait()

	if counter != totalGoroutines {
		t.Errorf("Ожидалось завершение %d горутин, завершено %d", totalGoroutines, counter)
	}
}
