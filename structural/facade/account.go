package main

import "fmt"

type Account struct {
	accountId string
}

func newAccount(accountId string) *Account {
	return &Account{
		accountId: accountId,
	}
}

func (ac *Account) checkAccount(accountId string) error {
	if ac.accountId != accountId {
		return fmt.Errorf("Account name is incorrect")
	}

	fmt.Println("Account Verified")
	return nil
}
