package main

type cutable interface {
	isCute() bool
}

type sizable interface {
	getSizeInMillimiter() int
}

type cutableSizable interface {
	cutable
	sizable
}
