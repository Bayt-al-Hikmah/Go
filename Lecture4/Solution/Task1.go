package main

import (
	"fmt"
	"errors"
)

type BankAccount struct{
	balance float64
}

func (b * BankAccount) Deposit(amount float64){
	b.balance += amount
}

func (b * BankAccount) Withdraw(amount float64) error{
	if amount > b.balance {
		return errors.New("insufficient funds")
	}

	b.balance -= amount
	return nil
}

func main(){
	var ali_account BankAccount
	ali_account.balance = 500
	fmt.Printf("You starting balance is %f \n",ali_account.balance)
	ali_account.Deposit(700)
	fmt.Printf("Balance now is %f \n",ali_account.balance)
	err := ali_account.Withdraw(1700)
	if err != nil{
		fmt.Println("We don't have enough to withdraw %s",err)
	}else{
		fmt.Println("Balance now is ",ali_account.balance)
	}

}