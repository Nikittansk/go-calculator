package main

import (
	"fmt"
	"time"
)

func main() {
	go calculateAdd(2, 2)
	go calculateAdd(3, 6)
	go calculateMultiple(7, 7)
	go calculateDivison(9, 3)
	time.Sleep(2 * time.Second)
}

func calculateAdd(a, b int) {
	sum := a + b
	fmt.Printf("%d + %d = %d\n", a, b, sum)
}

func calculateMultiple(a, b int) {
	multiple := a * b
	fmt.Printf("%d * %d = %d\n", a, b, multiple)
}

func calculateDivison(a, b int) {
	divison := a / b
	fmt.Printf("%d / %d = %d\n", a, b, divison)
}
