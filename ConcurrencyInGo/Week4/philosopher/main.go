package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Philosopher struct {
	number int
	left   *sync.Mutex
	right  *sync.Mutex
}

func (p *Philosopher) eat(w *sync.WaitGroup, sync, ack chan int) {
	times := 0
	for times < 3 {
		sync <- 1
		res := <-ack
		if res == 1 {
			continue
		}
		x := rand.Intn(2)
		if x == 0 {
			p.left.Lock()
			p.right.Lock()
		} else {
			p.right.Lock()
			p.left.Lock()
		}
		fmt.Printf("starting to eat %d\n", p.number)
		fmt.Printf("finishing eating %d\n", p.number)
		p.left.Unlock()
		p.right.Unlock()
		sync <- 2
		times++
	}
	w.Done()
}

func host(sync, ack, fin chan int) {
	nowEating := 0
	for {
		select {
		case i := <-sync:
			if i == 2 {
				nowEating--
			} else {
				if nowEating < 2 {
					nowEating++
					ack <- 0
				} else {
					ack <- 1
				}
			}
		case <-fin:
			return
		}
	}
}

func main() {
	chopsticks := make([]*sync.Mutex, 5)
	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		chopsticks[i] = new(sync.Mutex)
	}
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{i, chopsticks[i], chopsticks[(i+1)%5]}
	}

	var wg sync.WaitGroup
	wg.Add(5)

	sync := make(chan int)
	ack := make(chan int)
	fin := make(chan int)

	go host(sync, ack, fin)
	for i := 0; i < 5; i++ {
		go philosophers[i].eat(&wg, sync, ack)
	}
	wg.Wait()
	fin <- 1
}
