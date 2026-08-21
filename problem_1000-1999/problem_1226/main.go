package main

import "sync"

type DiningPhilosophers struct {
	spoons []chan struct{}
	seat   chan struct{}
	once   sync.Once
}

func (this *DiningPhilosophers) wantsToEat(
	philosopher int,
	pickLeftFork func(),
	pickRightFork func(),
	eat func(),
	putLeftFork func(),
	putRightFork func(),
) {
	this.once.Do(func() {
		this.spoons = make([]chan struct{}, 5)
		for i := range 5 {
			this.spoons[i] = make(chan struct{}, 1)
			this.spoons[i] <- struct{}{}
		}
		this.seat = make(chan struct{}, 4)
		for i := 0; i < 4; i++ {
			this.seat <- struct{}{}
		}
	})
	<-this.seat
	left, right := (philosopher+1)%5, philosopher
	<-this.spoons[left]
	pickLeftFork()
	<-this.spoons[right]
	pickRightFork()
	eat()
	putLeftFork()
	this.spoons[left] <- struct{}{}
	putRightFork()
	this.spoons[right] <- struct{}{}
	this.seat <- struct{}{}
	// TODO: implement your solution here
}
