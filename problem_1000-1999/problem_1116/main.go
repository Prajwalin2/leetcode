package main

type ZeroEvenOdd struct {
	n    int
	curr int
	odd  chan int
	even chan int
	zero chan int
}

func NewZeroEvenOdd(n int) *ZeroEvenOdd {
	zeo := &ZeroEvenOdd{
		n:    n,
		curr: 0,
		odd:  make(chan int),
		even: make(chan int),
		zero: make(chan int, 1),
	}
	zeo.zero <- 1
	return zeo
}

func (z *ZeroEvenOdd) Zero(printNumber func(int)) {
	for i := 0; i < z.n; i++ {

		<-z.zero
		printNumber(0)
		if z.curr%2 == 0 {
			z.odd <- 1
		} else {
			z.even <- 1
		}
	}
}

func (z *ZeroEvenOdd) Even(printNumber func(int)) {
	for i := 2; i <= z.n; i += 2 {
		<-z.even
		z.curr++
		printNumber(z.curr)
		z.zero <- 1
	}
}

func (z *ZeroEvenOdd) Odd(printNumber func(int)) {
	for i := 1; i <= z.n; i += 2 {
		<-z.odd
		z.curr++
		printNumber(z.curr)
		z.zero <- 1
	}
}
