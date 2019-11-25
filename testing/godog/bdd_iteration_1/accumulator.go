package accumulator

type acc struct {
	value int
}

func (a *acc) accumulate(x int) {
	a.value += x
}
