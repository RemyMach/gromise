package main

import (
	"fmt"
	"math/rand"
	"time"
)

func one() <-chan int32 {
	r := make(chan int32)

	go func() {
		defer close(r)
		dur := time.Millisecond * time.Duration(rand.Int63n(1500))
		fmt.Println(1, dur)
		time.Sleep(dur)
		r <- 1
	}()

	return r
}

func two() <-chan int32 {
	r := make(chan int32)

	go func() {
		defer close(r)

		// Simulate a workload.
		dur := time.Millisecond * time.Duration(rand.Int63n(10000))
		fmt.Println(2, dur)
		time.Sleep(dur)
		time.Sleep(dur)
		r <- 2
	}()

	return r
}

func main() {
	var r int32
	select {
	case r = <-one():
	case r = <-two():
	}

	fmt.Println(r)
}