package programs

import "fmt"

func Program2() {
	code := `
package main

import "fmt"

// Define a struct to represent a BankAccount
type BankAccount struct {
	accountHolder string
	balance       float64
}

// Method to deposit money into the account
func (b *BankAccount) Deposit(amount float64) {
	b.balance += amount
	fmt.Printf("Deposited: %.2f\n", amount)
}

// Method to withdraw money from the account
func (b *BankAccount) Withdraw(amount float64) {
	if amount > b.balance {
		fmt.Println("Insufficient funds!")
	} else {
		b.balance -= amount
		fmt.Printf("Withdrawn: %.2f\n", amount)
	}
}

// Method to check the current balance
func (b *BankAccount) CheckBalance() float64 {
	return b.balance
}

func main() {
	var accountHolder string
	var initialBalance, depositAmount, withdrawAmount float64

	// Get input for account holder's name and initial balance
	fmt.Print("Enter the account holder's name: ")
	fmt.Scanln(&accountHolder)
	fmt.Print("Enter the initial balance: ")
	fmt.Scanln(&initialBalance)

	// Create a new bank account with dynamic input
	account := BankAccount{accountHolder: accountHolder, balance: initialBalance}

	// Deposit money into the account
	fmt.Print("Enter the amount to deposit: ")
	fmt.Scanln(&depositAmount)
	account.Deposit(depositAmount)

	// Withdraw money from the account
	fmt.Print("Enter the amount to withdraw: ")
	fmt.Scanln(&withdrawAmount)
	account.Withdraw(withdrawAmount)

	// Check and print the current balance
	fmt.Printf("Current balance for %s: %.2f\n", account.accountHolder, account.CheckBalance())
}


`
	fmt.Println("Program 2 Code:")
	fmt.Println(code)
}
