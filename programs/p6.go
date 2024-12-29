package programs

import "fmt"

func Program6() {
	code := `
package main

import "fmt"

// Struct to represent a Book
type Book struct {
	title     string
	available bool
}

// Method to borrow a book
func (b *Book) borrow() {
	if b.available {
		b.available = false
		fmt.Println(b.title, "has been borrowed.")
	} else {
		fmt.Println(b.title, "is not available.")
	}
}

// Method to return a book
func (b *Book) returnBook() {
	if !b.available {
		b.available = true
		fmt.Println(b.title, "has been returned.")
	} else {
		fmt.Println(b.title, "was not borrowed.")
	}
}

// Method to check availability
func (b *Book) checkAvailability() {
	if b.available {
		fmt.Println(b.title, "is available.")
	} else {
		fmt.Println(b.title, "is not available.")
	}
}

func main() {
	// Create a new book
	book1 := Book{title: "Go Programming", available: true}

	// Borrow and return the book
	book1.checkAvailability()
	book1.borrow()
	book1.checkAvailability()
	book1.returnBook()
	book1.checkAvailability()
}


// third varient - with borrower info
// package main

// import "fmt"

// // Struct to represent a Book
// type Book struct {
// 	title     string
// 	available bool
// 	borrower  string
// }

// // Method to borrow a book with borrower's name
// func (b *Book) borrow(name string) {
// 	if b.available {
// 		b.available = false
// 		b.borrower = name
// 		fmt.Printf("%s has been borrowed by %s.\n", b.title, name)
// 	} else {
// 		fmt.Printf("%s is not available, currently borrowed by %s.\n", b.title, b.borrower)
// 	}
// }

// // Method to return a book
// func (b *Book) returnBook() {
// 	if !b.available {
// 		fmt.Printf("%s has been returned by %s.\n", b.title, b.borrower)
// 		b.available = true
// 		b.borrower = ""
// 	} else {
// 		fmt.Println(b.title, "was not borrowed.")
// 	}
// }

// func main() {
// 	// Create a new book
// 	book := Book{title: "Go Programming", available: true}

// 	// Borrow and return the book with borrower's name
// 	book.borrow("Alice")
// 	book.borrow("Bob") // Trying to borrow when already borrowed
// 	book.returnBook()
// 	book.borrow("Bob") // Now Bob can borrow
// }

`
	fmt.Println("Program 6 Code:")
	fmt.Println(code)
}
