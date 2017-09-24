package main

import "fmt"

func adderX(increment int) func(value int) (int, int) {
	calls := 0
	return func(value int) (int, int) {
		calls++
		return value + increment, calls
	}
}

func main() {
	adder22 := adderX(22)
	adder41 := adderX(41)

	for i := 0; i < 2; i++ {
		res, call := adder41(1)
		fmt.Printf("1 + adderX(41) = %d (calls: %d)\n", res, call)
		res, call = adder22(1)
		fmt.Printf("1 + adderX(22) = %d (calls: %d)\n", res, call)
	}
}
