package main

type FooBar struct {
	n  int
	c1 chan int
	c2 chan int
}

func NewFooBar(n int) *FooBar {
	fb := &FooBar{
		n:  n,
		c1: make(chan int),
		c2: make(chan int),
	}
	fb.c2 <- 1
	return fb
}

func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
		// printFoo() outputs "foo". Do not change or remove this line.
		<-fb.c2
		printFoo()
		fb.c1 <- 1
	}
}

func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
		// printBar() outputs "bar". Do not change or remove this line.
		<-fb.c1
		printBar()
		fb.c2 <- 1
	}
}
