package main

type livingOrganism struct {
	cute bool
}

type cat struct {
	livingOrganism
	sizeInCm int
}

type spider struct {
	livingOrganism
	sizeInMillimiter int
}
