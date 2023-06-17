package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result struct {
	Value int32
	Error error
}

func longRunningTask() <-chan Result {
	r := make(chan Result)

	go func() {
		defer close(r)
		
		// Simulate a workload.
		time.Sleep(time.Second * 3)
		
		if rand.Float32() < 0.5 { // Simulate an error 50% of the time.
			r <- Result{Error: fmt.Errorf("an error occurred")}
			return
		}
		
		r <- Result{Value: rand.Int31n(100)}
	}()

	return r
}

func main() {
	aCh, bCh, cCh := longRunningTask(), longRunningTask(), longRunningTask()
	a, b, c := <-aCh, <-bCh, <-cCh
	
	if a.Error != nil || b.Error != nil || c.Error != nil {
		fmt.Println("An error occurred in one or more goroutines")
		return
	}
	
	fmt.Println(a.Value, b.Value, c.Value)
}
