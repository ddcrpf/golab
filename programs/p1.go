package programs

import "fmt"

func Program1() {
	code := `
package main

import (
	"fmt"
	"math/big"
)

// Iterative approach to calculate factorial using math/big for large numbers
func factorialIterative(n int64) *big.Int {
	result := big.NewInt(1)
	for i := int64(2); i <= n; i++ {
		result.Mul(result, big.NewInt(i))
	}
	return result
}

// Recursive approach to calculate factorial using math/big for large numbers
func factorialRecursive(n int64) *big.Int {
	// Base case: 0! = 1 and 1! = 1
	if n == 0 || n == 1 {
		return big.NewInt(1)
	}
	// Recursive case: n! = n * (n-1)!
	result := big.NewInt(n)
	return result.Mul(result, factorialRecursive(n-1))
}

func main() {
	var num int64
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	// Iterative factorial
	fmt.Printf("Factorial of %d (Iterative): %v\n", num, factorialIterative(num))
	// Recursive factorial
	fmt.Printf("Factorial of %d (Recursive): %v\n", num, factorialRecursive(num))
}

`
	fmt.Println("Program 1 Code:")
	fmt.Println(code)
}
