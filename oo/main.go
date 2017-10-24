package main

import "fmt"

func printCutableAndSizable(name string, i cutableSizable) {
	fmt.Printf("%s: cute: %t - size in millimiters: %d\n",
		name,
		i.isCute(),
		i.getSizeInMillimiter())
}

func main() {
	acat := &cat{
		livingOrganism: livingOrganism{cute: true},
		sizeInCm:       30,
	}

	aspider := &spider{
		livingOrganism:   livingOrganism{cute: false},
		sizeInMillimiter: 30,
	}

	printCutableAndSizable("My cat", acat)
	printCutableAndSizable("Your spider", aspider)
}
