package main

import "fmt"

func superFunc(i int) (int, int) {
	return i, i + 19
}

func main() {
	a, b := superFunc(23)
	fmt.Printf("And the returns are %d and %d\n",
		a,
		b)
}
