package programs

import "fmt"

func Program5() {
	code := `
package main
import "fmt"
// Recursive function to compute Fibonacci number
func fibonacci(n int) int {
if n <= 0 {
return 0
} else if n == 1 {
return 1
}
return fibonacci(n-1) + fibonacci(n-2)
}
func main() {
var num int
fmt.Print("Enter the number of Fibonacci terms to compute: ")
fmt.Scan(&num)
fmt.Printf("Fibonacci sequence up to %d terms:\n", num)
for i := 0; i < num; i++ {
fmt.Printf("%d ", fibonacci(i))
}
fmt.Println()
}

// //Variant 2: Memoized Recursive Fibonacci
// package main
// import "fmt"
// // Memoization array to store Fibonacci numbers
// var memo = make(map[int]int)
// // Recursive function to compute Fibonacci number with memoization
// func fibonacci(n int) int {
// if n <= 0 {
//     return 0
// } else if n == 1 {
// return 1
// }
// if val, found := memo[n]; found {
// return val
// }
// // Store the computed Fibonacci number in the map
// memo[n] = fibonacci(n-1) + fibonacci(n-2)
// return memo[n]
// }
// func main() {
// var num int
// fmt.Print("Enter the number of Fibonacci terms to compute: ")
// fmt.Scan(&num)
// fmt.Printf("Fibonacci sequence up to %d terms:\n", num)
// for i := 0; i < num; i++ {
// fmt.Printf("%d ", fibonacci(i))
// }
// fmt.Println()
// }
`
	fmt.Println("Program 5 Code:")
	fmt.Println(code)
}
