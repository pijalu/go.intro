package main

import "fmt"

func printCutableAndSizable(name string, l livingOrganism, o sizable) {
	fmt.Printf("%s: cute: %t - size in millimiters: %d\n",
		name,
		l.isCute(),
		o.getSizeInMillimiter())
}

func main() {
	acat := &cat{
		root:     livingOrganism{cute: true},
		sizeInCm: 30,
	}

	aspider := &spider{
		root:             livingOrganism{cute: false},
		sizeInMillimiter: 30,
	}

	printCutableAndSizable("My cat", acat.root, acat)
	printCutableAndSizable("Your spider", aspider.root, aspider)
}
