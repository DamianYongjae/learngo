package accounts

import (
	"errors"
	"fmt"
)

// Account struct should be exported
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't withdraw")

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

//Deposit x amount on your account
func (theAccount *Account) Deposit(amount int) {
	theAccount.balance += amount
}

//Balance of your account
func (theAccount Account) Balance() int {
	return theAccount.balance
}

//Withdraw x amount from your account
func (theAccount *Account) Withdraw(amount int) error {
	if theAccount.balance < amount {
		return errNoMoney
	}
	theAccount.balance -= amount
	return nil
}

//ChangeOwner of the account
func (theAccount *Account) ChangeOwner(newOwner string) {
	theAccount.owner = newOwner
}

//Owner of the account
func (theAccount Account) Owner() string {
	return theAccount.owner
}

func (theAccount Account) String() string {
	return fmt.Sprint(theAccount.Owner(), "'s account.\nHas: ", theAccount.Balance())
}
