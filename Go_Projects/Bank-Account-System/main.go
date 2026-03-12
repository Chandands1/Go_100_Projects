package main

import (
	"fmt"
	"strings"
)

type Account struct {
	Name          string
	AccountNumber int
	Balance       float64
}

func main() {
	startBanking()
}

func startBanking() {
	account := createAccount()
	bankMenu(&account)
}

func createAccount() Account {
	var acc Account

	fmt.Print("Enter your name: ")
	fmt.Scanln(&acc.Name)

	fmt.Print("Enter your Account Number: ")
	fmt.Scanln(&acc.AccountNumber)

	fmt.Print("Enter initial deposit: ")
	fmt.Scanln(&acc.Balance)

	fmt.Println("Account created successfully!")

	return acc
}

func bankMenu(acc *Account) {
	for {
		fmt.Println("\nSelect Operation:")
		fmt.Println("D -> Deposit")
		fmt.Println("B -> Check Balance")
		fmt.Println("W -> Withdraw")
		fmt.Println("E -> Exit")

		var choice string
		fmt.Print("Enter choice: ")
		fmt.Scanln(&choice)

		choice = strings.ToUpper(choice)

		switch choice {

		case "D":
			deposit(acc)

		case "B":
			checkBalance(acc)

		case "W":
			withdraw(acc)

		case "E":
			fmt.Println("Thank you for using the banking system!")
			return

		default:
			fmt.Println("Invalid option. Try again.")
		}
	}
}

func deposit(acc *Account) {
	var amount float64

	fmt.Print("Enter amount to deposit: ")
	fmt.Scanln(&amount)

	if amount <= 0 {
		fmt.Println("Enter a valid amount")
		return
	}

	acc.Balance += amount

	fmt.Println("Deposit successful!")
	fmt.Println("Current Balance:", acc.Balance)
}

func checkBalance(acc *Account) {
	fmt.Println("Current Balance:", acc.Balance)
}

func withdraw(acc *Account) {
	var amount float64

	fmt.Print("Enter amount to withdraw: ")
	fmt.Scanln(&amount)

	if amount <= 0 {
		fmt.Println("Enter a valid amount")
		return
	}

	if amount > acc.Balance {
		fmt.Println("Insufficient balance")
		return
	}

	acc.Balance -= amount

	fmt.Println("Withdrawal successful!")
	fmt.Println("Remaining Balance:", acc.Balance)
}