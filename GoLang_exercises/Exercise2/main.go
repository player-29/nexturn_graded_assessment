// Exercise 2: Bank Transaction System 
// Topics Covered: Go Constants, Go Loops, Go Break and Continue, Go Functions, Go 
// Strings, Go Errors 
// Case Study: 
// You need to simulate a bank transaction system with the following features: 
// 1. Account Management: Each account has an ID, name, and balance. Store the 
// accounts in a slice. 
// 2. Deposit Function: A function to deposit money into an account. Validate if the 
// deposit amount is greater than zero. 
// 3. Withdraw Function: A function to withdraw money from an account. Ensure the 
// account has a sufficient balance before proceeding. Return appropriate errors 
// for invalid amounts or insufficient balance. 
// 4. Transaction History: Maintain a transaction history for each account as a string 
// slice. Use a loop to display the transaction history when requested. 
// 5. Menu System: Implement a menu-driven program where users can choose 
// actions like deposit, withdraw, view balance, or exit. Use constants for menu 
// options and break the loop to exit. 

package main

import (
	"errors"
	"fmt"
)

type Account struct {
	ID              string
	Name            string
	Balance         float64
	TransactionHistory []string
}

var accounts []Account

func main() {
	for {
		fmt.Println("\nBank Transaction System")
		fmt.Println("1. Create Account")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. View Balance")
		fmt.Println("5. Transaction History")
		fmt.Println("6. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}

		switch choice {
		case 1:
			createAccount()
		case 2:
			depositMoney()
		case 3:
			withdrawMoney()
		case 4:
			viewBalance()
		case 5:
			viewTransactionHistory()
		case 6:
			fmt.Println("Exiting the system. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}

func createAccount() {
	var id, name string
	var initialBalance float64

	fmt.Print("Enter Account ID: ")
	fmt.Scan(&id)
	fmt.Print("Enter Account Holder Name: ")
	fmt.Scan(&name)
	fmt.Print("Enter Initial Balance: ")
	_, err := fmt.Scan(&initialBalance)
	if err != nil {
		fmt.Println("Invalid input for balance. Please enter a valid number.")
		return
	}

	if initialBalance < 0 {
		fmt.Println("Initial balance cannot be negative.")
		return
	}

	for _, account := range accounts {
		if account.ID == id {
			fmt.Println("Account ID already exists. Please use a unique ID.")
			return
		}
	}

	newAccount := Account{
		ID:                id,
		Name:              name,
		Balance:           initialBalance,
		TransactionHistory: []string{"Account created with balance: " + fmt.Sprintf("%.2f", initialBalance)},
	}
	accounts = append(accounts, newAccount)
	fmt.Println("Account created successfully!")
}

func depositMoney() {
	account, err := findAccount()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var amount float64
	fmt.Print("Enter amount to deposit: ")
	_, err = fmt.Scan(&amount)
	if err != nil || amount <= 0 {
		fmt.Println("Invalid deposit amount. Please enter a valid number greater than zero.")
		return
	}

	account.Balance += amount
	account.TransactionHistory = append(account.TransactionHistory, fmt.Sprintf("Deposited: %.2f", amount))
	fmt.Println("Deposit successful!")
}

func withdrawMoney() {
	account, err := findAccount()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var amount float64
	fmt.Print("Enter amount to withdraw: ")
	_, err = fmt.Scan(&amount)
	if err != nil || amount <= 0 {
		fmt.Println("Invalid withdrawal amount. Please enter a valid number greater than zero.")
		return
	}

	if account.Balance < amount {
		fmt.Println("Insufficient balance.")
		return
	}

	account.Balance -= amount
	account.TransactionHistory = append(account.TransactionHistory, fmt.Sprintf("Withdrew: %.2f", amount))
	fmt.Println("Withdrawal successful!")
}

func viewBalance() {
	account, err := findAccount()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Account Balance: %.2f\n", account.Balance)
}

func viewTransactionHistory() {
	account, err := findAccount()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Transaction History:")
	for i, transaction := range account.TransactionHistory {
		fmt.Printf("%d. %s\n", i+1, transaction)
	}
}

func findAccount() (*Account, error) {
	var id string
	fmt.Print("Enter Account ID: ")
	fmt.Scan(&id)

	for i := range accounts {
		if accounts[i].ID == id {
			return &accounts[i], nil
		}
	}

	return nil, errors.New("account not found")
}
