package main

import (
	"fmt"
)

func main() {
	input := make(chan string)
	output := make(chan string)

	go func() {
		in := <-input
		output <- in + " No. 5"
	}()

	input <- "Channel"
	fmt.Println(<-output)
}
