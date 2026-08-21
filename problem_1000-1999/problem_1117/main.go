package problem1117

type H2O struct {
	H chan struct{}
	O chan struct{}
}

func NewH2O() *H2O {
	h := &H2O{
		H: make(chan struct{}, 2),
		O: make(chan struct{}),
	}
	h.H <- struct{}{}
	h.H <- struct{}{}
	return h
}

func (h *H2O) Hydrogen(releaseHydrogen func()) {
	<-h.H
	// releaseHydrogen() outputs "H". Do not change or remove this line.
	releaseHydrogen()
	h.O <- struct{}{}
}

func (h *H2O) Oxygen(releaseOxygen func()) {
	<-h.O
	// releaseOxygen() outputs "H". Do not change or remove this line.
	releaseOxygen()
	h.H <- struct{}{}
	h.H <- struct{}{}
}
