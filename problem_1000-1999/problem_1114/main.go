package main

import "sync"

type H2O struct {
	H     chan struct{}
	O     chan struct{}
	sync1 sync.Mutex
	sync2 sync.Mutex
}

func NewH2O() *H2O {
	h := &H2O{
		H: make(chan struct{}, 2),
		O: make(chan struct{}, 2),
	}
	h.H <- struct{}{}
	h.H <- struct{}{}
	return h
}

func (h *H2O) Hydrogen(releaseHydrogen func()) {
	h.sync1.Lock()
	<-h.H
	// releaseHydrogen() outputs "H". Do not change or remove this line.
	releaseHydrogen()
	h.O <- struct{}{}
	h.sync1.Unlock()
}

func (h *H2O) Oxygen(releaseOxygen func()) {
	h.sync2.Lock()
	<-h.O
	<-h.O
	// releaseOxygen() outputs "H". Do not change or remove this line.
	releaseOxygen()
	h.H <- struct{}{}
	h.H <- struct{}{}
	h.sync2.Unlock()
}
