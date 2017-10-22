package main

type livingOrganism struct {
	cute bool
}

type cat struct {
	root     livingOrganism
	sizeInCm int
}

type spider struct {
	root             livingOrganism
	sizeInMillimiter int
}
