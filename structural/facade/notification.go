package main

import "fmt"

type Notification struct {
}

func (n *Notification) sendAccountNotification() {
	fmt.Println("send account notification")
}
