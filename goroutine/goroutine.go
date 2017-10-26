package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello")
}

func world() {
	time.Sleep(1000)
	fmt.Println("Word")
}

func main() {
	go world()
	hello()
	time.Sleep(2000)
}
