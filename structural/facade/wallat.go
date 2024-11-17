package main

import "fmt"

type Wallat struct {
	balance int
}

func newWallat() *Wallat {
	return &Wallat{
		balance: 0,
	}
}

func (w *Wallat) creditWallat(amount int) {
	w.balance += amount
	fmt.Println("wallat balance added successfully")
	return
}

func (w *Wallat) debitBalance(amount int) error {
	if w.balance < amount {
		return fmt.Errorf("Balance is not sufficient")
	}
	fmt.Println("Wallet balance is Sufficient")
	w.balance = w.balance - amount
	return nil
}
