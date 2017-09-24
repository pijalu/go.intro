package main

import (
	"fmt"

	"github.com/pijalu/go.intro/apackage"
)

func main() {
	val := 41
	fmt.Printf("Addone(%d):%d\n", val, apackage.Addone(val))
}
