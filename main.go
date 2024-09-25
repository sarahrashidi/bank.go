package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"github.com/sarahrashidi/bank.go/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {

	presentOptions()

	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
		fmt.Println("-----------")
		panic("Can't continue, sorry!")

	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7!", randomdata.PhoneNumber())
	for i := 0; i < int(accountBalance); i++ {

		var choice int
		fmt.Print("Your choice : ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Your balance is : ", accountBalance)
		} else if choice == 2 {

			fmt.Println("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid deposit amount, must be greater than zero")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)

		} else if choice == 3 {
			fmt.Println("Withdrawal amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount > accountBalance {
				fmt.Println("Invalid withdraw amount, greater than aacount balance!")
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Println("Blance updated! New amount: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)

		} else {
			fmt.Println("Exit!")
			fmt.Println("Have a nice day!")
		}

	}
}
