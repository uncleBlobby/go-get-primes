package main

import (
	"fmt"
	"time"
)

func main() {
	var primeIndex uint64 = 0
	start := time.Now()
	lastPrimeFoundTime := time.Now()
	var number uint64 = 0
	for number = 0; number >= 0; number++ {
		if number%10000000 == 0 {
			fmt.Println(number)
		}
		if !isPrime(number) {
			continue
		} else {
			t := time.Now()
			if primeIndex%1000 == 0 {
				fmt.Printf("%d - %d is prime, %v since last prime, %v elapsed total\n", primeIndex, number, t.Sub(lastPrimeFoundTime), t.Sub(start))
			}

			lastPrimeFoundTime = time.Now()
			primeIndex++
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
