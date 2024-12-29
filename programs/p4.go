package programs

import "fmt"

func Program4() {
	code := `
package main
import "fmt"
// Function to swap two integers using pointers
func swap(a, b *int) {
temp := *a
*a = *b
*b = temp
}
func main() {
var x, y int
fmt.Print("Enter two integers: ")
fmt.Scan(&x, &y)
fmt.Printf("Before swap: x = %d, y = %d\n", x, y)
swap(&x, &y)
fmt.Printf("After swap: x = %d, y = %d\n", x, y)
}

// Variant 3 : Swap with Bitwise XOR
// package main
// import "fmt"
// // Function to swap two integers using bitwise XOR
// func swap(a, b *int) {
// *a = *a ^ *b
// *b = *a ^ *b
// *a = *a ^ *b
// }
// func main() {
// var x, y int
// fmt.Print("Enter two integers: ")
// fmt.Scan(&x, &y)
// fmt.Printf("Before swap: x = %d, y = %d\n", x, y)
// swap(&x, &y)
// fmt.Printf("After swap: x = %d, y = %d\n", x, y)
// }
`
	fmt.Println("Program 4 Code:")
	fmt.Println(code)
}
