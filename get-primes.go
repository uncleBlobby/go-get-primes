package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	var number uint64 = 0
	for number = 0; number >= 0; number++ {
		if number%10000000 == 0 {
			fmt.Println(number)
		}
		if !isPrime(number) {
			continue
		} else {
			t := time.Now()
			fmt.Printf("%d is prime, %v seconds elapsed\n", number, t.Sub(start))
		}

	}
	fmt.Println("Hello, world!")
}

func isPrime(target uint64) bool {
	var start uint64 = 2
	for start = 2; start < target/2; start++ {
		if target%start == 0 {
			return false
		}
	}
	return true
}
