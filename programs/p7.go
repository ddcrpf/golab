package programs

import "fmt"

func Program7() {
	code := `
package main

import "fmt"

// Define the LibraryItem interface
type LibraryItem interface {
	borrow()
	returnItem()
	checkAvailability()
}

// Book struct implementing LibraryItem interface
type Book struct {
	title     string
	available bool
}

func (b *Book) borrow() {
	if b.available {
		b.available = false
		fmt.Println(b.title, "has been borrowed.")
	} else {
		fmt.Println(b.title, "is not available.")
	}
}

func (b *Book) returnItem() {
	if !b.available {
		b.available = true
		fmt.Println(b.title, "has been returned.")
	} else {
		fmt.Println(b.title, "was not borrowed.")
	}
}

func (b *Book) checkAvailability() {
	if b.available {
		fmt.Println(b.title, "is available.")
	} else {
		fmt.Println(b.title, "is not available.")
	}
}

func main() {
	// Create a new book
	var item LibraryItem = &Book{title: "Go Programming", available: true}

	// Borrow, check availability, and return the book
	item.checkAvailability()
	item.borrow()
	item.checkAvailability()
	item.returnItem()
}

`
	fmt.Println("Program 7 Code:")
	fmt.Println(code)
}
