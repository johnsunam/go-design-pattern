package main

import "fmt"

type WallatFacade struct {
	account      *Account
	notification *Notification
	wallat       *Wallat
}

func newWallatFacade(accId string) *WallatFacade {
	return &WallatFacade{
		account:      newAccount(accId),
		notification: &Notification{},
		wallat:       newWallat(),
	}
}

func (wf *WallatFacade) addMoneyToWallat(accId string, amount int) error {
	fmt.Println("starting add money to wallat")
	err := wf.account.checkAccount(accId)
	if err != nil {
		return err
	}
	wf.wallat.creditWallat(amount)
	return nil
}
