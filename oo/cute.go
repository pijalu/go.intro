package main

type cutable interface {
	isCute()
}

func (l *livingOrganism) isCute() bool {
	return l.cute
}
