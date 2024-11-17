package main

import "fmt"

func main() {
	fmt.Println()
	walletFaced := newWallatFacade("sfsdf")
	err := walletFaced.addMoneyToWallat("sfsdf", 1200)
	if err != nil {
		fmt.Printf("\n errors: %+v \n", err)
	}

}
