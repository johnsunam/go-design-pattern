package main

import "fmt"

func main() {
	mag1, _ := newPublication("magazine", "times", 50, "The Times")
	mag2, _ := newPublication("magazine", "Ranga", 50, "The Ranga")
	nwp1, _ := newPublication("newspaper", "blast", 5, "The blast")
	pubDetails(mag1)
	pubDetails(mag2)
	pubDetails(nwp1)
}

func pubDetails(pub IPublication) {
	fmt.Printf("----------- \n")
	fmt.Printf("%s \n", pub)
	fmt.Printf("Type: %T \n", pub)
	fmt.Printf("Name: %s \n", pub.getName())
	fmt.Printf("Page: %d \n", pub.getPage())
	fmt.Printf("Publisher: %s \n", pub.getPublisher())
	fmt.Printf("----------- \n")
}
