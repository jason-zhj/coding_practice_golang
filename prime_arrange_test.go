package main

import (
	"testing"
	"fmt"
)

func numPrimeArrangements(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	numPrimes := 1
	for i:=3; i<=n; i++ {
		if isPrime(i) {
			numPrimes++
		}
	}

	return factorial(numPrimes) * factorial(n - numPrimes)
}

func isPrime(num int) bool {
	for i:=2; i<=num/2; i++ {
		if num % i == 0 {
			return false
		}
	}

	return true
}

func factorial(num int) int {
	p := 1
	for i:=1; i<=num; i++ {
		p *= i
	}
	return p
}

func TestNumPrimeArrange(t *testing.T) {
	fmt.Println(numPrimeArrangements(5))
}
