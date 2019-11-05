package main

import (
	"fmt"
	"sync"
)

func main() {
	var waitgroup sync.WaitGroup
	waitgroup.Add(1)
	go one(&waitgroup)
	go two(&waitgroup)
	// which finishes first is non-deterministic, hence race condition
	waitgroup.Wait()
}

func one(waitgroup *sync.WaitGroup) {
	fmt.Println(1)
	waitgroup.Done()
}

func two(waitgroup *sync.WaitGroup) {
	fmt.Println(2)
	waitgroup.Done()
}
