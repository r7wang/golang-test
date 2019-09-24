package test

import (
	"sync"
	"testing"
	"time"
)

func TestWaitOnGoroutine(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		t.Log("Goroutine")
	}()
	wg.Wait()
}

func TestGoroutineClosure(t *testing.T) {
	val := 0
	for goroutineIdx := 0; goroutineIdx < 5; goroutineIdx++ {
		go func() {
			time.Sleep(time.Millisecond * 5)
			t.Log("Goroutine")
			val++
		}()
	}
	for waitIdx := 0; waitIdx < 10; waitIdx++ {
		t.Logf("Current value: %d", val)
		time.Sleep(time.Millisecond * 1)
	}
}
